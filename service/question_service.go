package service

import (
	"dl/models"
	"fmt"

	"github.com/go-xorm/xorm"
)

type QuestionService interface {
	QuestionListByChapterId(chapterId int) []models.Question
	RandomQuestionList(chapterId int, num int) []models.Question
	AddQuestion(chapterId int, qContent string, qType string, qAnswer string, qAnalysis string, mark int,
		aShow int, bShow int, cShow int, dShow int, aContent string, bContent string,
		cContent string, dContent string) bool
}

type questionService struct {
	engine *xorm.Engine
}

func NewQuestionService(engine *xorm.Engine) *questionService {
	return &questionService{
		engine: engine,
	}
}

// 新增题目
func (qs *questionService) AddQuestion(chapterId int, qContent string, qType string, qAnswer string,
	qAnalysis string, mark int, aShow int, bShow int, cShow int, dShow int, aContent string,
	bContent string, cContent string, dContent string) bool {
	// 1. 新增题目
	question := models.Question{
		ChapterId: chapterId,
		Qcontent:  qContent,
		Qtype:     qType,
		Answer:    qAnswer,
		Analysis:  qAnalysis,
		Mark:      mark,
	}
	_, err1 := qs.engine.Insert(&question)
	fmt.Println(question.Id)
	fmt.Println(err1)

	// 2. 新增选项
	choice := models.Choice{
		QuestionId: question.Id,
		AContent:   aContent,
		BContent:   bContent,
		CContent:   cContent,
		DContent:   dContent,
		AShow:      aShow,
		BShow:      bShow,
		CShow:      cShow,
		DShow:      dShow,
	}
	_, err2 := qs.engine.Insert(&choice)
	fmt.Println(choice.Id)
	fmt.Println(err2)

	return true
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
