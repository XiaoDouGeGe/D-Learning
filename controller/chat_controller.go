package controller

import (
	"dl/service"

	iris "github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ChatController struct {
	Ctx         iris.Context
	ChatService service.ChatService
}

func (cc *ChatController) GetList() mvc.Result {
	userId, err := cc.Ctx.URLParamInt("userId")
	if err != nil {
		userId = -1
	}
	courseId, err := cc.Ctx.URLParamInt("courseId")
	if err != nil {
		courseId = -1
	}

	chats := cc.ChatService.ChatList(userId, courseId)
	chatList := make([]map[string]interface{}, 0)

	for _, chat := range chats {
		chatItem := make(map[string]interface{})
		chatItem["id"] = chat.Id
		chatItem["userId"] = chat.UserId
		chatItem["courseId"] = chat.CourseId
		chatItem["role"] = chat.Role
		chatItem["content"] = chat.Content
		chatItem["chatTimeF"] = chat.ChatTimeF
		chatItem["next"] = chat.Next

		chatList = append(chatList, chatItem)
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"errorno": 0,
			"data":    chatList,
		},
	}
}
