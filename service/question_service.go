package service

import (
	"dl/models"

	"github.com/go-xorm/xorm"
)

type QuestionService interface {
	QuestionListByChapterId(chapterId int) []models.Question
	RandomQuestionList(chapterId int, num int) []models.Question
}

type questionService struct {
	engine *xorm.Engine
}

func NewQuestionService(engine *xorm.Engine) *questionService {
	return &questionService{
		engine: engine,
	}
}

func (qs *questionService) QuestionListByChapterId(chapterId int) []models.Question { // 不做分页处理
	questionList := make([]models.Question, 0)
	err := qs.engine.Where("chapter_id = ?", chapterId).OrderBy("id").Find(&questionList)
	if err != nil {
		return questionList
	}
	return questionList
}

func (qs *questionService) RandomQuestionList(chapterId int, num int) []models.Question { // 随机获取题目，数量为num
	questionList := make([]models.Question, 0)
	err := qs.engine.Where("chapter_id = ?", chapterId).OrderBy("id").Limit(num).Find(&questionList)
	if err != nil {
		return questionList
	}
	return questionList
}
