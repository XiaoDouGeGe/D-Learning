package service

import (
	"dl/models"
	"strings"

	"github.com/go-xorm/xorm"
)

type UserService interface {
	Login(username, password string) (int, bool, error)
	Logout(userId int) (bool, error)
}

type userService struct {
	engine *xorm.Engine
}

func NewUserService(engine *xorm.Engine) *userService {
	return &userService{
		engine: engine,
	}
}

func (s *userService) Login(username, password string) (int, bool, error) {
	username = strings.TrimSpace(username)
	dbUser := models.User{Username: username, Password: password}
	get, err := s.engine.Get(&dbUser)
	if !get || err != nil {
		return -1, false, err
	}
	return dbUser.Id, true, nil
}

func (s *userService) Logout(userId int) (bool, error) {
	return true, nil
}
