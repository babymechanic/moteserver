package components

import (
	"github.com/babymechanic/motecommon/messages"
	"github.com/kidoman/embd"
	"github.com/kidoman/embd/sensor/l3gd20"

	_ "github.com/kidoman/embd/host/all"
)

type Gyroscope struct {
	bus            embd.I2CBus
	gyro           *l3gd20.L3GD20
	wasInitialized bool
}

func (gyroscope *Gyroscope) Orientation(_ struct{}, orientation *messages.Orientation) error {
	gyroscope.initialize()
	*orientation = gyroscope.orientation()
	return nil
}

func (gyroscope *Gyroscope) orientation() messages.Orientation {
	orientations, err := gyroscope.gyro.Orientations()
	if err != nil {
		panic(err)
	}
	orientation := <-orientations
	return messages.Orientation{X: orientation.X, Y: orientation.Y, Z: orientation.Z}
}

func (gyroscope *Gyroscope) initialize() {
	if !gyroscope.wasInitialized {
		return
	}
	if err := embd.InitI2C(); err != nil {
		panic(err)
	}
	gyroscope.bus = embd.NewI2CBus(1)
	gyroscope.gyro = l3gd20.New(gyroscope.bus, l3gd20.R250DPS)
	gyroscope.gyro.Start()
	gyroscope.wasInitialized = true
}

func (gyroscope *Gyroscope) Destroy() {
	gyroscope.gyro.Close()
	gyroscope.bus.Close()
	embd.CloseI2C()
}

func init() {
	registeredComponents.Register("Gyroscope", &Gyroscope{})
}
