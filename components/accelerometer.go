package components

import (
	"github.com/babymechanic/motecommon/messages"
)

type Accelerometer struct {
}

func (accelerometer *Accelerometer) Heading(_ struct{}, heading *messages.Heading) error {
	*heading = messages.Heading{Angle: 90.9}
	return nil
}

func init() {
	registeredComponents.Register("Accelerometer", &Accelerometer{})
}
