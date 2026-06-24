package rx

import (
	"github.com/xjslang/xjs/ast"
	"github.com/xjslang/xjs/builder"
	"github.com/xjslang/xjs/js"
	"github.com/xjslang/xjs/parser"
)

type RxVariable struct {
	ast.BaseExpr
	Variable *js.Variable
}

type RxLetStmt struct {
	ast.BaseStmt
	LetStmt *js.LetStmt
}

type RxAssignStmt struct {
	ast.BaseStmt
	AssignStmt *js.AssignStmt
}

func Plugin(b *builder.Builder) {
	b.UseStmtParser(func(p *parser.Parser, next func() (ast.Stmt, error)) (node ast.Stmt, err error) {
		if node, err = next(); err != nil {
			return
		}
		switch v := node.(type) {
		case *js.LetStmt:
			if v.Name.Literal[0] == '$' {
				w := &RxLetStmt{}
				w.LetStmt = v
				return w, nil
			}
		case *js.AssignStmt:
			if v.Name.Literal[0] == '$' {
				w := &RxAssignStmt{}
				w.AssignStmt = v
				return w, nil
			}
		}
		return
	})

	b.UseExprParser(func(p *parser.Parser, next func() (ast.Expr, error)) (node ast.Expr, err error) {
		if node, err = next(); err != nil {
			return
		}
		if node, ok := node.(*js.Variable); ok {
			if node.Name.Literal[0] == '$' {
				v := &RxVariable{}
				v.Variable = node
				return v, nil
			}
		}
		return
	})
}
