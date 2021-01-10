package service

import "github.com/golobby/container/pkg/container"

func InitContainer(c *container.Container) {
	c.Singleton(func () SayHelloService {
		return &sayHelloServiceImpl{}
	})
}