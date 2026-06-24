package rx

import (
	"github.com/xjslang/xjs/ast"
	"github.com/xjslang/xjs/printer"
)

const Runtime = `function rx(initVal) {
  let _val = initVal;
  let _listeners = [];
  return {
    get: function () {
      return _val;
    },
    set: function (val) {
      _val = val;
      for (let listener of _listeners) {
        listener();
      }
    },
  };
}`

func Compiler(pr *printer.Printer, node ast.Node, next func(node ast.Node) error) error {
	switch v := node.(type) {
	case *rxVariable:
		pr.Print(v.Variable.Name, ".get()")
	case *rxLetStmt:
		pr.Print("let ", v.LetStmt.Name, " = rx(", v.LetStmt.Value, ")")
	case *rxAssignStmt:
		pr.Print(v.AssignStmt.Name, ".set(", v.AssignStmt.Value, ")")
	default:
		return next(node)
	}
	return nil
}
