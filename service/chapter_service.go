package service

import (
	"dl/models"

	"github.com/go-xorm/xorm"
)

type ChapterService interface {
	ChapterList(courseId int) []models.Chapter
	GetContent(chapterId int) (string, string)
	GetChapterName(chapterId int) (string, string)
	AddChapter(chapterTitle string, chapterContent string, courseId int) bool
	DeleteChapter(chapterId int) bool
}

type chapterService struct {
	engine *xorm.Engine
}

func NewChapterService(engine *xorm.Engine) *chapterService {
	return &chapterService{
		engine: engine,
	}
}

func (cs *chapterService) ChapterList(courseId int) []models.Chapter {
	chapterList := make([]models.Chapter, 0)
	// err := cs.engine.Where(" course_id = ? ", strconv.Itoa(courseId)).OrderBy("id").Find(&chapterList)
	err := cs.engine.Where(" status = 'Y' ").Where(" course_id = ? ", courseId).OrderBy("id").Find(&chapterList)
	// err := cs.engine.OrderBy("id").Find(&chapterList)

	if err != nil {
		return chapterList
	}
	return chapterList
}

func (cs *chapterService) GetContent(chapterId int) (string, string) {
	title, content := "", ""

	item := models.Chapter{Id: chapterId}
	get, _ := cs.engine.Get(&item)
	// fmt.Println(get)
	if get { // 已有
		// fmt.Println(item.ChapterContent)
		title = item.ChapterTitle
		content = item.ChapterContent
	}

	return title, content
}

// 根据chapterId，获取章节名称和课程名称
func (cs *chapterService) GetChapterName(chapterId int) (string, string) {
	title, courseName := "", ""

	chapterItem := models.Chapter{Id: chapterId}
	get1, _ := cs.engine.Get(&chapterItem)
	if get1 { // 已有
		title = chapterItem.ChapterTitle
		// 根据courseId获取课程名称
		courseItem := models.Course{Id: chapterItem.CourseId}
		get2, _ := cs.engine.Get(&courseItem)
		if get2 {
			courseName = courseItem.CourseName
		}
	}

	return title, courseName
}

// 新增章节
func (cs *chapterService) AddChapter(chapterTitle string, chapterContent string, courseId int) bool {
	chapter := models.Chapter{
		ChapterTitle:   chapterTitle,
		ChapterContent: chapterContent,
		CourseId:       courseId,
	}
	_, err := cs.engine.Insert(&chapter)
	return err == nil
}

// 删除章节
func (cs *chapterService) DeleteChapter(chapterId int) bool {
	chapter := models.Chapter{Status: "N"}
	_, err := cs.engine.Where(" id = ? ", chapterId).Update(&chapter)
	return err == nil
}
