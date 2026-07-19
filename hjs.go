package hjs

import (
	"github.com/xjslang/hjs/html"
	"github.com/xjslang/xjs"
	"github.com/xjslang/xjs/ast"
	"github.com/xjslang/xjs/js"
	"github.com/xjslang/xjs/jsextended"
	"github.com/xjslang/xjs/printer"
)

func Parse(input []byte) (*js.Program, error) {
	p := xjs.PluginBuilder().
		Install(jsextended.Plugin).
		Install(html.Plugin).
		Build(input)
	return js.ParseProgram(p)
}

func Compile(result ast.Node) (string, error) {
	pr := xjs.PrinterBuilder().
		UsePrinter(jsextended.Printer).
		UsePrinter(html.Compiler).
		Build(printer.Compact())
	pr.Print(result)
	return pr.Output()
}

func Format(result ast.Node, opts ...printer.Option) (string, error) {
	pr := xjs.PrinterBuilder().
		UsePrinter(jsextended.Printer).
		UsePrinter(html.Formatter).
		Build(opts...)
	pr.Print(result)
	return pr.Output()
}
