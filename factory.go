package main

import (
	"errors"
	"fmt"
)

# 简单工厂模式

type Product interface {
	Desc()
}

type Candy struct{}

func (c *Candy) Desc() {
	fmt.Println("这个是糖果")
}

type Cookie struct{}

func (c *Cookie) Desc() {
	fmt.Println("这个是饼干")
}

func Factory(name string) (p Product, err error) {
	switch name {
	case "candy":
		p = new(Candy)
	case "cookie":
		p = new(Cookie)
	default:
		err = errors.New("unknow product")
	}
	return
}

func Desc(p Product) {
	p.Desc()
}

func main() {
	p, err := Factory("candy")
	if err != nil {
		panic(err)
	}
	Desc(p)

	p, err = Factory("cookie")
	if err != nil {
		panic(err)
	}
	Desc(p)
}
