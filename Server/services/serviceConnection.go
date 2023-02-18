package services

type connectionWithParams[Arg any, Return any] struct {
	handler func(Arg) Return
}

func (c *connectionWithParams[Arg, Return]) Provide(handler func(Arg) Return) {
	c.handler = handler
}

func (c *connectionWithParams[Arg, Return]) Request(arg Arg) Return {
	return c.handler(arg)
}

type ServiceConnection struct {
	ConnectionIDByUserEmail connectionWithParams[string, string]
}

func NewServiceConnection() *ServiceConnection {
	return &ServiceConnection{connectionWithParams[string, string]{}}
}
