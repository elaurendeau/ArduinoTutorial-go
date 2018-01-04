package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main()  {

	firmataAdapter := firmata.NewAdaptor("COM3")
	led := gpio.NewLedDriver(firmataAdapter, "13")
	button := gpio.NewButtonDriver(firmataAdapter, "8")


	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			led.On()
			println("button pressed")
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			led.Off()
			println("button released")
		})

	}
	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdapter},
		[]gobot.Device{button, led},
		work)

	robot.Start()
	
	
}