package service

import (
	"split-bill/backend/dto/request"
	"split-bill/backend/dto/response"
)

type UserService interface {
	ForgetPassword(request.ForgetUserPasswordRequest) error
	ResetPassword(request.ResetUserPasswordRequest) error
	ChangePassword(request.ChangeUserPasswordRequest) error
	UpdateUser(request.UpdateUserRequest) (response.UserResponse, error)
	DeleteUser(request.DeleteUserRequest) error
}
