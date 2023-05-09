package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/maximilienandile/demo-grpc/controllers"
	pb "github.com/maximilienandile/demo-grpc/gen/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

var collection *mongo.Collection

type syncUnityServer struct {
	pb.UnimplementedSyncUnityServer
}

func (s *syncUnityServer) Echo(ctx context.Context, req *pb.NoSql) (*pb.NoSql, error) {
	fmt.Println(collection)
	controllers.SaveData(req, collection, ctx)
	return req, nil
}

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("echo", "localhost:8081", "end-point de recebimento de json")
)

func main() {
	go func() {
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/"))
		if err != nil {
			log.Fatal(err)
		}

		err = client.Connect(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		collection = client.Database("blogdb").Collection("blog")
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		mux := runtime.NewServeMux()

		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		pb.RegisterSyncUnityHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
		http.ListenAndServe(":8080", mux)
	}()
	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSyncUnityServer(grpcServer, &syncUnityServer{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(err)
	}
}
