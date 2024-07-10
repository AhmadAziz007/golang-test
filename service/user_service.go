package service

import (
	"azizdev/model/web/create"
	"azizdev/model/web/response"
	"azizdev/model/web/update"
	"context"
)

type UserService interface {
	Create(ctx context.Context, request create.UserCreateRequest) response.UserResponse
	Update(ctx context.Context, request update.UserUpdateRequest) response.UserResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) response.UserResponse
	FindAll(ctx context.Context) []response.UserResponse
}
