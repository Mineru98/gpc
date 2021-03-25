package node

import (
	"fmt"
	"log"
	"net"
	"time"
)

type SlaveNode interface {
	Mode(Slave, string, string)
}

type Slave struct {
	id   int
	ip   string
	port string
}

func SlaveMode(slave Slave, ip, port, t_ip, t_port string) {
	fmt.Println("Started GPC Slave Mode Service")
	slave.ip = ip
	slave.port = port
	conn, err := net.Dial("tcp", t_ip+":"+t_port)
	if nil != err {
		log.Fatalf("failed to connect to server")
	}
	fmt.Println(ip)
	for {
		conn.Write([]byte("ping from " + ip + ":" + port))
		time.Sleep(time.Duration(3) * time.Second)
	}
}

func init() {
}
