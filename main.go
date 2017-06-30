package main

import (
	"fmt"
	"go-ev3/serialapi"
	"log"
	"net/http"
)

var ev3 serialapi.EV3

// Demo app main entry point
func main() {
	// Run REST API
	// TODO: Read values from config/input params
	initRestApi(8081, "/dev/rfcomm0", 9600)

	// Execute DEMO code
	// runDemo("/dev/rfcomm8", 9600)
}

// Initializes EV3 interface REST API
func initRestApi(httpPort int, serialPort string, serialBaud int) {
	ev3 = serialapi.EV3{
		PortName: serialPort,
		PortBaud: serialBaud,
		DebugOn:  false,
	}

	router := NewRouter()
	log.Println("Go EV3 REST API server initialized")
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", httpPort), router))
}

// Execute demo code
//func runDemo(serialPort string, serialBaud int) {
//	ev3 := serialapi.EV3{
//		PortName: serialPort,
//		PortBaud: serialBaud,
//		DebugOn:  false,
//	}
//
//	portsStatus, _ := ev3.GetPortsStatus()
//	fmt.Printf("Port status:\n%v\n", portsStatus)
//
//	val, _ := ev3.GetDistance(serialapi.SensorPort1)
//	fmt.Printf("Distance: %d cm\n", val)
//
//	val, _ = ev3.GetClickCount(serialapi.SensorPort2)
//	fmt.Printf("Click count: %d\n", val)
//
//	val, _ = ev3.GetIsClicked(serialapi.SensorPort2)
//	fmt.Printf("Is clicked: %d\n", val)
//
//	val, _ = ev3.GetColor(serialapi.SensorPort3)
//	fmt.Printf("Color: %s\n", serialapi.ColorStr(val))
//
//	val, _ = ev3.GetLuminosity(serialapi.SensorPort3)
//	fmt.Printf("Luminosity: %d\n", val)
//
//	val, _ = ev3.GetGyroAngle(serialapi.SensorPort4)
//	fmt.Printf("Gyro angle: %d\n", val)
//
//	val, _ = ev3.GetGyroGravity(serialapi.SensorPort4)
//	fmt.Printf("Gyro gravity: %d\n", val)
//
//	ev3.MoveMotorAngle(serialapi.MotorPortA, 50, 45, 0)
//	time.Sleep(2 * time.Second)
//
//	ev3.MoveMotorSpeed(serialapi.MotorPortB, 0)
//	ev3.MoveMotorStart(serialapi.MotorPortB)
//	ev3.MoveMotorSpeed(serialapi.MotorPortB, 100)
//	time.Sleep(1 * time.Second)
//
//	ev3.MoveMotorSpeed(serialapi.MotorPortB, 30)
//	time.Sleep(1 * time.Second)
//
//	ev3.MoveMotorSpeed(serialapi.MotorPortB, 0)
//	ev3.MoveMotorStop(serialapi.MotorPortB, 0)
//	time.Sleep(2 * time.Second)
//
//	ev3.PlaySound(2, 1000, 200)
//}
