package main

import (
	"fmt"
)

type ITarget interface {
	TargetMethod(int, int, int)
}

type Target struct{}

func NewTarget() *Target {
	return &Target{}
}

func (t *Target) TargetMethod(a, b, c int) {
	fmt.Printf("this targetMethod, a:%d b:%d c:%d\n", a, b, c)
}

type IAdapter interface {
	AdapterMethod(int, int)
}

type Adapter struct {
	*Target
}

func NewAdapter(target *Target) *Adapter {
	return &Adapter{
		Target: target,
	}
}

func (ad *Adapter) AdapterMethod(a, b int) {
	fmt.Printf("this adapterMethod, a:%d b:%d\n", a, b)
	ad.Target.TargetMethod(a, b, 0)
}

func Product(adapter IAdapter) {
	adapter.AdapterMethod(1, 2)
}

func main() {
	target := NewTarget()
	adapter := NewAdapter(target)
	Product(adapter)
}
