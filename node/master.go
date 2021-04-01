package node

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "../utils/proto"
	"google.golang.org/grpc"
)

type MasterNode interface {
	Mode(Master, string, string)
}

type Master struct {
	ip   string
	port string
}

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func MasterMode(master Master, ip string, port string) {
	fmt.Println("Started GPC Master Mode Service")
	master.ip = ip
	master.port = port
	lis, err := net.Listen("tcp", ip+":"+port)
	if nil != err {
		log.Fatalf("fail to bind address to %v; err: %v", port, err)
	}
	s := grpc.NewServer()
	pb.RegisterGPCServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func init() {
}
