package main

import (
	"fmt"
)

// 模板方法
// 侧重的地方是抽象算法，封装规则

type ICoffeeinBeverage interface {
	BoilWater()
	Brew()
	PourInCup()
	AddCondiments()
	WantCondiment() bool
}

type CoffeeinBeverage struct {
	ICoffeeinBeverage
}

func (c *CoffeeinBeverage) BoilWater() {
	fmt.Println("boil water")
}

func (c *CoffeeinBeverage) PourInCup() {
	fmt.Println("Pour in cup")
}

func (c *CoffeeinBeverage) PrepareRecipe() {
	c.BoilWater()
	c.ICoffeeinBeverage.Brew()
	c.PourInCup()
	if c.ICoffeeinBeverage.WantCondiment() {
		c.ICoffeeinBeverage.AddCondiments()
	}
}

type Teal struct {
	*CoffeeinBeverage
}

func NewTeal() *Teal {
	teal := new(Teal)
	coffeeinBeverage := new(CoffeeinBeverage)
	coffeeinBeverage.ICoffeeinBeverage = teal
	teal.CoffeeinBeverage = coffeeinBeverage
	return teal
}

func (t *Teal) Brew() {
	fmt.Println("brew teal")
}

func (t *Teal) AddCondiments() {
	fmt.Println("add lemon")
}

func (t *Teal) WantCondiment() bool {
	return true
}

type Coffee struct {
	*CoffeeinBeverage
}

func NewCoffee() *Coffee {
	coffee := new(Coffee)
	coffeeinBeverage := new(CoffeeinBeverage)
	coffeeinBeverage.ICoffeeinBeverage = coffee
	coffee.CoffeeinBeverage = coffeeinBeverage
	return coffee
}

func (c *Coffee) Brew() {
	fmt.Println("brew coffee")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("add milk and sugar")
}

func (c *Coffee) WantCondiment() bool {
	return false
}

func main() {
	teal := NewTeal()
	teal.PrepareRecipe()
	fmt.Println("===================")
	coffee := NewCoffee()
	coffee.PrepareRecipe()
}
