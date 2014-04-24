package components

import (
	"github.com/babymechanic/motecommon/messages"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/all"
	"github.com/kidoman/embd/sensor/lsm303"
)

type Accelerometer struct {
	bus            embd.I2CBus
	accel          *lsm303.LSM303
	wasInitialized bool
}

func (accelerometer *Accelerometer) Heading(_ struct{}, heading *messages.Heading) error {
	accelerometer.initialize()
	headingValue, err := accelerometer.accel.Heading()
	if err != nil {
		panic(err)
	}
	*heading = messages.Heading{Angle: headingValue}
	return err
}

func (accelerometer *Accelerometer) initialize() {
	if accelerometer.wasInitialized {
		return
	}
	if err := embd.InitI2C(); err != nil {
		panic(err)
	}
	accelerometer.bus = embd.NewI2CBus(1)
	accelerometer.accel = lsm303.New(accelerometer.bus)
	accelerometer.accel.Run()
	accelerometer.wasInitialized = true
}

func (accelerometer *Accelerometer) close() {
	accelerometer.accel.Close()
	accelerometer.bus.Close()
}

func init() {
	registeredComponents.Register("Accelerometer", &Accelerometer{})
}
