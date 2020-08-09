package main

import "fmt"

// Operand は処理対象を表すインタフェース
type Operand interface {
	GetOperandString() string
}

// Operator は処理を表すインタフェース
type Operator interface {
	Execute() Operand
}

// Ingredient は処理対象を表すクラス
type Ingredient struct {
	name string
}

// NewIngredient is XXX
func NewIngredient(name string) *Ingredient {
	return &Ingredient{
		name: name,
	}
}

// GetOperandString is XXX
func (i *Ingredient) GetOperandString() string {
	return i.name
}

// Expression は処理結果を表すクラス
type Expression struct {
	operator Operator
}

// NewExpression is XXX
func NewExpression(operator Operator) *Expression {
	return &Expression{
		operator: operator,
	}
}

// GetOperandString is XXX
func (i *Expression) GetOperandString() string {
	return i.operator.Execute().GetOperandString()
}

// Plus は2つの operand を足し合わせる処理を表すクラス
type Plus struct {
	operand1 Operand
	operand2 Operand
}

// NewPlus is XXX
func NewPlus(operand1, operand2 Operand) *Plus {
	return &Plus{
		operand1: operand1,
		operand2: operand2,
	}
}

// Execute is XXX
func (p *Plus) Execute() Operand {
	return NewIngredient(fmt.Sprintf("%sと%sを足したもの", p.operand1.GetOperandString(), p.operand2.GetOperandString()))
}

// Wait は2つの operand を足し合わせる処理を表すクラス
type Wait struct {
	operand Operand
	minutes uint32
}

// NewWait is XXX
func NewWait(operand Operand, minutes uint32) *Wait {
	return &Wait{
		operand: operand,
		minutes: minutes,
	}
}

// Execute is XXX
func (w *Wait) Execute() Operand {
	return NewIngredient(fmt.Sprintf("%sを%d分置いたもの", w.operand.GetOperandString(), w.minutes))
}

func main() {
	i1 := NewIngredient("ingredient1")
	i2 := NewIngredient("ingredient2")

	resPlus := NewPlus(i1, i2).Execute()
	fmt.Println(resPlus.GetOperandString())

	resWait := NewWait(resPlus, 10).Execute()
	fmt.Println(resWait.GetOperandString())
}
