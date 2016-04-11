package main

import (
	"fmt"
)

// 代理模式

type Animal interface {
	Run()
}

type Human struct{}

func NewHuman() *Human {
	return new(Human)
}

func (h *Human) Run() {
	fmt.Println("有人在跑...")
}

type Duck struct{}

func NewDuck() *Duck {
	return new(Duck)
}

func (d *Duck) Run() {
	fmt.Println("鸭子在跑")
}

func Run(a Animal) {
	a.Run()
}

func main() {
	h := NewHuman()
	Run(h)

	d := NewDuck()
	Run(d)
}
