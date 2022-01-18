package main

import (
	"machine"
	"time"
)

func main() {

	delay := func(t uint16) {
		time.Sleep(time.Duration(1000000 * uint32(t)))
	}

	leds := []machine.Pin{
		machine.D2,
		machine.D3,
		machine.D4,
		machine.D5,
		machine.D6,
		machine.D7,
		machine.D8,
	}

	for _, led := range leds {
		led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	index, delta := 0, 1

	for {

		for i, led := range leds {
			led.Set(i == index)
		}

		index += delta

		if index == 0 || index == len(leds)-1 {
			delta *= -1
		}

		delay(75)
	}
}
