package controller

import (
	"dl/service"

	iris "github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type QuestionController struct {
	Ctx             iris.Context
	QuestionService service.QuestionService
	ChoiceService   service.ChoiceService
}

func (qc *QuestionController) GetRandomQuestions() mvc.Result {

	chapterId, err := qc.Ctx.URLParamInt("chapterId")
	if err != nil {
		chapterId = -1
	}

	num, err := qc.Ctx.URLParamInt("num")
	if err != nil {
		num = 20
	}

	questions := qc.QuestionService.RandomQuestionList(chapterId, num)

	questionList := make([]map[string]interface{}, 0)

	for _, question := range questions {
		questionItem := make(map[string]interface{})
		questionItem["id"] = question.Id
		questionItem["qContent"] = question.Qcontent
		questionItem["qType"] = question.Qtype
		// questionItem["qAnswer"] = question.Answer
		// questionItem["qAnalysis"] = question.Analysis
		questionItem["mark"] = question.Mark

		// get choices
		choice := qc.ChoiceService.GetChoice(question.Id)

		questionItem["choices"] = choice

		questionList = append(questionList, questionItem)
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"errorno": 0,
			"data":    questionList,
		},
	}
}
