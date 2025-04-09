package service

import (
	"dl/models"

	"github.com/go-xorm/xorm"
)

type CourseService interface {
	CourseList() []models.Course
	AddCourse(courseName string, courseDesc string) bool
}

type courseService struct {
	engine *xorm.Engine
}

func NewCourseService(engine *xorm.Engine) *courseService {
	return &courseService{
		engine: engine,
	}
}

func (cs *courseService) CourseList() []models.Course {
	courseList := make([]models.Course, 0)
	err := cs.engine.OrderBy("id").Find(&courseList)
	if err != nil {
		return courseList
	}
	return courseList
}

// 新增课程
func (cs *courseService) AddCourse(courseName string, courseDesc string) bool {
	course := models.Course{
		CourseName: courseName,
		CourseDesc: courseDesc,
	}
	_, err := cs.engine.Insert(&course)
	return err == nil
}
