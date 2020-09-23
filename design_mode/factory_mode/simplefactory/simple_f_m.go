package simplefactory

import "fmt"

/* 简单工厂:New创建方法对包外可见,用于唯一工厂来创建对象*/
type API interface {
	Say(name string) string
}

type hiAPI struct{}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi , %s", name)
}

type helloAPI struct{}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}
