package controller

import (
	"dl/service"

	iris "github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ChapterController struct {
	Ctx            iris.Context
	ChapterService service.ChapterService
}

func (cc *ChapterController) GetContent() mvc.Result {
	chapterId, err := cc.Ctx.URLParamInt("chapterId")
	if err != nil {
		chapterId = -1
	}

	content := cc.ChapterService.GetContent(chapterId)

	return mvc.Response{
		Object: map[string]interface{}{
			"errorno": 0,
			"data":    content,
		},
	}
}
