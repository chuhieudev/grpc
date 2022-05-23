package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/chuhieudev/grpc.git/calculatorpb"
)

type server struct()

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8008")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s:=grpc.NewServer()

	calculatorpb.RegisterCalculatorServer(s, &server{})

	fmt.Println("runing...")

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
