package components

import (
	"fmt"
	"github.com/babymechanic/motecommon/messages"
)

type Display struct {
}

func (display *Display) Resolution(_ struct{}, resolution *messages.Resolution) error {
	*resolution = messages.Resolution{Width: 1920, Height: 1080}
	return nil
}

func (display *Display) SetResolution(resolution messages.Resolution, _ *struct{}) error {
	fmt.Println("setting rez as ", resolution.Height, " x ", resolution.Width)
	return nil
}

func init() {
	registeredComponents.Register("Display", &Display{})
}
