package controller

import (
	"dl/service"
	"time"

	iris "github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ExerciseController struct {
	Ctx            iris.Context
	Service        service.ExerciseService
	ChapterService service.ChapterService
	UserService    service.UserService
}

func (ec *ExerciseController) PostCommit() mvc.Result {
	userId := ec.Ctx.PostValueIntDefault("userId", 0)
	chapterId := ec.Ctx.PostValueIntDefault("chapterId", 0)
	content := ec.Ctx.PostValue("content")

	// 获取当前时间
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	// fmt.Println("格式化后的时间: ", formattedTime)

	// 计算分数
	score, _ := ec.Service.CalculateScore(content)

	isOK := ec.Service.CommitContent(userId, chapterId, content, formattedTime, score)

	// 提交失败
	if !isOK {
		return mvc.Response{
			Object: map[string]interface{}{
				"errorno": -1,
				"detail":  nil,
			},
		}
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"errorno": 0,
			"detail":  nil,
		},
	}

}

func (ec *ExerciseController) GetList() mvc.Result {
	userId, err := ec.Ctx.URLParamInt("userId") // userId=0 表示所有用户
	if err != nil {
		userId = -1
	}

	exs := ec.Service.GetExerciseList(userId)
	exList := make([]map[string]interface{}, 0)

	for _, ex := range exs {
		exItem := make(map[string]interface{})
		exItem["id"] = ex.Id
		exItem["userId"] = ex.UserId
		exItem["userName"] = ec.UserService.GetUserNameById(ex.UserId)
		// exItem["chapterId"] = ex.ChapterId
		// exItem["content"] = ex.Content
		exItem["score"] = ex.Score
		exItem["eTime"] = ex.ETime

		// 根据chapterId，获取章节名称和课程名称
		chapterName, courseName := ec.ChapterService.GetChapterName(ex.ChapterId)
		exItem["courseName"] = courseName
		exItem["chapterName"] = chapterName

		exList = append(exList, exItem)
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"errorno": 0,
			"data":    exList,
		},
	}
}

// 获取练习详情
func (ec *ExerciseController) GetDetail() mvc.Result {
	detail := make(map[string]interface{})

	exId, err := ec.Ctx.URLParamInt("exId")
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"errorno": 0,
				"data":    detail,
			},
		}
	}

	detail = ec.Service.GetExerciseDetail(exId)

	return mvc.Response{
		Object: map[string]interface{}{
			"errorno": 0,
			"data":    detail,
		},
	}

}
