package rx

import (
	"github.com/xjslang/xjs/ast"
	"github.com/xjslang/xjs/printer"
)

func Formatter(pr *printer.Printer, node ast.Node, next func(node ast.Node) error) error {
	switch v := node.(type) {
	case *RxVariable:
		pr.Print(v.Variable)
	case *RxLetStmt:
		pr.Print(v.LetStmt)
	case *RxAssignStmt:
		pr.Print(v.AssignStmt)
	default:
		return next(node)
	}
	return nil
}
