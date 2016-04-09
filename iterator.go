package main

import (
	"fmt"
)

// 迭代器与组合

// 抽象出来迭代
// 通过组合拼装

const (
	MAXITEM = 6
)

type Iterator interface {
	HasNext() bool
	Next() IMenuItem
}

type IMenuItem interface {
	Description() string
}

type IMenu interface {
	CreateIterator() Iterator
}

// MenuItem
type MenuItem struct {
	name        string
	description string
	vegetarin   bool
	price       float64
}

func NewMenuItem(name string, description string, vegetarin bool, price float64) *MenuItem {
	return &MenuItem{
		name:        name,
		description: description,
		vegetarin:   vegetarin,
		price:       price,
	}
}

func (m *MenuItem) Description() string {
	return fmt.Sprintf("%s %s %v %f", m.name, m.description, m.vegetarin, m.price)
}

// DinnerMenu
type DinnerMenu struct {
	maxItem   int
	menuItems []*MenuItem
}

func NewDinnerMenu() *DinnerMenu {
	return &DinnerMenu{
		maxItem:   MAXITEM,
		menuItems: make([]*MenuItem, 0, MAXITEM),
	}
}

func (d *DinnerMenu) AddMenuItem(menu *MenuItem) {
	if len(d.menuItems) > d.maxItem {
		fmt.Println("too many menu")
		return
	}

	d.menuItems = append(d.menuItems, menu)
}

func (d *DinnerMenu) AddMenu(name string, description string, vegetarin bool, price float64) {
	menuItem := NewMenuItem(name, description, vegetarin, price)
	d.AddMenuItem(menuItem)
}

func (d *DinnerMenu) CreateIterator() *DinnerIterator {
	return &DinnerIterator{
		index:      0,
		DinnerMenu: d,
	}
}

func (d *DinnerMenu) GetLength() int {
	return len(d.menuItems)
}

func (d *DinnerMenu) GetMenuItem(index int) *MenuItem {
	if index < len(d.menuItems) {
		return d.menuItems[index]
	}
	return nil
}

// DinnerIterator implements Iterator
type DinnerIterator struct {
	index int
	*DinnerMenu
}

func (d *DinnerIterator) HasNext() bool {
	if d.index < d.DinnerMenu.GetLength() {
		return true
	}
	return false
}

func (d *DinnerIterator) Next() IMenuItem {
	menuItem := d.DinnerMenu.GetMenuItem(d.index)
	d.index++
	return menuItem
}

// breakfastMenu
type BreakfastMenu struct {
	menuItems []*MenuItem
}

func NewBreakfastMenu() *BreakfastMenu {
	return &BreakfastMenu{
		menuItems: make([]*MenuItem, 0, MAXITEM),
	}
}

func (b *BreakfastMenu) AddMenuItem(menu *MenuItem) {
	b.menuItems = append(b.menuItems, menu)
}

func (b *BreakfastMenu) AddMenu(name string, description string, vegetarin bool, price float64) {
	menuItem := NewMenuItem(name, description, vegetarin, price)
	b.AddMenuItem(menuItem)
}

func (b *BreakfastMenu) CreateIterator() *BreakfastMenuIterator {
	return &BreakfastMenuIterator{
		index:         0,
		BreakfastMenu: b,
	}
}

func (b *BreakfastMenu) GetLength() int {
	return len(b.menuItems)
}

func (b *BreakfastMenu) GetMenuItem(index int) *MenuItem {
	if index < len(b.menuItems) {
		return b.menuItems[index]
	}
	return nil
}

// BreakfastMenuIterator implements Iterator
type BreakfastMenuIterator struct {
	index int
	*BreakfastMenu
}

func (b *BreakfastMenuIterator) HasNext() bool {
	if b.index < b.BreakfastMenu.GetLength() {
		return true
	}

	return false
}

func (b *BreakfastMenuIterator) Next() IMenuItem {
	menuItem := b.BreakfastMenu.GetMenuItem(b.index)
	b.index++
	return menuItem
}

// display
func Display(iterator Iterator) {
	for iterator.HasNext() {
		imenuItem := iterator.Next()
		fmt.Println(imenuItem.Description())
	}
}

func main() {
	breakfastMenu := NewBreakfastMenu()
	breakfastMenu.AddMenu("牛奶", "牛奶description", false, 3.0)
	breakfastMenu.AddMenu("油条", "油条description", false, 1.0)
	breakfastMenu.AddMenu("馒头", "馒头description", true, 1.0)
	breakfastMenu.AddMenu("豆浆", "DoujiangDescription", true, 1.5)

	Display(breakfastMenu.CreateIterator())

	dinnerMenu := NewDinnerMenu()
	dinnerMenu.AddMenu("香菇豆腐饭", "香菇豆腐", false, 10.5)
	dinnerMenu.AddMenu("蛋炒饭", "哈哈", false, 8.5)
	dinnerMenu.AddMenu("鱼香肉丝", "你猜", true, 15.5)
	dinnerMenu.AddMenu("咖啡", "大杯的哦", true, 10)
	Display(dinnerMenu.CreateIterator())
}
