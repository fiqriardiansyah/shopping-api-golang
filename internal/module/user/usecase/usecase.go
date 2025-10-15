package usecase

import (
	"github.com/fiqriardiansyah/shopping-api-golang/internal/helper"
)

type UserUseCase struct {
	GrpcClient *helper.GrpcClient
}

func NewUserUseCase(client *helper.GrpcClient) *UserUseCase {
	return &UserUseCase{
		GrpcClient: client,
	}
}
