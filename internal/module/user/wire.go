//go:build wireinject
// +build wireinject

package user

import (
	"github.com/fiqriardiansyah/shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/module/user/usecase"
	"github.com/google/wire"
)

func InitializeUserHandler(*helper.GrpcClient) *UserController {
	wire.Build(NewUserController, usecase.NewUserUseCase)
	return nil
}
