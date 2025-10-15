package helper

import userpb "github.com/fiqriardiansyah/shopping-proto/gen/go/user"

type Config struct {
	Prefix         string
	BaseUrl        string
	UserServiceUrl string
}

type GrpcClient struct {
	UserClient userpb.UserServiceClient
}
