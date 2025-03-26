package service

import (
	"dl/models"
	"fmt"

	"github.com/go-xorm/xorm"
)

type ProgressService interface {
	IsCompleted(userId int, chapterId int) bool
	SetStatus(userId int, chapterId int, status string) bool
}

type progressService struct {
	engine *xorm.Engine
}

func NewProgressService(engine *xorm.Engine) *progressService {
	return &progressService{
		engine: engine,
	}
}

func (ps *progressService) IsCompleted(userId int, chapterId int) bool {
	completedItem := models.Progress{UserId: userId, ChapterId: chapterId, Status: "completed"}
	get, err := ps.engine.Get(&completedItem)

	if !get || err != nil {
		return false
	}

	return true
}

func (ps *progressService) SetStatus(userId int, chapterId int, status string) bool {
	fmt.Println(userId, chapterId, status)
	item := models.Progress{UserId: userId, ChapterId: chapterId}
	get, err := ps.engine.Get(&item)
	fmt.Println(get)
	if !get || err != nil { // 没有，创建
		progress := models.Progress{UserId: userId, ChapterId: chapterId, Status: status}
		ps.engine.Insert(&progress)
		fmt.Println(progress.Id)
	} else { // 已有，更新
		progress := models.Progress{UserId: userId, ChapterId: chapterId, Status: status}
		ps.engine.Where(" user_id = ? ", userId).Where(" chapter_id = ? ", chapterId).Update(&progress)
	}

	return true
}
