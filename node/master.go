package master

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
	id   int
	ip   string
	port string
}

func Mode(master Master, ip string, port string) {
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
	fmt.Println("Master Node Init")
}
