package parser

import (
	"fmt"
	"strings"
	"yocto/src/yoctoparser/grammer/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type CreateColumnDefine struct {
	Cname      string
	Datatype   int
	Clength    string
	Cprecision string
	Constraint []CreateColumnConstraint
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
	TableList     []string
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
	objectName := strings.Split(ctx.TableName().GetText(), ".")
	if len(objectName) == 2 {
		this.DB = objectName[0]
		this.TableList = append(this.TableList, objectName[1])
	} else {
		this.TableList = append(this.TableList, objectName[0])
	}
}

func (tableColumn *CreateColumnDefine) getDataLength(length interface{}) {
	switch length.(type) {
	case *parser.LengthOneDimensionContext:
		tableColumn.Clength = length.(*parser.LengthOneDimensionContext).DecimalLiteral().GetText()

	case *parser.LengthTwoDimensionContext:
		tableColumn.Clength = length.(*parser.LengthTwoDimensionContext).DecimalLiteral(0).GetText()
		tableColumn.Cprecision = length.(*parser.LengthTwoDimensionContext).DecimalLiteral(1).GetText()

	case *parser.LengthTwoOptionalDimensionContext:
		tableColumn.Clength = length.(*parser.LengthTwoOptionalDimensionContext).DecimalLiteral(0).GetText()
		tableColumn.Cprecision = length.(*parser.LengthTwoOptionalDimensionContext).DecimalLiteral(1).GetText()
	}
}

func (tableColumn *CreateColumnDefine) GetColumnInfo(ctx *parser.ColumnDefinitionContext) {
	for _, columnDefinition := range ctx.GetChildren() {
		//fmt.Println(reflect.TypeOf(columnDefinition))
		switch columnDefinition.(type) {
		case *parser.DimensionDataTypeContext:
			{
				tableColumn.Datatype = columnDefinition.(*parser.DimensionDataTypeContext).GetTypeName().GetTokenType()
				for _, length := range columnDefinition.GetChildren() {
					tableColumn.getDataLength(length)
				}
			}

		case *parser.StringDataTypeContext:
			{
				tableColumn.Datatype = columnDefinition.(*parser.StringDataTypeContext).GetTypeName().GetTokenType()
				for _, length := range columnDefinition.GetChildren() {
					tableColumn.getDataLength(length)
				}
			}

		case *parser.NationalStringDataTypeContext:
			{
				tableColumn.Datatype = columnDefinition.(*parser.NationalStringDataTypeContext).GetTypeName().GetTokenType()
				for _, length := range columnDefinition.GetChildren() {
					tableColumn.getDataLength(length)
				}
			}

		case *parser.NationalVaryingStringDataTypeContext:
			{
				tableColumn.Datatype = columnDefinition.(*parser.NationalVaryingStringDataTypeContext).GetTypeName().GetTokenType()
				for _, length := range columnDefinition.GetChildren() {
					tableColumn.getDataLength(length)
				}
			}

		case *parser.SimpleDataTypeContext:
			{
				tableColumn.Datatype = columnDefinition.(*parser.SimpleDataTypeContext).GetTypeName().GetTokenType()
				for _, length := range columnDefinition.GetChildren() {
					tableColumn.getDataLength(length)
				}
			}

		case *parser.CollectionDataTypeContext:
			{
				tableColumn.Datatype = columnDefinition.(*parser.CollectionDataTypeContext).GetTypeName().GetTokenType()
				for _, length := range columnDefinition.GetChildren() {
					tableColumn.getDataLength(length)
				}
			}

		case *parser.SpatialDataTypeContext:
			{
				tableColumn.Datatype = columnDefinition.(*parser.SpatialDataTypeContext).GetTypeName().GetTokenType()
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
				tableColumn.Constraint = append(tableColumn.Constraint, pkConstraint)
			}

		case *parser.CommentColumnConstraintContext:
			{
				cmtConstraint := CreateColumnConstraint{
					constraintType: parser.MySqlParserCOMMENT,
					flag:           true,
					str:            columnDefinition.(*parser.CommentColumnConstraintContext).STRING_LITERAL().GetText()}
				tableColumn.Constraint = append(tableColumn.Constraint, cmtConstraint)
			}

		case *parser.DefaultColumnConstraintContext:
			{
				dftConstraint := CreateColumnConstraint{
					constraintType: parser.MySqlParserDEFAULT,
					flag:           true,
					str:            columnDefinition.(*parser.DefaultColumnConstraintContext).DefaultValue().GetText()}
				tableColumn.Constraint = append(tableColumn.Constraint, dftConstraint)
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
				tableColumn.Constraint = append(tableColumn.Constraint, nullConstraint)
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
				tableColumn.Cname = value.(*parser.UidContext).SimpleId().GetText()

			case *parser.ColumnDefinitionContext:
				tableColumn.GetColumnInfo(value.(*parser.ColumnDefinitionContext))
			}
		}
		this.CreateColumns = append(this.CreateColumns, *tableColumn)
	}

}

func test(q, db string) {
	p := Build_paser(q, db)
	tree := p.Root()
	ss := new(SQLObject)
	ss.DB = db
	antlr.ParseTreeWalkerDefault.Walk(ss, tree)
	fmt.Println(ss.CreateColumns)
	fmt.Println(ss)
	fmt.Println(tree.ToStringTree(p.GetTokenNames(), p.BaseParser))
}

func YoctoPaser(query, db string) (s *SQLObject) {
	ss := new(SQLObject)
	ss.DB = db
	ss.SQLQuery = query
	p := Build_paser(query, db)
	tree := p.Root()
	antlr.ParseTreeWalkerDefault.Walk(ss, tree)
	fmt.Println(ss)
	//s = *ss
	return ss
}

func main() {
	YoctoPaser("CREATE TABLE T1"+
		"(C1 INT(14) NOT NULL PRIMARY KEY COMMENT 'JEIHEI', "+
		"C2 VARCHAR(120) DEFAULT 'ABCD' COMMENT 'WANGBADAN', "+
		"C3 DECIMAL(12, 13) NOT NULL,"+
		"C4 DATE NULL"+
		") ENGINE=INNODB AUTO_INCREMENT=6 DEFAULT CHARSET=UTF8MB4 ", "aaaa")
	//test("CREATE TABLE DB1.T1 (C2 VARCHAR(120) DEFAULT 'ABCD' COMMENT 'WANGBADAN') ENGINE=INNODB AUTO_INCREMENT=6 DEFAULT CHARSET=UTF8MB4 ", "")
}
