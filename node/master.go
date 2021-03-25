package node

import (
	"fmt"
	"log"
	"net"

	"../utils/network"
)

type MasterNode interface {
	Mode(Master, string, string)
}

type Master struct {
	ip   string
	port string
}

func MasterMode(master Master, ip string, port string) {
	fmt.Println("Started GPC Master Mode Service")
	master.ip = ip
	master.port = port
	l, err := net.Listen("tcp", ip+":"+port)
	if nil != err {
		log.Fatalf("fail to bind address to %v; err: %v", port, err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if nil != err {
			log.Printf("fail to accept; err: %v", err)
			continue
		}

		go network.ConnHandler(conn)
	}
}

func init() {
}
