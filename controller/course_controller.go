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

func (cc *CourseController) PostAdd() mvc.Result {
	courseName := cc.Ctx.PostValue("courseName")
	courseDesc := cc.Ctx.PostValue("courseDesc")

	if courseName == "" || courseDesc == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"errorno": 1,
				"msg":     "课程名称和描述不能为空",
			},
		}
	}

	if cc.CourseService.AddCourse(courseName, courseDesc) {
		return mvc.Response{
			Object: map[string]interface{}{
				"errorno": 0,
				"msg":     "添加成功",
			},
		}
	} else {
		return mvc.Response{
			Object: map[string]interface{}{
				"errorno": 1,
				"msg":     "添加失败",
			},
		}
	}
}

// 删除课程
func (cc *CourseController) PostDelete() mvc.Result {
	courseId := cc.Ctx.PostValueIntDefault("courseId", 0)
	if courseId == 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"errorno": 1,
				"msg":     "课程ID不能为空",
			},
		}
	}

	ok := cc.CourseService.DeleteCourse(courseId)
	if ok {
		return mvc.Response{
			Object: map[string]interface{}{
				"errorno": 0,
				"msg":     "删除成功",
			},
		}
	} else {
		return mvc.Response{
			Object: map[string]interface{}{
				"errorno": 1,
				"msg":     "删除失败",
			},
		}
	}
}

func (cc *CourseController) GetChapters() mvc.Result {

	courseOne := make(map[string]interface{})

	userId, err := cc.Ctx.URLParamInt("userId")
	if err != nil {
		userId = -1
	}

	courseId, err := cc.Ctx.URLParamInt("courseId")
	if err != nil {
		courseId = -1
	}

	completedNum, totalNum := 0, 0

	chapters := cc.ChapterService.ChapterList(courseId)
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

	courseOne["chapters"] = chaptersList
	courseOne["progress"] = strconv.Itoa(completedNum) + "/" + strconv.Itoa(totalNum)

	return mvc.Response{
		Object: map[string]interface{}{
			"errorno": 0,
			"data":    courseOne,
		},
	}

}
