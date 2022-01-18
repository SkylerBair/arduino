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

	machine.InitADC()
	ldr := machine.ADC{machine.ADC0}
	ldr.Configure()

	led := machine.Pin(machine.D9)
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {

		print(ldr.Get())

		if ldr.Get() > 40000 {
			led.Set(true)
		} else {
			led.Set(false)
		}

		time.Sleep(time.Millisecond * 100)
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
