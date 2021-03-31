package node

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "../utils/define"
	"google.golang.org/grpc"
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
	// conn, err := net.Dial("tcp", t_ip+":"+t_port)
	conn, err := grpc.Dial(t_ip+":"+t_port, grpc.WithInsecure(), grpc.WithBlock())
	if nil != err {
		log.Fatalf("failed to connect to server")
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	name := "gpc"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
	// fmt.Println(ip)
	// for {
	// 	conn.Write([]byte("ping from " + ip + ":" + port))
	// 	time.Sleep(time.Duration(3) * time.Second)
	// }
}

func init() {
}
