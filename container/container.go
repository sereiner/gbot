package container

import "github.com/codegangsta/inject"

type Container struct {
	inject.Injector
}

func NewContainer() *Container {
	return &Container{
		Injector: inject.New(),
	}
}

func (c *Container) Register(v interface{}) {
	c.Map(v)
}
