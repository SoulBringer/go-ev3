package main

import (
	"go-ev3/serialapi"
	"github.com/gorilla/websocket"
	"fmt"
)

type WSResponse struct {
	Action string
	Index uint
	Value interface{}
}

// Response helper functionality
func writeResponse(w *websocket.Conn, request WSMsg, v interface{}, err error) {
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err == nil {
		response := WSResponse{ }
		response.Action = request.Action
		response.Index = request.Index
		response.Value = v
		if err = w.WriteJSON(response); err != nil {
			fmt.Println(err)
		}
	} else {
		w.WriteJSON(err)
	}
}

// Get port status
func GetPortsStatus(w *websocket.Conn, r WSMsg) {
	result, err := ev3.GetPortsStatus()
	writeResponse(w, r, result, err)
}

// Plays sound on brick side
func PlaySound(w *websocket.Conn, r WSMsg) {
	volume := r.Volume
	frequency := r.Frequency
	duration := r.Duration

	err := ev3.PlaySound(uint8(volume), uint16(frequency), uint16(duration))
	writeResponse(w, r, nil, err)
}

// Get color value for connected sensor
func GetColor(w *websocket.Conn, r WSMsg) {
	port := r.PortId
	result, err := ev3.GetColor(uint8(port))
	resultStr := serialapi.ColorStr(result)
	writeResponse(w, r, resultStr, err)
}

// Get luminosity value for connected sensor
func GetLuminosity(w *websocket.Conn, r WSMsg) {
	port := r.PortId
	result, err := ev3.GetLuminosity(uint8(port))
	writeResponse(w, r, result, err)
}

// Get distance value for connected sensor
func GetDistance(w *websocket.Conn, r WSMsg) {
	port := r.PortId
	result, err := ev3.GetDistance(uint8(port))
	writeResponse(w, r, result, err)
}

// Get is clickable sensor clicked
func GetIsClicked(w *websocket.Conn, r WSMsg) {
	port := r.PortId
	result, err := ev3.GetIsClicked(uint8(port))
	writeResponse(w, r, result, err)
}

// Get click count for connected sensor
func GetClickCount(w *websocket.Conn, r WSMsg) {
	port := r.PortId
	result, err := ev3.GetClickCount(uint8(port))
	writeResponse(w, r, result, err)
}

// Get gyro angle for connected sensor
func GetGyroAngle(w *websocket.Conn, r WSMsg) {
	port := r.PortId
	result, err := ev3.GetGyroAngle(uint8(port))
	writeResponse(w, r, result, err)
}

// Get gyro angle for connected sensor
func GetGyroGravity(w *websocket.Conn, r WSMsg) {
	port := r.PortId
	result, err := ev3.GetGyroGravity(uint8(port))
	writeResponse(w, r, result, err)
}

// Get value for generic sensor
func GetSensorValue(w *websocket.Conn, r WSMsg) {
	port := r.PortId
	result, err := ev3.GetSensorValue(uint8(port), 0xFF)
	writeResponse(w, r, result, err)
}

// Start motor
func MotorStart(w *websocket.Conn, r WSMsg) {
	port := r.PortId
	err := ev3.MoveMotorStart(uint8(port))
	writeResponse(w, r,nil, err)
}

// Stop motor and (not) apply brake
func MotorStop(w *websocket.Conn, r WSMsg) {
	port := r.PortId
	brake := r.Brake
	err := ev3.MoveMotorStop(uint8(port), uint8(brake))
	writeResponse(w, r,nil, err)
}

// Set motor speed
func SetMotorSpeed(w *websocket.Conn, r WSMsg) {
	port := r.PortId
	speed := r.Speed
	err := ev3.MoveMotorSpeed(uint8(port), int8(speed))
	writeResponse(w, r,nil, err)
}

// Get motor current angle
func GetMotorState(w *websocket.Conn, r WSMsg) {
	port := r.PortId
	result, err := ev3.GetMotorAngle(uint8(port))
	writeResponse(w, r, result, err)
}
