package usecase

import (
	"context"

	userpb "github.com/fiqriardiansyah/shopping-proto/gen/go/user"
	"github.com/google/uuid"
)

func (u *UserUseCase) GetUser(id uuid.UUID) (*userpb.GetUserResponse, error) {

	userResponse, err := u.GrpcClient.UserClient.GetUser(context.Background(), &userpb.GetUserRequest{
		Id: id.String(),
	})
	if err != nil {
		return nil, err
	}
	return userResponse, nil
}
