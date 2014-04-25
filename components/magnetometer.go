package components

import (
	"github.com/babymechanic/motecommon/messages"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/all"
	"github.com/kidoman/embd/sensor/lsm303"
)

type Magnetometer struct {
	bus            embd.I2CBus
	magnet         *lsm303.LSM303
	wasInitialized bool
}

func (magnetometer *Magnetometer) Heading(_ struct{}, heading *messages.Heading) error {
	magnetometer.initialize()
	headingValue, err := magnetometer.magnet.Heading()
	if err != nil {
		panic(err)
	}
	*heading = messages.Heading{Angle: headingValue}
	return nil
}

func (magnetometer *Magnetometer) initialize() {
	if magnetometer.wasInitialized {
		return
	}
	if err := embd.InitI2C(); err != nil {
		panic(err)
	}
	magnetometer.bus = embd.NewI2CBus(1)
	magnetometer.magnet = lsm303.New(magnetometer.bus)
	magnetometer.magnet.Run()
	magnetometer.wasInitialized = true
}

func (magnetometer *Magnetometer) close() {
	magnetometer.magnet.Close()
	magnetometer.bus.Close()
}

func init() {
	registeredComponents.Register("Magnetometer", &Magnetometer{})
}
