package main

// run command = tinygo flash -target arduino -port /dev/ttyACM0 github.com/SkylerBair/knight-rider

import (
	"machine"
	"time"
)

func main() {

	centerSensor := machine.ADC{Pin: machine.ADC0}
	centerSensor.Configure(machine.ADCConfig{})

	leftSensor := machine.ADC{Pin: machine.ADC1}
	leftSensor.Configure(machine.ADCConfig{})

	rightSensor := machine.ADC{Pin: machine.ADC2}
	rightSensor.Configure(machine.ADCConfig{})

	machine.InitADC()

	motorSpeed()

	delay := func(t uint16) {
		time.Sleep(time.Duration(1000000 * uint32(t)))
	}

	for {

		print(leftSensor.Get())
		print(centerSensor.Get())
		print(rightSensor.Get())

		if leftSensor.Get() < 6000 {
			sensorTest()
			reverse()
			delay(1000)
			deviateRight()
			delay(2000)

		} else if rightSensor.Get() < 6000 {
			sensorTest2()
			reverse()
			delay(1000)
			deviateLeft()
			delay(2000)
		} else if centerSensor.Get() < 6000 {
			sensorTest3()
			reverse()
			delay(100)
			turnAround()
			delay(3000)
		} else if leftSensor.Get()&centerSensor.Get() < 6000 {
			reverse()
			delay(1000)
			deviateRight()
			delay(2000)

		} else if rightSensor.Get()&centerSensor.Get() < 6000 {
			reverse()
			delay(1000)
			deviateLeft()
			delay(2000)

		} else if leftSensor.Get()&rightSensor.Get()&centerSensor.Get() < 6000 {
			reverse()
			delay(1000)
			turnAround()
			delay(3000)
		} else {
			forward()
		}
	}

}

// the pin config shit needs to be added as struct then used when function is called, set up in main.
func sensorTest() {

	lightCheck := machine.D2
	lightCheck.Configure(machine.PinConfig{Mode: machine.PinOutput})
	lightCheck.Set(true)
	delay := func(t uint16) {
		time.Sleep(time.Duration(1000000 * uint32(t)))
	}
	delay(500)
	lightCheck.Set(false)
}

func sensorTest2() {

	lightCheck := machine.D3
	lightCheck.Configure(machine.PinConfig{Mode: machine.PinOutput})
	lightCheck.Set(true)
	delay := func(t uint16) {
		time.Sleep(time.Duration(1000000 * uint32(t)))
	}
	delay(500)
	lightCheck.Set(false)
}
func sensorTest3() {

	lightCheck := machine.D4
	lightCheck.Configure(machine.PinConfig{Mode: machine.PinOutput})
	lightCheck.Set(true)
	delay := func(t uint16) {
		time.Sleep(time.Duration(1000000 * uint32(t)))
	}
	delay(500)
	lightCheck.Set(false)
}

func motorSpeed() {

	pwm := machine.Timer1
	rightMotorSpeed := machine.D5

	err := pwm.Configure(machine.PWMConfig{})
	if err != nil {
		println(err.Error())
	}

	ch, err := pwm.Channel(rightMotorSpeed)
	if err != nil {
		println(err.Error())
	}

	pwm.Set(ch, pwm.Top()/4)

	pwm2 := machine.Timer1
	leftMotorSpeed := machine.D9

	err2 := pwm2.Configure(machine.PWMConfig{})
	if err != nil {
		println(err.Error())
	}

	ch2, err2 := pwm2.Channel(leftMotorSpeed)
	if err != nil {
		println(err2.Error())
	}

	pwm2.Set(ch2, pwm.Top()/4)

}

func deviateRight() {
	//pin Config's
	//motorSpeed(250, 250)
	leftMotorReverse := machine.D6
	leftMotorReverse.Configure(machine.PinConfig{Mode: machine.PinOutput})

	leftMotorForward := machine.D10
	leftMotorForward.Configure(machine.PinConfig{Mode: machine.PinOutput})

	rightMotorForward := machine.D13
	rightMotorForward.Configure(machine.PinConfig{Mode: machine.PinOutput})

	rightMotorReverse := machine.D11
	rightMotorReverse.Configure(machine.PinConfig{Mode: machine.PinOutput})
	//motor set up

	//motorSpeed(250, 250)
	leftMotorForward.High()
	rightMotorForward.Low()
	leftMotorReverse.Low()
	rightMotorReverse.Low()

}

func deviateLeft() {
	//pin Config's
	leftMotorReverse := machine.D6
	leftMotorReverse.Configure(machine.PinConfig{Mode: machine.PinOutput})

	leftMotorForward := machine.D10
	leftMotorForward.Configure(machine.PinConfig{Mode: machine.PinOutput})

	rightMotorForward := machine.D13
	rightMotorForward.Configure(machine.PinConfig{Mode: machine.PinOutput})

	rightMotorReverse := machine.D11
	rightMotorReverse.Configure(machine.PinConfig{Mode: machine.PinOutput})
	//motor set up

	//motorSpeed(250, 250)
	leftMotorForward.Low()
	rightMotorForward.High()
	leftMotorReverse.Low()
	rightMotorReverse.Low()

}

func forward() {
	//pin Config's
	leftMotorReverse := machine.D6
	leftMotorReverse.Configure(machine.PinConfig{Mode: machine.PinOutput})

	leftMotorForward := machine.D10
	leftMotorForward.Configure(machine.PinConfig{Mode: machine.PinOutput})

	rightMotorForward := machine.D13
	rightMotorForward.Configure(machine.PinConfig{Mode: machine.PinOutput})

	rightMotorReverse := machine.D11
	rightMotorReverse.Configure(machine.PinConfig{Mode: machine.PinOutput})
	//motor set up
	//motorSpeed(250, 250)
	leftMotorForward.High()
	rightMotorForward.High()
	leftMotorReverse.Low()
	rightMotorReverse.Low()

}

func reverse() {
	//pin Config's
	leftMotorReverse := machine.D6
	leftMotorReverse.Configure(machine.PinConfig{Mode: machine.PinOutput})

	leftMotorForward := machine.D10
	leftMotorForward.Configure(machine.PinConfig{Mode: machine.PinOutput})

	rightMotorForward := machine.D13
	rightMotorForward.Configure(machine.PinConfig{Mode: machine.PinOutput})

	rightMotorReverse := machine.D11
	rightMotorReverse.Configure(machine.PinConfig{Mode: machine.PinOutput})
	//motor set up
	//motorSpeed(250, 250)
	rightMotorForward.Low()
	rightMotorReverse.High()
	leftMotorForward.Low()
	leftMotorReverse.High()

}

func turnAround() {
	//pin Config's
	leftMotorReverse := machine.D6
	leftMotorReverse.Configure(machine.PinConfig{Mode: machine.PinOutput})

	leftMotorForward := machine.D10
	leftMotorForward.Configure(machine.PinConfig{Mode: machine.PinOutput})

	rightMotorForward := machine.D13
	rightMotorForward.Configure(machine.PinConfig{Mode: machine.PinOutput})

	rightMotorReverse := machine.D11
	rightMotorReverse.Configure(machine.PinConfig{Mode: machine.PinOutput})
	//motor set up

	//motorSpeed(250, 250)
	leftMotorForward.Low()
	rightMotorForward.High()
	leftMotorReverse.High()
	rightMotorReverse.Low()

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
