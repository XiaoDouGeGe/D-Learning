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
	api_key, base_url := "b3ea194e-7be9-45e5-b7dd-aae02190161b", "https://api-inference.modelscope.cn/v1/"
	// api_key, base_url := "123456", "http://10.10.185.2:30003/v1"
	cfg := openai.DefaultConfig(api_key)
	cfg.BaseURL = base_url

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
		// Model:    openai.GPT3Dot5Turbo,
		Model:    "deepseek-ai/DeepSeek-R1",
		Messages: messages,
	}

	// 发送请求并获取响应
	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	// 输出响应内容
	receive := ""
	// fmt.Println(resp.Choices[0].Message.Content)
	if strings.Contains(resp.Choices[0].Message.Content, "</think>") {
		res := strings.Split(resp.Choices[0].Message.Content, "</think>")
		if len(res) >= 2 {
			// fmt.Println("=============")
			// fmt.Println(strings.TrimSpace(res[len(res)-1]))
			// fmt.Println("=============")

			receive = strings.TrimSpace(res[len(res)-1])
		}
	} else {
		receive = resp.Choices[0].Message.Content
	}

	// 1.存到数据库
	chatRow2 := models.Chat{UserId: userId, CourseId: courseId, Role: "assistant", Content: receive}
	_, err2 := cs.engine.Insert(&chatRow2)
	fmt.Println(chatRow2.Id)
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	// 2.返回给接口
	return receive, nil

}
