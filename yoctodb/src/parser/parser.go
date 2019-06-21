package parser

import (
	"fmt"
	"yoctodb/src/server/grammer/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
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
	input := antlr.NewInputStream("SELECT NAME FROM AT WHERE NAME LIKE 'makang%'")
	lexer := parser.NewMySqlLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewMySqlParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Root()
	//antlr.ParseTreeWalkerDefault.Walk()
	//antlr.ParseTreeVisitor(tree)
	fmt.Println(tree.ToStringTree(p.GetTokenNames(), p.BaseParser))
}
