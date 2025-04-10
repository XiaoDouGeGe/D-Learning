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

// 新增题目和选项
func (qc *QuestionController) PostAdd() mvc.Result {
	chapterId := qc.Ctx.PostValueIntDefault("chapterId", 0)
	qContent := qc.Ctx.PostValue("qContent")
	qType := qc.Ctx.PostValue("qType")
	qAnswer := qc.Ctx.PostValue("qAnswer")
	qAnalysis := qc.Ctx.PostValue("qAnalysis")
	mark := qc.Ctx.PostValueIntDefault("mark", 0)

	aShow := qc.Ctx.PostValueIntDefault("aShow", 0)
	bShow := qc.Ctx.PostValueIntDefault("bShow", 0)
	cShow := qc.Ctx.PostValueIntDefault("cShow", 0)
	dShow := qc.Ctx.PostValueIntDefault("dShow", 0)

	aContent := qc.Ctx.PostValue("aContent")
	bContent := qc.Ctx.PostValue("bContent")
	cContent := qc.Ctx.PostValue("cContent")
	dContent := qc.Ctx.PostValue("dContent")

	qc.QuestionService.AddQuestion(chapterId, qContent, qType, qAnswer, qAnalysis, mark,
		aShow, bShow, cShow, dShow, aContent, bContent, cContent, dContent)

	return mvc.Response{
		Object: map[string]interface{}{
			"errorno": 0,
			"msg":     "ok",
		},
	}
}

// 根据chapterId查询问题列表
func (qc *QuestionController) GetList() mvc.Result {
	chapterId, err := qc.Ctx.URLParamInt("chapterId")
	if err != nil {
		chapterId = -1
	}

	questions := qc.QuestionService.QuestionListByChapterId(chapterId)

	questionList := make([]map[string]interface{}, 0)

	for _, question := range questions {
		questionItem := make(map[string]interface{})
		questionItem["id"] = question.Id
		questionItem["qContent"] = question.Qcontent
		questionItem["qType"] = question.Qtype
		questionItem["qAnswer"] = question.Answer
		questionItem["qAnalysis"] = question.Analysis
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
			"msg":     "ok",
		},
	}
}

// 删除题目
func (qc *QuestionController) PostDelete() mvc.Result {
	questionId := qc.Ctx.PostValueIntDefault("questionId", 0)
	if questionId == 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"errorno": 1,
				"msg":     "题目ID不能为空",
			},
		}
	}

	ok := qc.QuestionService.DeleteQuestion(questionId)
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
