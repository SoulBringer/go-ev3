package main

import (
	"fmt"
	"go-ev3/serialapi"
)

// Demo app main entry point
func main() {
	fmt.Println("EV3 API demo")

	ev3 := serialapi.EV3{
		PortName: "/dev/rfcomm2",
		PortBaud: 9600,
		DebugOn:  true,
	}
	ev3.PlaySound(2, 1000, 1000)
	//ev3.GetPortsStatus()
	//ev3.GetColorValue(serialapi.SensorPort1)
}
