package controller

import (
	"dl/service"

	iris "github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ProgressController struct {
	Ctx     iris.Context
	Service service.ProgressService
}

func (pc *ProgressController) PostStatus() mvc.Result {
	userId := pc.Ctx.PostValueIntDefault("userId", 0)
	chapterId := pc.Ctx.PostValueIntDefault("chapterId", 0)
	status := pc.Ctx.PostValue("status")
	pc.Service.SetStatus(userId, chapterId, status)

	return mvc.Response{
		Object: map[string]interface{}{
			"errorno": 0,
			"detail":  nil,
		},
	}
}
