package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {

	firmataAdapter := firmata.NewAdaptor("COM3")
	led := gpio.NewLedDriver(firmataAdapter, "13")
	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
			println("Light toggle")
		})
	}
	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdapter},
		[]gobot.Device{led},
		work)

	robot.Start()
}
