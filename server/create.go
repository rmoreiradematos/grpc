package main

import (
	"context"
	"fmt"
	pb "github.com/maximilienandile/demo-grpc/gen/proto"
	"log"
)

func (s *Server) Echo(ctx context.Context, in *pb.NoSql) (*pb.NoSql, error) {
	log.Printf("Creating first collection with %v\n", in)
	data := SyncUnity{
		QuestionTypes: in.GetQuestionTypes(),
	}
	res, err := server.collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.ErrorF(codes.Internal,
			fmt.Sprintf("Internal error %v\n", in))
	}
	fmt.Println(res)
}
