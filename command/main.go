package main

import (
	"fmt"
)

// Command は処理内容を表すインタフェースです
type Command interface {
	Execute()
}

// Source はコマンドの処理対象です
type Source struct {
	name string
}

// NewSource is XXX
func NewSource(name string) *Source {
	return &Source{
		name: name,
	}
}

// GetName is XXX
func (s *Source) GetName() string {
	return s.name
}

// CommandA is XXX
type CommandA struct {
	source *Source
}

// NewCommandA is XXX
func NewCommandA(s *Source) *CommandA {
	return &CommandA{
		source: s,
	}
}

// Execute is XXX
func (c *CommandA) Execute() {
	fmt.Println(c.source.GetName() + " by CommandA")
}

// CommandB is XXX
type CommandB struct {
	source *Source
}

// NewCommandB is XXX
func NewCommandB(s *Source) *CommandB {
	return &CommandB{
		source: s,
	}
}

// Execute is XXX
func (c *CommandB) Execute() {
	fmt.Println(c.source.GetName() + " by CommandB")
}

// CommandManager は複数のコマンドを管理します
type CommandManager struct {
	commands []Command
}

// NewCommandManager is XXX
func NewCommandManager(commands ...Command) *CommandManager {
	return &CommandManager{
		commands: commands,
	}
}

// Run is XXX
func (c *CommandManager) Run() {
	for _, command := range c.commands {
		command.Execute()
	}
}

func main() {
	source := NewSource("source1")
	cs := []Command{
		NewCommandA(source),
		NewCommandB(source),
	}

	NewCommandManager(cs...).Run()
}
