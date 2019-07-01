package main

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
}

type SQLObject struct {
	DB            string
	SQLQuery      string
	SQLType       int
	SQLCommand    int
	CreateColumns []CreateColumnDefine
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

func (s *SQLObject) EnterColumnDefinition(ctx *parser.ColumnDefinitionContext) {
	//fmt.Println(ctx.DataType().GetText())
}

func (this *SQLObject) EnterCreateDefinitions(ctx *parser.CreateDefinitionsContext) {
	tableColumn := new(CreateColumnDefine)
	// for all cols
	for _, createDefinition := range ctx.AllCreateDefinition() {
		//for one col, including uid and definitions
		for _, value := range createDefinition.GetChildren() {

			switch value.(type) {

			case *parser.UidContext:
				tableColumn.cname = value.(*parser.UidContext).SimpleId().GetText()

			case *parser.ColumnDefinitionContext:
				for _, columnDefinition := range value.GetChildren() {

					switch columnDefinition.(type) {
					case *parser.DimensionDataTypeContext:
						tableColumn.datatype = columnDefinition.(*parser.DimensionDataTypeContext).GetTypeName().GetTokenType()

						if columnDefinition.GetChildCount() > 1 {
							for _, length := range columnDefinition.GetChildren() {
								switch length.(type) {

								case *parser.LengthOneDimensionContext:
									tableColumn.clength = length.(*parser.LengthOneDimensionContext).DecimalLiteral().GetText()

								case *parser.LengthTwoDimensionContext:
									tableColumn.clength = length.(*parser.LengthTwoDimensionContext).DecimalLiteral(0).GetText()
									tableColumn.cprecision = length.(*parser.LengthTwoDimensionContext).DecimalLiteral(0).GetText()

								case *parser.LengthTwoOptionalDimensionContext:
									tableColumn.clength = length.(*parser.LengthTwoOptionalDimensionContext).DecimalLiteral(0).GetText()
									tableColumn.cprecision = length.(*parser.LengthTwoOptionalDimensionContext).DecimalLiteral(0).GetText()
								}
							}
						}

						//case *parser.StringDataTypeContext:
						//case *parser.NationalStringDataTypeContext:
						//case *parser.NationalVaryingStringDataTypeContext:
						//case *parser.SimpleDataTypeContext:
						//case *parser.CollectionDataTypeContext:
						//case *parser.SpatialDataTypeContext:
						//case *parser.
						//	fmt.Println(value.(*parser.ColumnDefinitionContext).DataType())
						//	if columnDefinition.GetChildCount() == 1 {
						//		fmt.Println(reflect.TypeOf(columnDefinition.GetChild(0)))
						//	}

						//case *parser.PrimaryKeyColumnConstraintContext:
						//	fmt.Println(columnDefinition.(*parser.PrimaryKeyColumnConstraintContext).PRIMARY().GetText())
						//case *parser.NullColumnConstraintContext:
						//	fmt.Println(columnDefinition.(*parser.NullColumnConstraintContext).NullNotnull().GetText())
						//case *parser.CommentColumnConstraintContext:
						//	fmt.Println(columnDefinition.(*parser.CommentColumnConstraintContext).COMMENT().GetText())
						//case *parser.DefaultColumnConstraintContext:
						//	fmt.Println(columnDefinition.(*parser.DefaultColumnConstraintContext).DEFAULT().GetText())
						//fmt.Println(columnDefinition.(*parser.ColumnDefinitionContext).DataType())
					}
				}
			}
			//
			//if reflect.TypeOf(value).String() == "*parser.ColumnDefinitionContext" {
			//	fmt.Println(value.(*parser.ColumnDefinitionContext).DataType().GetText())
			//}

			//if reflect.TypeOf(value).String() == ""
			//value.(*parser.ColumnConstraintContext).
			//fmt.Println(value, reflect.TypeOf(value))
			//for i, cv := range value.GetChildren() {
			//	fmt.Println(i, cv, "haha", reflect.TypeOf(cv))
			//}
		}
	}
}

func test(q, db string) {
	p := Build_paser(q, "")
	tree := p.Root()
	ss := new(SQLObject)
	antlr.ParseTreeWalkerDefault.Walk(ss, tree)
	fmt.Println(ss.SQLType)
	fmt.Println(ss.SQLCommand)
	fmt.Println(tree.ToStringTree(p.GetTokenNames(), p.BaseParser))
}

func main() {
	test("CREATE TABLE DB1.T1"+
		"(C1 INT(14) NOT NULL PRIMARY KEY COMMENT 'JEIHEI', "+
		"C2 VARCHAR(120) DEFAULT 'ABCD' COMMENT 'WANGBADAN', "+
		"C3 DECIMAL(12, 13),"+
		"C4 DATE"+
		") ENGINE=INNODB AUTO_INCREMENT=6 DEFAULT CHARSET=UTF8MB4 ", "")
	fmt.Println()
}
