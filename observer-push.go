package main

import (
	"fmt"
)

// 观察者模式
// 推模式

// 观察者
type Obserable interface {
	Update(string)
	GetId() int
}

// 主题接口
type Subject interface {
	RegistterObserver(Obserable)
	RemoveObserver(Obserable)
	NotifyObserver()
}

type Guest struct {
	Id int
}

func (g *Guest) Update(info string) {
	fmt.Printf("%d update %s\n", g.Id, info)
}

func (g *Guest) GetId() int {
	return g.Id
}

type Kitchen struct {
	Obserables map[int]Obserable
}

func NewKitchen() *Kitchen {
	return &Kitchen{
		Obserables: make(map[int]Obserable),
	}
}

func (k *Kitchen) RegistterObserver(obs Obserable) {
	k.Obserables[obs.GetId()] = obs
}

func (k *Kitchen) RemoveObserver(obs Obserable) {
	delete(k.Obserables, obs.GetId())
}

func (k *Kitchen) NotifyObserver(info string) {
	for _, obs := range k.Obserables {
		obs.Update(info)
	}
}

func (k *Kitchen) Serving(info string) {
	k.NotifyObserver(info)
}

func main() {
	guestNum := 10
	k := NewKitchen()
	guests := make([]*Guest, 0, guestNum)
	for i := 0; i < guestNum; i++ {
		guest := &Guest{
			Id: i,
		}
		k.RegistterObserver(guest)
		guests = append(guests, guest)
	}

	k.Serving("上菜了")

	for _, guest := range guests {
		k.RemoveObserver(guest)
	}
}
