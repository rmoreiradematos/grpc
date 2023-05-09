package controllers

import (
	"context"
	"fmt"
	pb "github.com/maximilienandile/demo-grpc/gen/proto"
	"go.mongodb.org/mongo-driver/mongo"
)

func SaveData(req *pb.NoSql, collection *mongo.Collection, ctx context.Context) {
	fmt.Println(req.GetQuestionTypes())

	data := SyncUnity{
		QuestionTypes: req.GetQuestionTypes(),
	}
	res, err := collection.InsertOne(ctx, data)

	fmt.Println(err)

	if err != nil {
		fmt.Sprintf("Internal error")
	}
	fmt.Println(res)
}
