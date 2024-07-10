package implement

import (
	"azizdev/exception"
	"azizdev/helper"
	"azizdev/helper/model"
	"azizdev/model/domain"
	"azizdev/model/web/create"
	"azizdev/model/web/response"
	"azizdev/model/web/update"
	"azizdev/repository"
	"azizdev/service"
	"context"
	"database/sql"
	"github.com/go-playground/validator"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) service.UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate}
}

func (service *UserServiceImpl) Create(ctx context.Context, request create.UserCreateRequest) response.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		RoleId:   request.RoleId,
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	user = service.UserRepository.Save(ctx, tx, user)

	return model.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request update.UserUpdateRequest) response.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	user.RoleId = request.RoleId
	user.Name = request.Name
	user.Email = request.Email
	user.Password = request.Password

	user = service.UserRepository.Update(ctx, tx, user)

	return model.ToUserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, id)
	helper.PanicIfError(err)

	service.UserRepository.Delete(ctx, tx, user)
}

func (service *UserServiceImpl) FindById(ctx context.Context, id int) response.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return model.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []response.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := service.UserRepository.FindAll(ctx, tx)
	return model.ToUserResponses(user)
}
