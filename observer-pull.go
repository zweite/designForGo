package main

import (
	"fmt"
)

// 观察者模式
// 拉模式

// 观察者
type Obserable interface {
	Update(Subject)
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

func (g *Guest) Update(s Subject) {
	k, ok := s.(*Kitchen)
	if ok {
		fmt.Printf("%d update %s\n", g.Id, k.Info)
	}
}

func (g *Guest) GetId() int {
	return g.Id
}

type Kitchen struct {
	Info       string
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

func (k *Kitchen) NotifyObserver() {
	for _, obs := range k.Obserables {
		obs.Update(k)
	}
}

func (k *Kitchen) Serving(info string) {
	k.Info = info
	k.NotifyObserver()
}

func main() {
	guestNum := 10
	subject := NewKitchen()
	guests := make([]*Guest, 0, guestNum)
	for i := 0; i < guestNum; i++ {
		guest := &Guest{
			Id: i,
		}
		subject.RegistterObserver(guest)
		guests = append(guests, guest)
	}

	subject.Serving("上菜了")

	for _, guest := range guests {
		subject.RemoveObserver(guest)
	}
}
