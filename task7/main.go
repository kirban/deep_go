package main

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrWrongConstructor = errors.New("wrong constructor")
)

type Container struct {
	items map[string]any
}

func NewContainer() *Container {
	return &Container{
		items: make(map[string]any),
	}
}

func (c *Container) RegisterType(name string, constructor interface{}) {
	con, err := assertToConstructor(constructor)

	if err != nil {
		panic(err)
	}

	c.items[name] = con
}

func (c *Container) Resolve(name string) (interface{}, error) {
	val, ok := c.items[name]

	if !ok {
		return nil, fmt.Errorf("resolve error: item %s %w", name, ErrNotFound)
	}

	con, err := assertToConstructor(val)

	if err != nil {
		return nil, fmt.Errorf("resolve error: %w", ErrWrongConstructor)
	}

	return con(), nil
}

func assertToConstructor(value interface{}) (func() any, error) {
	con, ok := value.(func() any)
	if !ok {
		return nil, ErrWrongConstructor
	}
	return con, nil
}
