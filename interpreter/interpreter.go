package interpreter

import (
	"strings"
)

// Context 表示模板引擎的上下文
type Context struct {
	Variables map[string]string
}

// Expression 是模板中的一个抽象表达式
type Expression interface {
	Interpret(context *Context) string
}

// Literal 是一个表示文本字面量的表达式
type Literal struct {
	Text string
}

func (l *Literal) Interpret(context *Context) string {
	return l.Text
}

// Variable 是一个表示变量的表达式
type Variable struct {
	Name string
}

func (v *Variable) Interpret(context *Context) string {
	return context.Variables[v.Name]
}

// Template 解析模板并生成表达式列表
func Template(template string) []Expression {
	var expressions []Expression
	tokens := strings.Fields(template)

	for _, token := range tokens {
		if strings.HasPrefix(token, "{{") && strings.HasSuffix(token, "}}") {
			varName := strings.TrimSuffix(strings.TrimPrefix(token, "{{"), "}}")
			expressions = append(expressions, &Variable{varName})
		} else {
			expressions = append(expressions, &Literal{token})
		}
	}
	return expressions
}

// Render 渲染模板
func Render(expressions []Expression, context *Context) string {
	var result strings.Builder

	for _, expression := range expressions {
		result.WriteString(expression.Interpret(context))
		result.WriteString(" ")
	}

	return result.String()
}
