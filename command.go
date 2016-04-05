package main

import (
	"fmt"
)

// 命令模式

type TV struct {
	channel int
	brand   string
}

func NewTV(brand string, channel int) *TV {
	return &TV{
		channel: channel,
		brand:   brand,
	}
}

func (t *TV) TurnOn() {
	fmt.Printf("%s tv turnOn\n", t.brand)
}

func (t *TV) TurnOff() {
	fmt.Printf("%s tv turnOff\n", t.brand)
}

func (t *TV) ChangeChannel(channel int) {
	fmt.Printf("%s tv channel %d to %d\n", t.brand, t.channel, channel)
	t.channel = channel
}

type Command interface {
	Execute()
}

type CommandOn struct {
	*TV
}

func NewCommandOn(tv *TV) *CommandOn {
	return &CommandOn{
		TV: tv,
	}
}

func (c *CommandOn) Execute() {
	c.TV.TurnOn()
}

type CommandOff struct {
	*TV
}

func NewCommandOff(tv *TV) *CommandOff {
	return &CommandOff{
		TV: tv,
	}
}

func (c *CommandOff) Execute() {
	c.TV.TurnOff()
}

type CommandChange struct {
	Channel int
	*TV
}

func NewCommandChange(tv *TV) *CommandChange {
	return &CommandChange{
		TV: tv,
	}
}

func (c *CommandChange) Execute() {
	c.TV.ChangeChannel(c.Channel)
}

type Control struct {
	onCommand     Command
	offCommand    Command
	changeCommand Command
}

func NewControl(on, off, change Command) *Control {
	return &Control{
		onCommand:     on,
		offCommand:    off,
		changeCommand: change,
	}
}

func (c *Control) TurnOn() {
	c.onCommand.Execute()
}

func (c *Control) TurnOff() {
	c.offCommand.Execute()
}

func (c *Control) ChangeChannel(channel int) {
	if change, ok := c.changeCommand.(*CommandChange); ok {
		change.Channel = channel
	}

	c.changeCommand.Execute()
}

func main() {
	tv := NewTV("TCL", 19)
	commandOn := NewCommandOn(tv)
	commandOff := NewCommandOff(tv)
	commandChange := NewCommandChange(tv)
	control := NewControl(commandOn, commandOff, commandChange)
	control.TurnOn()
	control.ChangeChannel(2)
	control.TurnOff()
}
