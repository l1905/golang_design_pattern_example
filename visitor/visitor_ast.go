package visitor

import (
	"fmt"
	"go/ast"
)

// 定义一个实现ast.Visitor接口的结构体和Visit方法
type FuncVisitor struct{}

func (v *FuncVisitor) Visit(node ast.Node) ast.Visitor {
	if fn, ok := node.(*ast.FuncDecl); ok {
		fmt.Printf("Function name: %s\n", fn.Name.Name)
	}
	return v
}
