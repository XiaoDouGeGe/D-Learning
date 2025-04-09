package controller

import (
	"dl/service"

	iris "github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {
	Ctx     iris.Context
	Service service.UserService
}

func (us *UserController) PostLogin() mvc.Result {
	username := us.Ctx.PostValue("username")
	password := us.Ctx.PostValue("password")
	userId, isSuccess, secret, err := us.Service.Login(username, password)
	data := make(map[string]interface{})
	data["userId"] = userId
	data["isSuccess"] = isSuccess
	data["secret"] = secret
	data["err"] = err
	if !isSuccess {
		return mvc.Response{
			Object: map[string]interface{}{
				"errorno": -1,
				"detail":  data,
			},
		}
	}
	return mvc.Response{
		Object: map[string]interface{}{
			"errorno": 0,
			"detail":  data,
		},
	}
}

func (us *UserController) PostLogout() mvc.Result {
	userId := us.Ctx.PostValueIntDefault("userId", 1)
	isSuccess, err := us.Service.Logout(userId)
	data := make(map[string]interface{})
	data["isSuccess"] = isSuccess
	data["err"] = err
	if !isSuccess {
		return mvc.Response{
			Object: map[string]interface{}{
				"errorno": -1,
				"detail":  data,
			},
		}
	}
	return mvc.Response{
		Object: map[string]interface{}{
			"errorno": 0,
			"detail":  data,
		},
	}
}
