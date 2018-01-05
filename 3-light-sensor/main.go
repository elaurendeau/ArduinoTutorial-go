package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/aio"
	"gobot.io/x/gobot/platforms/firmata"
	"time"
)

func main() {

	firmataAdapter := firmata.NewAdaptor("COM3")
	lightSensor := aio.NewAnalogSensorDriver(firmataAdapter, "0", 1*time.Second)


	work := func() {

		lightSensor.On(lightSensor.Event("data"), func(data interface{}) {
			brightness := uint8(
				gobot.ToScale(gobot.FromScale(float64(data.(int)), 0, 1024), 0, 255),
			)
			println("Brightness: ", brightness)
		})

	}
	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdapter},
		[]gobot.Device{lightSensor},
		work)

	robot.Start()
}
