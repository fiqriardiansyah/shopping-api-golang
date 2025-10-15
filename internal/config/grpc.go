package config

import (
	"fmt"
	"os"

	"github.com/fiqriardiansyah/shopping-api-golang/internal/helper"
	userpb "github.com/fiqriardiansyah/shopping-proto/gen/go/user"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGrpcConnection() *grpc.ClientConn {
	port := os.Getenv("GRPC_USER_CONNECTION_PORT")

	conn, err := grpc.NewClient(fmt.Sprintf(":%s", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	logrus.Info(fmt.Sprintf("GRPC client running on port: %s", port))
	return conn
}

func NewGrpcClient() *helper.GrpcClient {
	conn := NewGrpcConnection()
	userClient := userpb.NewUserServiceClient(conn)
	return &helper.GrpcClient{
		UserClient: userClient,
	}
}
