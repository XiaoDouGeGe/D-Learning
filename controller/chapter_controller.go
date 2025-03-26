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

	title, content := cc.ChapterService.GetContent(chapterId)

	data := make(map[string]interface{})
	data["title"] = title
	data["content"] = content

	return mvc.Response{
		Object: map[string]interface{}{
			"errorno": 0,
			"data":    data,
		},
	}
}
