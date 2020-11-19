package java

import (
	"log"

	"github.com/ZupIT/horusec-engine/java/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type JavaTreeListener struct {
	*parser.BaseJavaParserListener
}

func NewJavaTreeListener() *JavaTreeListener {
	return new(JavaTreeListener)
}

func (treeListener *JavaTreeListener) EnterImportDeclaration(ctx antlr.ParserRuleContext) {
	log.Println(ctx.GetText())
	log.Println("enter import declaration")
}

func (treeListener *JavaTreeListener) EnterPackageDeclaration(ctx antlr.ParserRuleContext) {
	log.Println(ctx.GetText())
	log.Println("enter package declaration")
}

// func (treeListener *JavaTreeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
// 	log.Println(ctx.GetText())
// 	log.Println("every rule triggered")
// }

func LoadJavaSourceFile(filename string) {
	log.SetFlags(0)
	log.SetPrefix("engine> ")

	inputFileStream, err := antlr.NewFileStream(filename)

	if err != nil {
		log.Printf("error: %s", err.Error())
	}

	lexer := parser.NewJavaLexer(inputFileStream)

	stream := antlr.NewCommonTokenStream(lexer, 0)

	javaParser := parser.NewJavaParser(stream)
	// javaParser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	javaParser.BuildParseTrees = true

	compilationUnit := javaParser.CompilationUnit()

	ctx := parser.NewCompilationUnitContext(javaParser, compilationUnit, 0)

	packages := ctx.PackageDeclaration()
	imports := ctx.AllImportDeclaration()

	antlr.ParseTreeWalkerDefault.Walk(NewJavaTreeListener(), packages)
	antlr.ParseTreeWalkerDefault.Walk(NewJavaTreeListener(), imports)
}
