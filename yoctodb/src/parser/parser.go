package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"yocto/yoctodb/src/parser/grammer/parser"
)

type TreeShapeListener struct {
	*parser.BaseMySqlParserListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	//input, _ := antlr.NewFileStream(os.Args[1])
	//input := antlr.NewInputStream("SELECT NAME FROM AT WHERE NAME LIKE 'makang%'")
	input := antlr.NewInputStream("CREATE TABLE 'makang' (C1 INT, C2 VARCHAR(120))")
	lexer := parser.NewMySqlLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewMySqlParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.DdlStatement()
	//antlr.ParseTreeWalkerDefault.Walk()
	//antlr.ParseTreeVisitor(tree)
	//fmt.Println(tree.get)
	fmt.Println(tree.ToStringTree(p.GetTokenNames(), p.BaseParser))
}
