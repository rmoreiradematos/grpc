package main

import (
	"context"
	"log"

	pb "github.com/maximilienandile/demo-grpc/gen/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	client := pb.NewTesteApiClient(conn)

	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{})

	if err != nil {
		log.Println(err)
	}

	//fmt.Println(resp)
	//fmt.Println(resp.name)
	//fmt.Println(resp.idade)
}
