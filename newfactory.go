package main

import (
	"fmt"
	"reflect"
)

// 工厂模式，动态注册生产线

type NewProduct interface {
	Name() string
	Desc()
}

type NewFactory struct {
	lines map[string]NewProduct
}

func (n *NewFactory) Register(np NewProduct) {
	n.lines[np.Name()] = np
}

func (n *NewFactory) GetNewProduct(name string) (pr NewProduct) {
	pr, ok := n.lines[name]
	if !ok {
		return
	}

	tp := reflect.TypeOf(pr)
	vp := reflect.Zero(tp)

	ip := vp.Interface()
	pr, ok = ip.(NewProduct)
	if !ok {
		fmt.Println("类型转换出错")
	}
	return
}

func GetNewFactory() *NewFactory {
	return &NewFactory{
		lines: make(map[string]NewProduct),
	}
}

type NewCandy struct{}

func (n *NewCandy) Name() string {
	return "newCandy"
}

func (n *NewCandy) Desc() {
	fmt.Println("这个是新糖果")
}

type NewCookie struct{}

func (n *NewCookie) Name() string {
	return "newCookie"
}

func (n *NewCookie) Desc() {
	fmt.Println("这个是新饼干")
}

func main() {
	factory := GetNewFactory()
	factory.Register(new(NewCandy))
	factory.Register(new(NewCookie))

	pro := factory.GetNewProduct("newCandy")
	if pro != nil {
		pro.Desc()
	}
}
