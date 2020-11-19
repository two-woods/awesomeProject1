package main

import (
	"fmt"
)

/*工厂方法模式使用子类的方式延迟生成对象到子类中实现*/

//Operaor 是被封装的实际类接口，加减运算的抽象
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

//工厂接口，创建Operator类，运算的工厂的抽象
type OperatorFactory interface {
	Create() Operator
}

//实现Operator的基类，封装公共方法
type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	fmt.Println("SetASetASetASetA")
	o.a = a
}

func (o *OperatorBase) SetB(b int) {
	fmt.Println("SetBSetBSetBSetB")
	o.b = b
}

//--------------------------------------------------
//PlusOperatorFactory  是PlusOperator的工厂类
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	//子类中实现父类PlusOperator
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

type PlusOperator struct {
	*OperatorBase
}

func (o PlusOperator) Result() int {
	return o.a + o.b
}

//----------------------------------------------------
//MinusOperatorFactory 是MinusOperator的工厂类
type MinusOperatorFactory struct{}

//实现OperatorFactory接口
func (MinusOperatorFactory) Create() Operator {
	//子类中实现父类PlusOperator
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

type MinusOperator struct {
	*OperatorBase
}

func (o MinusOperator) Result() int {
	return o.a - o.b
}

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func main() {
	var (
		factory OperatorFactory
	)
	factory = PlusOperatorFactory{}
	r1 := compute(factory, 1, 2)
	println("-=-=-=", r1)
	factory = MinusOperatorFactory{}
	r2 := compute(factory, 4, 2)
	println("-=-=-=", r2)

}
