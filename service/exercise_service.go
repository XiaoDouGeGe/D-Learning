package service

import (
	"dl/models"
	"encoding/json"
	"fmt"

	"github.com/go-xorm/xorm"
)

type ExerciseService interface {
	CommitContent(userId int, chapterId int, content string, etime string, score int) bool
	CalculateScore(content string) (int, error)
	GetExerciseList(userId int) []models.Exercise
	GetExerciseDetail(exId int) map[string]interface{}
}

type exerciseService struct {
	engine *xorm.Engine
}

func NewExerciseService(engine *xorm.Engine) *exerciseService {
	return &exerciseService{
		engine: engine,
	}
}

// 提交练习
func (es *exerciseService) CommitContent(userId int, chapterId int, content string, etime string, score int) bool {
	// fmt.Println(userId, chapterId, content, etime)
	item := models.Exercise{UserId: userId, ChapterId: chapterId, Content: content, ETime: etime, Score: score}
	es.engine.Insert(&item)
	fmt.Println(item.Id)
	return true
}

// 计算练习分数
func (es *exerciseService) CalculateScore(content string) (int, error) {
	eScore := 0

	// 解析JSON字符串
	var data []map[string]interface{}
	err := json.Unmarshal([]byte(content), &data)
	if err != nil {
		return 0, err
	}
	for _, d := range data {
		qId := d["qId"]
		uAnswer := d["uAnswer"]

		// 这里需要根据题目ID查询正确答案
		if num, ok := qId.(int); ok {
			qItem := models.Question{Id: num}
			get, err := es.engine.Get(&qItem)
			if get && err == nil {
				if qItem.Answer == uAnswer {
					eScore += qItem.Mark
				}
			}
		}
	}

	return eScore, nil
}

// 获取练习列表
func (es *exerciseService) GetExerciseList(userId int) []models.Exercise {
	exerciseList := make([]models.Exercise, 0)
	err := es.engine.Where("user_id = ?", userId).OrderBy("-id").Find(&exerciseList)
	if err != nil {
		return exerciseList
	}
	return exerciseList
}

// 获取练习详情
func (es *exerciseService) GetExerciseDetail(exId int) map[string]interface{} {
	exercise := models.Exercise{Id: exId}
	get1, err := es.engine.Get(&exercise)
	if err != nil || !get1 {
		return nil
	}

	// 解析JSON字符串
	var data []map[string]interface{}
	err = json.Unmarshal([]byte(exercise.Content), &data)
	if err != nil {
		return nil
	}

	exerciseDetail := make(map[string]interface{})
	exerciseDetail["id"] = exercise.Id
	exerciseDetail["userId"] = exercise.UserId
	exerciseDetail["chapterId"] = exercise.ChapterId
	exerciseDetail["score"] = exercise.Score
	exerciseDetail["eTime"] = exercise.ETime

	questions := make([]map[string]interface{}, 0)

	for _, d := range data {
		qId := d["qId"]
		// uAnswer := d["uAnswer"]

		// 问题及答案
		if num, ok := qId.(int); ok {
			qItem := models.Question{Id: num}
			get2, err := es.engine.Get(&qItem)
			if get2 && err == nil {
				d["qContent"] = qItem.Qcontent
				d["qType"] = qItem.Qtype
				d["qAnswer"] = qItem.Answer
				d["qAnalysis"] = qItem.Analysis
				d["mark"] = qItem.Mark

				// 获取选项
				var choices []map[string]string
				cItem := models.Choice{QuestionId: num}
				get3, _ := es.engine.Get(&cItem)
				if get3 { // 存在
					if cItem.AShow == 1 {
						choices = append(choices, map[string]string{"A": cItem.AContent})
					}
					if cItem.BShow == 1 {
						choices = append(choices, map[string]string{"B": cItem.BContent})
					}
					if cItem.CShow == 1 {
						choices = append(choices, map[string]string{"C": cItem.CContent})
					}
					if cItem.DShow == 1 {
						choices = append(choices, map[string]string{"D": cItem.DContent})
					}
				}
				d["choices"] = choices
			}
		}
		questions = append(questions, d)
	}

	exerciseDetail["questions"] = questions

	return exerciseDetail
}
