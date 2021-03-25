package network

import (
	"io"
	"log"
	"net"
	"time"
)

func ConnHandler(conn net.Conn) {
	recvBuf := make([]byte, 4096) // receive buffer: 4kb
	for {
		n, err := conn.Read(recvBuf)
		if nil != err {
			if io.EOF == err {
				log.Printf("connection is closed from client; %v", conn.RemoteAddr().String())
				return
			}
			log.Printf("fail to receive data; err: %v", err)
			return
		}
		if 0 < n {
			data := recvBuf[:n]
			log.Println(string(data))
		}
	}
}

func Client(connInfo string) {
	conn, err := net.Dial("tcp", connInfo)
	if nil != err {
		log.Fatalf("failed to connect to server")
	}

	for {
		conn.Write([]byte("ping"))
		time.Sleep(time.Duration(3) * time.Second)
	}
}
