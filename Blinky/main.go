package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutout})
	for {
		led.High()
		time.sleep(time.Millisecond * 500)
		led.Low()
		time.sleep(time.Millisecond * 500)
	}
}
