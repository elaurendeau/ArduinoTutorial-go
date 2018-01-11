package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/aio"
	"gobot.io/x/gobot/platforms/firmata"
	"time"
	"gobot.io/x/gobot/drivers/gpio"
)

func main() {

	firmataAdapter := firmata.NewAdaptor("COM3")
	lightSensor := aio.NewAnalogSensorDriver(firmataAdapter, "3", 50*time.Millisecond)
	led := gpio.NewLedDriver(firmataAdapter, "11")

	work := func() {

		lightSensor.On(lightSensor.Event("data"), func(data interface{}) {
			brightness := uint8(
				gobot.ToScale(gobot.FromScale(float64(data.(int)), 0, 1024), 0, 255),
			)

			if (brightness < 100) {
				led.Brightness(250-brightness*2)
			} else {
				led.Off()
			}
		})

	}
	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdapter},
		[]gobot.Device{lightSensor, led},
		work)

	robot.Start()
}
