package main

import (
	"fmt"
)

// 装饰者模式

type Seasoning interface {
	Desc()
}

type Egg struct{}

func (e *Egg) Desc() {
	fmt.Println("操蛋啦")
}

type Oil struct {
	Seasoning
}

func (o *Oil) Desc() {
	o.Seasoning.Desc()
	fmt.Println("放油")
}

type Salt struct {
	Seasoning
}

func (s *Salt) Desc() {
	s.Seasoning.Desc()
	fmt.Println("放盐")
}

func main() {
	var s *Salt = &Salt{
		Seasoning: &Oil{
			Seasoning: new(Egg),
		},
	}
	Start(s)
}

func Start(s Seasoning) {
	s.Desc()
}
