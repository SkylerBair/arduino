package main

import (
	"machine"
	"time"
)

/*
motor1 = AFMotor //motor 1 is the drive motor
motor2 = AFMotor //motor 2 is the turning servo motor
var val int
bt = serial.read
motor1.setSpeed(255)
motor2.setspeed(225)
serial.begin(9600)
*/
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

/*
if bt == "1" {
	motor1 = forward
}

func forward() int {
	motor1.run(forward)
	return
}

func movebackward() string {
	motor1 := 2
	return
}

func turnleft() string {
	motor2 := 3
	return
}

func turnright() string {
	motor2 := 4
	return
}

func stopcar() string {
	motor1 := 0
	return

}
*/
