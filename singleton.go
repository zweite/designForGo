package main

import (
	"errors"
	"fmt"
	"sync"
)

// 单例模式
// 需要注意并发。此Demo使用加锁解决并发

var bCard *busCard

type busCard struct {
	mx   sync.Mutex
	Cost float64
}

func (b *busCard) Recharge(cost float64) {
	b.mx.Lock()
	defer b.mx.Unlock()
	b.Cost += cost
	return
}

func (b *busCard) Consumption(cost float64) (err error) {
	b.mx.Lock()
	defer b.mx.Unlock()
	if b.Cost < cost {
		err = errors.New("余额不足！")
		return
	}

	b.Cost -= cost
	return
}

func (b *busCard) Balance() (cost float64) {
	return b.Cost
}

func InstanceBusCard() *busCard {
	if bCard == nil {
		bCard = new(busCard)
	}

	return bCard
}

func main() {
	busCard1 := InstanceBusCard()
	busCard2 := InstanceBusCard()
	fmt.Println("busCard1 equest busCard2 ?", busCard1 == busCard2)
	busCard1.Recharge(100)
	fmt.Println("busCard2 balance =", busCard2.Balance())
}
