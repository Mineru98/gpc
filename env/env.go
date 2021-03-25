package env

import (
	"net"
	"os"
)

// Currrent Build Info
const BuildVersion = "0.0"
const BuildCode = "1"

// Debugging mode
var Debug bool = false

// Current Ip
var Ip string = "127.0.0.1"

// GPC Master Mode Port
var PortM string = "21324"

// GPC Slave Mode Port
var PortS string = "21325"

func init() {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	Ip = localAddr.IP.String()
}
