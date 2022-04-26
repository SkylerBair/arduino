package main

// run command = tinygo flash -target arduino -port /dev/ttyACM0 github.com/SkylerBair/knight-rider

import (
	"machine"
	"time"
)

func main() {
	delay := func(t uint16) {
		time.Sleep(time.Duration(1000000 * uint32(t)))
	}

	delay(100)
}

// the pin config shit needs to be added as struct then used when function is called, set up in main.

func deviateRight() {
	//pin Config's
	leftMotorForward := machine.D6
	leftMotorForward.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leftMotorReverse := machine.D10 //P9 is the Arduino output 9 that inits the Right motor to turn forward.
	leftMotorReverse.Configure(machine.PinConfig{Mode: machine.PinOutput})
	rightMotorForward := machine.D9
	rightMotorForward.Configure(machine.PinConfig{Mode: machine.PinOutput})
	rightMotorReverse := machine.D11 //P9 is the Arduino output 9 that inits the Right motor to turn forward.
	leftMotorReverse.Configure(machine.PinConfig{Mode: machine.PinOutput})
	//motor set up
	leftMotorForward.High()
	rightMotorForward.Low()
	leftMotorReverse.Low()
	rightMotorReverse.Low()

}

func deviateLeft() {
	//pin Config's
	leftMotorForward := machine.D6
	leftMotorForward.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leftMotorReverse := machine.D10 //P9 is the Arduino output 9 that inits the Right motor to turn forward.
	leftMotorReverse.Configure(machine.PinConfig{Mode: machine.PinOutput})
	rightMotorForward := machine.D9
	rightMotorForward.Configure(machine.PinConfig{Mode: machine.PinOutput})
	rightMotorReverse := machine.D11 //P9 is the Arduino output 9 that inits the Right motor to turn forward.
	leftMotorReverse.Configure(machine.PinConfig{Mode: machine.PinOutput})
	//motor set up
	reverse()
	leftMotorForward.Low()
	rightMotorForward.High()
	leftMotorReverse.High()
	rightMotorReverse.Low()

}

func forward() {
	//pin Config's
	leftMotorForward := machine.D6
	leftMotorForward.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leftMotorReverse := machine.D10 //P9 is the Arduino output 9 that inits the Right motor to turn forward.
	leftMotorReverse.Configure(machine.PinConfig{Mode: machine.PinOutput})
	rightMotorForward := machine.D9
	rightMotorForward.Configure(machine.PinConfig{Mode: machine.PinOutput})
	rightMotorReverse := machine.D11 //P9 is the Arduino output 9 that inits the Right motor to turn forward.
	leftMotorReverse.Configure(machine.PinConfig{Mode: machine.PinOutput})
	//motor set up
	reverse()
	leftMotorForward.High()
	rightMotorForward.High()
	leftMotorReverse.Low()
	rightMotorReverse.Low()

}

func reverse() {
	//pin Config's
	leftMotorForward := machine.D6
	leftMotorForward.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leftMotorReverse := machine.D10 //P9 is the Arduino output 9 that inits the Right motor to turn forward.
	leftMotorReverse.Configure(machine.PinConfig{Mode: machine.PinOutput})
	rightMotorForward := machine.D9
	rightMotorForward.Configure(machine.PinConfig{Mode: machine.PinOutput})
	rightMotorReverse := machine.D11 //P9 is the Arduino output 9 that inits the Right motor to turn forward.
	leftMotorReverse.Configure(machine.PinConfig{Mode: machine.PinOutput})
	//motor set up
	leftMotorForward.Low()
	rightMotorForward.Low()
	leftMotorReverse.High()
	rightMotorReverse.High()

}

//             *Knight-rider Example*
// func main() {

// 	delay := func(t uint16) {
// 		time.Sleep(time.Duration(1000000 * uint32(t)))
// 	}

// 	leds := []machine.Pin{
// 		machine.D2,
// 		machine.D3,
// 		machine.D4,
// 		machine.D5,
// 		machine.D6,
// 		machine.D7,
// 		machine.D8,
// 	}

// 	for _, led := range leds {
// 		led.Configure(machine.PinConfig{Mode: machine.PinOutput})
// 	}

// 	index, delta := 0, 1

// 	for {

// 		for i, led := range leds {
// 			led.Set(i == index)
// 		}

// 		index += delta

// 		if index == 0 || index == len(leds)-1 {
// 			delta *= -1
// 		}

// 		delay(200)
// 	}
// }
