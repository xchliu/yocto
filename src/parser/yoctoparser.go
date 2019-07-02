package yoctoparser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"yocto/src/parser/grammer/parser"
)

type CreateColumnDefine struct {
	cname      string
	datatype   int
	clength    string
	cprecision string
	constraint []CreateColumnConstraint
}

type CreateColumnConstraint struct {
	constraintType int
	flag           bool
	str            string
}

type QueryColumnDefine struct {
}

type SQLObject struct {
	DB            string
	SQLQuery      string
	SQLType       int
	SQLCommand    int
	CreateColumns []CreateColumnDefine
	QueryColumns  []QueryColumnDefine
	YoctoSQLBaseListener
}

type YoctoSQLBaseListener struct {
	*parser.BaseMySqlParserListener
}

func Build_paser(query, db string) (p *parser.MySqlParser) {
	input := antlr.NewInputStream(query)
	lexer := parser.NewMySqlLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p = parser.NewMySqlParser(stream)
	//p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	return p
}

func (this *SQLObject) EnterDdlStatement(ctx *parser.DdlStatementContext) {
	this.SQLType = parser.MySqlParserRULE_ddlStatement
}

func (this *SQLObject) EnterDmlStatement(ctx *parser.DmlStatementContext) {
	this.SQLType = parser.MySqlParserRULE_dmlStatement
}

func (this *SQLObject) EnterColumnCreateTable(ctx *parser.ColumnCreateTableContext) {
	this.SQLCommand = parser.MySqlParserRULE_createTable
}

func (tableColumn *CreateColumnDefine) getDataLength(length interface{}) {
	switch length.(type) {
	case *parser.LengthOneDimensionContext:
		tableColumn.clength = length.(*parser.LengthOneDimensionContext).DecimalLiteral().GetText()

	case *parser.LengthTwoDimensionContext:
		tableColumn.clength = length.(*parser.LengthTwoDimensionContext).DecimalLiteral(0).GetText()
		tableColumn.cprecision = length.(*parser.LengthTwoDimensionContext).DecimalLiteral(1).GetText()

	case *parser.LengthTwoOptionalDimensionContext:
		tableColumn.clength = length.(*parser.LengthTwoOptionalDimensionContext).DecimalLiteral(0).GetText()
		tableColumn.cprecision = length.(*parser.LengthTwoOptionalDimensionContext).DecimalLiteral(1).GetText()
	}
}

