package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/validator.v2"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/ws", wsHandler)

	return router
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	//if r.Header.Get("Origin") != "http://"+r.Host {
	//	http.Error(w, "Origin not allowed", 403)
	//	return
	//}
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}

	go processWsMsg(conn)
}

type WSMsg struct {
	Action string
	Index uint

	Volume uint8 `validate:"min=0,max=100"`
	Frequency uint16 `validate:"min=0,max=65535"`
	Duration uint16 `validate:"min=0,max=65535"`

	PortId uint8 `validate:"min=0,max=255"`
	Speed int8 `validate:"min=0,max=255"`
	Brake uint8 `validate:"min=0,max=1"`
}

func processWsMsg(conn *websocket.Conn) {
	var index uint
	for {
		m := WSMsg{}

		err := conn.ReadJSON(&m)
		if err != nil {
			fmt.Println("Error reading json.", err)
			break
		}

		if errs := validator.Validate(m); errs != nil {
			fmt.Println("Validation failed", err)
			continue
		}

		index++
		m.Index = index

		fmt.Printf("Got message: %#v\n", m)

		switch m.Action {
			case "GetPortsStatus": 	GetPortsStatus(conn, m)
			case "PlaySound": 	PlaySound(conn, m)
			case "MotorStart": 	MotorStart(conn, m)
			case "MotorStop": 	MotorStop(conn, m)
			case "SetMotorSpeed": 	SetMotorSpeed(conn, m)
			case "GetMotorState": 	GetMotorState(conn, m)
			case "GetColor": 	GetColor(conn, m)
			case "GetLuminosity": 	GetLuminosity(conn, m)
			case "GetDistance": 	GetDistance(conn, m)
			case "GetIsClicked": 	GetIsClicked(conn, m)
			case "GetClickCount": 	GetClickCount(conn, m)
			case "GetGyroAngle": 	GetGyroAngle(conn, m)
			case "GetGyroGravity": 	GetGyroGravity(conn, m)
			case "GetSensorValue": 	GetSensorValue(conn, m)
		}
	}
}
