package service

import (
	"dl/models"
	"fmt"

	"github.com/go-xorm/xorm"
)

type ChatService interface {
	ChatList(userId int, courseId int) []models.Chat
}

type chatService struct {
	engine *xorm.Engine
}

func NewChatService(engine *xorm.Engine) *chatService {
	return &chatService{
		engine: engine,
	}
}

func (cs *chatService) ChatList(userId int, courseId int) []models.Chat {
	chatList := make([]models.Chat, 0)
	err := cs.engine.Where("user_id = ?", userId).Where("course_id = ?", courseId).Where("id not in (1)").OrderBy("id").Find(&chatList)
	if err != nil {
		fmt.Println(err)
		return chatList
	}
	// fmt.Println(len(chatList))
	return chatList
}

func (cs *chatService) ChatCompletion(userId int, courseId int) bool {
	return true
}
