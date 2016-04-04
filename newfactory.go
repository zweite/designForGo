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

	// 拿到类型
	tp := reflect.TypeOf(pr)

	// 新建对象
	vp := reflect.New(tp)
	vp = reflect.Indirect(vp)
	// 这里为解决，拿到的interface为nil
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

type NewCandy struct {
	Color string
}

func (n *NewCandy) Name() string {
	return "newCandy"
}

func (n *NewCandy) Desc() {
	if len(n.Color) > 0 {
		fmt.Println(n.Color)
	}
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
		candy, ok := pro.(*NewCandy)
		if ok && candy != nil {
			candy.Color = "红色"
			candy.Desc()
		}
	}
}
