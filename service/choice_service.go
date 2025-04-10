package service

import (
	"dl/models"

	"github.com/go-xorm/xorm"
)

type ChoiceService interface {
	GetChoice(questionId int) []map[string]string
}

type choiceService struct {
	engine *xorm.Engine
}

func NewChoiceService(engine *xorm.Engine) *choiceService {
	return &choiceService{
		engine: engine,
	}
}

func (cs *choiceService) GetChoice(questionId int) []map[string]string {

	var choices []map[string]string

	item := models.Choice{QuestionId: questionId}

	get, _ := cs.engine.Where(" status = 'Y' ").Get(&item)

	if get { // 存在
		if item.AShow == 1 {
			choices = append(choices, map[string]string{"A": item.AContent})
		}
		if item.BShow == 1 {
			choices = append(choices, map[string]string{"B": item.BContent})
		}
		if item.CShow == 1 {
			choices = append(choices, map[string]string{"C": item.CContent})
		}
		if item.DShow == 1 {
			choices = append(choices, map[string]string{"D": item.DContent})
		}
	}

	return choices
}
