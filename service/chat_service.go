package service

import (
	"context"
	"dl/models"
	"fmt"
	"strings"

	"github.com/go-xorm/xorm"
	openai "github.com/sashabaranov/go-openai"
)

type ChatService interface {
	ChatList(userId int, courseId int) []models.Chat
	SendContent(userId int, courseId int, content string) (string, error)
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

func (cs *chatService) SendContent(userId int, courseId int, content string) (string, error) {

	// 创建客户端配置
	cfg := openai.DefaultConfig("YOUR_API_KEY")
	cfg.BaseURL = "http://10.10.185.2:30003/v1"

	// 使用自定义配置创建客户端
	client := openai.NewClientWithConfig(cfg)
	ctx := context.Background()

	// 多轮对话（拼装历史记录）
	messages := make([]openai.ChatCompletionMessage, 0)

	chatRow := models.Chat{UserId: userId, CourseId: courseId, Role: "user", Content: content}
	cs.engine.Insert(&chatRow)

	chatList := make([]models.Chat, 0)
	err := cs.engine.Where("user_id = ?", userId).Where("course_id = ?", courseId).OrderBy("id").Find(&chatList)
	if err == nil {
		for _, chat := range chatList {
			if chat.Role == "system" {
				messages = append(messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleSystem, Content: chat.Content})
			} else if chat.Role == "user" {
				messages = append(messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: chat.Content})
			} else if chat.Role == "assistant" {
				messages = append(messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleAssistant, Content: chat.Content})
			}
		}
	}

	// messages = append(messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: content})

	// fmt.Println("=============")
	// fmt.Println(len(messages))
	// fmt.Println("=============")

	// 创建聊天请求
	req := openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: messages,
	}

	// 发送请求并获取响应
	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	// 输出响应内容
	// fmt.Println(resp.Choices[0].Message.Content)
	res := strings.Split(resp.Choices[0].Message.Content, "</think>")
	if len(res) >= 2 {
		// fmt.Println("=============")
		// fmt.Println(strings.TrimSpace(res[len(res)-1]))
		// fmt.Println("=============")

		receive := strings.TrimSpace(res[len(res)-1])

		// 1.存到数据库
		chatRow := models.Chat{UserId: userId, CourseId: courseId, Role: "assistant", Content: receive}
		_, err := cs.engine.Insert(&chatRow)
		fmt.Println(chatRow.Id)
		if err != nil {
			fmt.Println(err.Error())
		}

		// 2.返回给接口
		return receive, nil
	}

	return "", nil
}
