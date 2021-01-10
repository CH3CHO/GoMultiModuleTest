package service

type SayHelloService interface {
	SayHello(name string) string
}

type sayHelloServiceImpl struct {
}

func (s *sayHelloServiceImpl) SayHello(name string) string {
	return "Hello " + name + "!"
}

func init() {
}