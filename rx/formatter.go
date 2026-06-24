package rx

import (
	"github.com/xjslang/xjs/ast"
	"github.com/xjslang/xjs/printer"
)

func Formatter(pr *printer.Printer, node ast.Node, next func(node ast.Node) error) error {
	switch v := node.(type) {
	case *rxVariable:
		pr.Print(v.Variable)
	case *rxLetStmt:
		pr.Print(v.LetStmt)
	case *rxAssignStmt:
		pr.Print(v.AssignStmt)
	default:
		return next(node)
	}
	return nil
}
