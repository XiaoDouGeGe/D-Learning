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

// 新增章节
func (cc *ChapterController) PostAdd() mvc.Result {
	courseId := cc.Ctx.PostValueIntDefault("courseId", 0)
	chapterTitle := cc.Ctx.PostValue("chapterTitle")
	chapterContent := cc.Ctx.PostValue("chapterContent")

	if chapterTitle == "" || chapterContent == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"errorno": 1,
				"msg":     "章节标题或内容不能为空",
			},
		}
	}

	ok := cc.ChapterService.AddChapter(chapterTitle, chapterContent, courseId)
	if ok {
		return mvc.Response{
			Object: map[string]interface{}{
				"errorno": 0,
				"msg":     "新增章节成功",
			},
		}
	} else {
		return mvc.Response{
			Object: map[string]interface{}{
				"errorno": 1,
				"msg":     "新增章节失败",
			},
		}
	}
}
