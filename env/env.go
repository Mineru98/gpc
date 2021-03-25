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
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				Ip = ipnet.IP.String()
			}
		}
	}
}
