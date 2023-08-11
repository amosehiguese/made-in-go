package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/amosehiguese/grpc-app/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8089"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}


	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Amos", "Alice", "Bob", "Mike", "Sam", "Edith", "Jack", "Hue", "Neyo"},
	}

	// CallSayHello(client)
	// CallSayHelloServerStreaminig(client, names)
	// CallSayHelloClientStream(client, names)
	CallSayHelloBidirectionalStream(client, names)
}

func CallSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("%s", res.Message)
}

func CallSayHelloServerStreaminig(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Streaming started")
	ctx := context.Background()
	// defer cancel()

	stream, err := client.SayHelloServerStreaming(ctx, names)
	if err != nil {
		log.Fatalf("SayHello stream failed: %v", err)
	}

	for {
		res, err :=stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while streaming: %v", err)
		}

		log.Println(res.Message)
	}
	log.Println("Streaming finished")
}

func CallSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client streaming started")
	stream, err := client.SayhelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v ", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{Name:name}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("Sent the req with name: %v", name)

		time.Sleep(2 * time.Second)
	}

	reply, err := stream.CloseAndRecv()
	log.Printf("Client streaming finished")
	if err != nil {
		log.Fatalf("Error while receiving %v", err)
	}

	log.Printf("%v", reply.Messages)

}


func CallSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional streaming started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	waitc := make(chan struct{})
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming: %v", err)
			}

			log.Println(res)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		nm := &pb.HelloRequest{Name: name}
		if err := stream.Send(nm); err != nil {
			log.Fatalf("Got Error while sending request: %v", err)
		}

		time.Sleep(2 * time.Second)
	}

	stream.CloseSend()
	<- waitc

	log.Printf("Bidirectional streaming finished")
}
