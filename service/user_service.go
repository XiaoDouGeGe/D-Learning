package service

import (
	"dl/models"
	"strings"

	"github.com/go-xorm/xorm"
)

type UserService interface {
	Login(username, password string) (int, bool, string, error)
	Logout(userId int) (bool, error)
	GetUserNameById(userId int) string
}

type userService struct {
	engine *xorm.Engine
}

func NewUserService(engine *xorm.Engine) *userService {
	return &userService{
		engine: engine,
	}
}

func (s *userService) Login(username, password string) (int, bool, string, error) {
	username = strings.TrimSpace(username)
	dbUser := models.User{Username: username, Password: password}
	get, err := s.engine.Where(" status = 'Y' ").Get(&dbUser)
	if !get || err != nil {
		return -1, false, "", err
	}
	return dbUser.Id, true, dbUser.Secret, nil
}

func (s *userService) Logout(userId int) (bool, error) {
	return true, nil
}

// 根据用户ID获取用户名称
func (s *userService) GetUserNameById(userId int) string {
	user := models.User{Id: userId}
	has, err := s.engine.Get(&user)
	if err != nil {
		return ""
	}
	if !has {
		return ""
	}
	return user.Username
}