func (tableColumn *CreateColumnDefine) GetColumnInfo(ctx *parser.ColumnDefinitionContext) {
	for _, columnDefinition := range ctx.GetChildren() {
		//fmt.Println(reflect.TypeOf(columnDefinition))
		switch columnDefinition.(type) {
		case *parser.DimensionDataTypeContext:
			{
				tableColumn.datatype = columnDefinition.(*parser.DimensionDataTypeContext).GetTypeName().GetTokenType()
				for _, length := range columnDefinition.GetChildren() {
					tableColumn.getDataLength(length)
				}
			}

		case *parser.StringDataTypeContext:
			{
				tableColumn.datatype = columnDefinition.(*parser.StringDataTypeContext).GetTypeName().GetTokenType()
				for _, length := range columnDefinition.GetChildren() {
					tableColumn.getDataLength(length)
				}
			}

		case *parser.NationalStringDataTypeContext:
			{
				tableColumn.datatype = columnDefinition.(*parser.NationalStringDataTypeContext).GetTypeName().GetTokenType()
				for _, length := range columnDefinition.GetChildren() {
					tableColumn.getDataLength(length)
				}
			}

		case *parser.NationalVaryingStringDataTypeContext:
			{
				tableColumn.datatype = columnDefinition.(*parser.NationalVaryingStringDataTypeContext).GetTypeName().GetTokenType()
				for _, length := range columnDefinition.GetChildren() {
					tableColumn.getDataLength(length)
				}
			}

		case *parser.SimpleDataTypeContext:
			{
				tableColumn.datatype = columnDefinition.(*parser.SimpleDataTypeContext).GetTypeName().GetTokenType()
				for _, length := range columnDefinition.GetChildren() {
					tableColumn.getDataLength(length)
				}
			}

		case *parser.CollectionDataTypeContext:
			{
				tableColumn.datatype = columnDefinition.(*parser.CollectionDataTypeContext).GetTypeName().GetTokenType()
				for _, length := range columnDefinition.GetChildren() {
					tableColumn.getDataLength(length)
				}
			}

		case *parser.SpatialDataTypeContext:
			{
				tableColumn.datatype = columnDefinition.(*parser.SpatialDataTypeContext).GetTypeName().GetTokenType()
				for _, length := range columnDefinition.GetChildren() {
					tableColumn.getDataLength(length)
				}
			}

		case *parser.PrimaryKeyColumnConstraintContext:
			{
				pkConstraint := CreateColumnConstraint{
					constraintType: parser.MySqlParserPRIMARY,
					flag:           true,
					str:            ""}
				tableColumn.constraint = append(tableColumn.constraint, pkConstraint)
			}

		case *parser.CommentColumnConstraintContext:
			{
				cmtConstraint := CreateColumnConstraint{
					constraintType: parser.MySqlParserCOMMENT,
					flag:           true,
					str:            columnDefinition.(*parser.CommentColumnConstraintContext).STRING_LITERAL().GetText()}
				tableColumn.constraint = append(tableColumn.constraint, cmtConstraint)
			}

		case *parser.DefaultColumnConstraintContext:
			{
				dftConstraint := CreateColumnConstraint{
					constraintType: parser.MySqlParserDEFAULT,
					flag:           true,
					str:            columnDefinition.(*parser.DefaultColumnConstraintContext).DefaultValue().GetText()}
				tableColumn.constraint = append(tableColumn.constraint, dftConstraint)
			}

		case *parser.NullColumnConstraintContext:
			{
				nullConstraint := CreateColumnConstraint{}
				isNullNotnull := columnDefinition.(*parser.NullColumnConstraintContext).NullNotnull().GetText()
				if isNullNotnull == "NOTNULL" {
					nullConstraint = CreateColumnConstraint{
						constraintType: parser.MySqlParserNULLIF,
						flag:           true,
						str:            ""}
				} else {
					nullConstraint = CreateColumnConstraint{
						constraintType: parser.MySqlParserNULLIF,
						flag:           false,
						str:            ""}
				}
				tableColumn.constraint = append(tableColumn.constraint, nullConstraint)
			}
		}
	}
}

func (this *SQLObject) EnterCreateDefinitions(ctx *parser.CreateDefinitionsContext) {
	// for all cols
	for _, createDefinition := range ctx.AllCreateDefinition() {
		tableColumn := new(CreateColumnDefine)
		//for one col, including uid and definitions
		for _, value := range createDefinition.GetChildren() {
			switch value.(type) {

			case *parser.UidContext:
				tableColumn.cname = value.(*parser.UidContext).SimpleId().GetText()

			case *parser.ColumnDefinitionContext:
				tableColumn.GetColumnInfo(value.(*parser.ColumnDefinitionContext))
			}
		}
		this.CreateColumns = append(this.CreateColumns, *tableColumn)
	}

}

func test(q, db string) {
	p := Build_paser(q, "")
	tree := p.Root()
	ss := new(SQLObject)
	antlr.ParseTreeWalkerDefault.Walk(ss, tree)
	fmt.Println(ss.CreateColumns)
	fmt.Println(tree.ToStringTree(p.GetTokenNames(), p.BaseParser))
}

func YoctoPaser(query, db string) (s SQLObject) {
	p := Build_paser(query, db)
	tree := p.Root()
	antlr.ParseTreeWalkerDefault.Walk(s, tree)
	return s
}

func main() {
	test("CREATE TABLE DB1.T1"+
		"(C1 INT(14) NOT NULL PRIMARY KEY COMMENT 'JEIHEI', "+
		"C2 VARCHAR(120) DEFAULT 'ABCD' COMMENT 'WANGBADAN', "+
		"C3 DECIMAL(12, 13) NOT NULL,"+
		"C4 DATE NULL"+
		") ENGINE=INNODB AUTO_INCREMENT=6 DEFAULT CHARSET=UTF8MB4 ", "")
	//test("CREATE TABLE DB1.T1 (C2 VARCHAR(120) DEFAULT 'ABCD' COMMENT 'WANGBADAN') ENGINE=INNODB AUTO_INCREMENT=6 DEFAULT CHARSET=UTF8MB4 ", "")
}
