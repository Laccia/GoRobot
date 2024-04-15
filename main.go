package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/drivers/gpio"
	"gobot.io/x/gobot/v2/platforms/adaptors"
	"gobot.io/x/gobot/v2/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	led := gpio.NewLedDriver(r, "33")
	s := raspi.NewAdaptor(adaptors.WithPWMDefaultPeriod(30000000), adaptors.WithPWMUsePiBlaster())
	//adaptors.WithPWMDefaultPeriodForPin("12", 2000000), adaptors.WithPWMServoAngleRangeForPin("12", 90, 180))
	sr := gpio.NewServoDriver(s, "pwm0")
	s.ServoWrite("pwm0", 90)
	work := func() {
		sr.Move(90)
		s.ServoWrite("pwm0", 180)
		fmt.Println(s)
		gobot.Every(3*time.Second, func() {
			led.Toggle()
			s.ServoWrite("pwm0", 180)
			sr.Move(90)
			sr.ToCenter()

		})

	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Connection{s},
		[]gobot.Device{led},
		[]gobot.Device{sr},
		[]gobot.Adaptor{s},
		s.ServoWrite("PWM0", 180),
		sr.Move(90),
		sr.ToCenter(),

		work,
	)

	robot.Start()
}
