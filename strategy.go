package main

// 策略模式

import (
	"fmt"
)

type Runer interface {
	Run()
}

type Person struct{}

func (p *Person) Run() {
	fmt.Println("person run...")
}

type Dack struct{}

func (d *Dack) Run() {
	fmt.Println("dack run...")
}

func Play(runer Runer) {
	runer.Run()
}

func main() {
	p := new(Person)
	d := new(Dack)
	Play(p)
	Play(d)
}
