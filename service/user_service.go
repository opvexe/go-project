package service

import "go-projecct/repository"

type IUserService interface {
	GetUserByUid(uid int)
}

func NewUserService() IUserService {
	return &userService{
		userRepo: repository.NewUserRepository(),
	}
}

type userService struct {
	userRepo repository.IUserRepository
}


func (u *userService) GetUserByUid(uid int) {

	/*
		业务操作
	 */
}