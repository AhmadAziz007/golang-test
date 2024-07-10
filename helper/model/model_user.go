package model

import (
	"azizdev/model/domain"
	"azizdev/model/web/response"
)

func ToUserResponse(user domain.User) response.UserResponse {
	return response.UserResponse{
		Id:       user.Id,
		RoleId:   user.RoleId,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
func ToUserResponses(user []domain.User) []response.UserResponse {
	var userResponses []response.UserResponse
	for _, userList := range user {
		userResponses = append(userResponses, ToUserResponse(userList))
	}
	return userResponses
}
