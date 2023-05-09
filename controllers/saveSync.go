package controllers

import (
	pb "github.com/maximilienandile/demo-grpc/gen/proto"
)

type SyncUnity struct {
	QuestionTypes []*pb.NoSql_QuestionTypes `bson:"question_types"`
}

func SaveSync(data *SyncUnity) *pb.NoSql {
	return &pb.NoSql{QuestionTypes: data.QuestionTypes}
}
