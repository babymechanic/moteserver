package components

import (
	"fmt"
	"net/rpc"
)

type RegisteredComponents map[string]interface{}

var registeredComponents = RegisteredComponents{}

func (registeredComponents RegisteredComponents) Register(name string, component interface{}) {
	registeredComponents[name] = component
	rpc.Register(component)
	fmt.Println("registered component", name)
}
