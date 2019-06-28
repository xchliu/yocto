package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"yocto/src/parser/grammer/parser"
)

type SQLObject struct {
	DB         string
	SQLQuery   string
	SQLType    int
	SQLCommand int
	*YoctoSQLBaseListener
}

type YoctoSQLBaseListener struct {
	*parser.BaseMySqlParserListener
}

func YoctoSQLListener() *YoctoSQLBaseListener {
	//a := new(SQLObject)
	return new(YoctoSQLBaseListener)
}

type YYoctoSQLListener struct {
	antlr.ParseTreeListener
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

func test(q string) {
	s := new(SQLObject)
	s.SQLQuery = q
	p := Build_paser(s.SQLQuery, "")
	tree := p.Root()
	antlr.ParseTreeWalkerDefault.Walk(YoctoSQLListener(), tree)
	//fmt.Println(sql_obj.pTree.ToStringTree(sql_obj.p.GetTokenNames(), sql_obj.p.BaseParser))
	fmt.Println(s.SQLType)
}

func main() {
	test("CREATE TABLE T1(C1 INT)")
	fmt.Println()
}
