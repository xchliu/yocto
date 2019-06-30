package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"reflect"
	"yocto/src/parser/grammer/parser"
)

type QueryColumnDefine struct {
	cname    string
	datatype int
}

type SQLObject struct {
	DB           string
	SQLQuery     string
	SQLType      int
	SQLCommand   int
	QueryColumns []string
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

func (s *SQLObject) EnterCreateDefinitions(ctx *parser.CreateDefinitionsContext) {

	//for i, v := range ctx.GetChildren() {
	//	fmt.Println(ctx.GetChild(i), v, reflect.TypeOf(v))
	//}
	for _, v := range ctx.GetChildren() {
		//fmt.Println(i, v, reflect.TypeOf(v))
		for j, k := range v.GetChildren() {
			fmt.Println(j, k, reflect.TypeOf(k))
		}
	}
	for _, columnDefnition := range ctx.AllCreateDefinition() {
		//fmt.Println(i, v, reflect.TypeOf(v))
		for _, value := range columnDefnition.GetChildren() {
			if reflect.TypeOf(value).String() == "*parser.UidContext" {
				fmt.Println(value.(*parser.UidContext).GetText())
			}

			fmt.Println(value, reflect.TypeOf(value))
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
	test("CREATE TABLE DB1.T1(C1 INT PRIMARY KEY COMMENT 'JEIHEI', C2 VARCHAR(120)) ENGINE=INNODB AUTO_INCREMENT=6 DEFAULT CHARSET=UTF8MB4 ", "")
	fmt.Println()
}
