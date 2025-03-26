package service

import (
	"dl/models"

	"github.com/go-xorm/xorm"
)

type ChapterService interface {
	ChapterList(courseId int) []models.Chapter
	GetContent(chapterId int) string
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
	err := cs.engine.Where(" course_id = ? ", courseId).OrderBy("id").Find(&chapterList)
	// err := cs.engine.OrderBy("id").Find(&chapterList)

	if err != nil {
		return chapterList
	}
	return chapterList
}

func (cs *chapterService) GetContent(chapterId int) string {
	content := ""

	item := models.Chapter{Id: chapterId}
	get, _ := cs.engine.Get(&item)
	// fmt.Println(get)
	if get { // 已有，更新
		// fmt.Println(item.ChapterContent)
		content = item.ChapterContent
	}

	return content
}
