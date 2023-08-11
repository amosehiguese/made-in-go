package main

import (
	"context"
	"io"
	"log"
	"net"
	"time"

	pb "github.com/amosehiguese/grpc-app/proto"
	"google.golang.org/grpc"
)

//define the port
const (
	port = ":8089"
)

//this is the struct to be created, pb is imported upstairs
type helloServer struct {
	pb.UnimplementedGreetServiceServer
	// pb.GreetServiceServer
}

// simple rpc
func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}

// server streaming rpc
func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("got request with names: %v", req.Names)

	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil{
			return err
		}
		time.Sleep(2 * time.Second)
	}

	return nil
}

// client streaming rpc
func (s *helloServer) SayhelloClientStreaming(stream pb.GreetService_SayhelloClientStreamingServer) error {
	var names []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			messages := pb.MessagesList{
				Messages: names,
			}
			return stream.SendAndClose(&messages)
		}

		if err != nil {
			return err
		}

		log.Printf("Got request with name: %v", req.Name)

		names = append(names, "Hello " + req.Name)
	}
}

// bi-directional streaming rpc
func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("Got request with name: %v", req.Name)

		resp := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}

		if err := stream.Send(resp); err != nil {
			return err
		}

	}
}

func main() {
	//listen on the port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
	// create a new gRPC server
	grpcServer := grpc.NewServer()

	// register the greet service
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Server started at %v", lis.Addr())

	//list is the port, the grpc server needs to start there
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
