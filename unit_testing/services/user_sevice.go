package service

import "unit-testing/repo"

type UserSevice interface {
	Login(username string) (string, error)
	Register(username, password string) error
	Logout(username string) error
}

type UserSeviceImp struct {
	userRepo repo.UserRepo
}

func NewUserSevice(userRepo repo.UserRepo) UserSevice {
	return &UserSeviceImp{
		userRepo: userRepo,
	}
}

func (u *UserSeviceImp) Login(username string) (string, error) {
	token, err := u.userRepo.GetUser(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *UserSeviceImp) Register(username, password string) error {
	return u.userRepo.AddUser(username, password)
}

func (u *UserSeviceImp) Logout(username string) error {
	_, err := u.userRepo.GetUser(username)
	if err != nil {
		return err
	}

	return err
}
