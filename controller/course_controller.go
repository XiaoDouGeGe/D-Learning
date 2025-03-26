package controller

import (
	"dl/service"
	"strconv"

	iris "github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type CourseController struct {
	Ctx             iris.Context
	CourseService   service.CourseService
	ChapterService  service.ChapterService
	ProgressService service.ProgressService
}

func (cc *CourseController) GetList() mvc.Result {
	userId, err := cc.Ctx.URLParamInt("userId")
	if err != nil {
		userId = -1
	}

	courses := cc.CourseService.CourseList()
	courseList := make([]map[string]interface{}, 0)

	for _, course := range courses {
		courseItem := make(map[string]interface{})
		courseItem["id"] = course.Id
		courseItem["name"] = course.CourseName
		courseItem["desc"] = course.CourseDesc
		courseItem["img"] = course.CourseImg
		courseItem["status"] = course.Status

		completedNum, totalNum := 0, 0

		chapters := cc.ChapterService.ChapterList(course.Id)
		chaptersList := make([]map[string]interface{}, 0)

		for _, chapter := range chapters {
			chapterItem := make(map[string]interface{})
			chapterItem["id"] = chapter.Id
			chapterItem["index"] = chapter.ChapterIndex
			chapterItem["title"] = chapter.ChapterTitle
			chapterItem["status"] = chapter.Status

			// check progress
			chapterItem["isCompleted"] = false
			isCompleted := cc.ProgressService.IsCompleted(userId, chapter.Id)
			if isCompleted {
				completedNum++
				chapterItem["isCompleted"] = true
			}

			chaptersList = append(chaptersList, chapterItem)
		}

		totalNum = len(chaptersList)

		courseItem["chapters"] = chaptersList
		courseItem["progress"] = strconv.Itoa(completedNum) + "/" + strconv.Itoa(totalNum)

		courseList = append(courseList, courseItem)
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"errorno": 0,
			"data":    courseList,
		},
	}
}
