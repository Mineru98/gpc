package node

import (
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

func SlaveMode(slave Slave, ip string, port string) {
	conn, err := net.Dial("tcp", ip+":"+port)
	if nil != err {
		log.Fatalf("failed to connect to server")
	}

	for {
		conn.Write([]byte("ping from " + ip + ":" + port))
		time.Sleep(time.Duration(3) * time.Second)
	}
}

func init() {

}
