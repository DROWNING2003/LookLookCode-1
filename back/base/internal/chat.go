package internal

import (
	"context"
	"log"

	"github.com/cloudwego/eino-ext/components/model/ollama"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

type ChatService struct {
	agent    *ollama.ChatModel
	template *prompt.DefaultChatTemplate
}

func NewChatService(ctx context.Context) (*ChatService, error) {
	chatModel := createOllamaChatModel(ctx)
	// todoTools := []tool.BaseTool{}
	// ragent, err := react.NewAgent(ctx, &react.AgentConfig{
	// 	Model: chatModel,
	// 	// ToolsConfig: compose.ToolsNodeConfig{
	// 	// 	Tools: todoTools,
	// 	// },
	// })
	// if err != nil {
	// 	log.Fatalln(err)
	// 	return nil, err
	// }
	return &ChatService{
		agent:    chatModel,
		template: createTemplate(),
	}, nil
}

func (s *ChatService) GetResponse(ctx context.Context, messages []Messages, userInput string) (*schema.StreamReader[*schema.Message], error) {
	// messages = append(messages, Message{Role: "user", Content: userInput})
	var chatHistory []*schema.Message
	for _, msg := range messages[:len(messages)-1] {
		if msg.Role == "user" {
			chatHistory = append(chatHistory, schema.UserMessage(msg.Content))
		} else {
			chatHistory = append(chatHistory, schema.AssistantMessage(msg.Content, nil))
		}
	}
	msgs, err := s.template.Format(ctx, map[string]any{
		"role":         "程序员鼓励师",
		"style":        "暧昧、积极、温暖且专业",
		"question":     userInput,
		"chat_history": chatHistory,
	})
	if err != nil {
		return nil, err
	}
	result, err := s.agent.Stream(ctx, msgs)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func createTemplate() *prompt.DefaultChatTemplate {
	return prompt.FromMessages(schema.FString,
		//schema.SystemMessage("你是一个{role}。你需要用{style}的语气回答问题。你的目标是帮助程序员保持积极乐观的心态，提供技术建议的同时也要关注他们的心理健康。"),
		schema.SystemMessage("你是一个{role}。你需要用{style}的语气回答问题。你的目标是在原来的代码添加注释"),
		schema.MessagesPlaceholder("chat_history", true),
		schema.UserMessage("问题: {question}"),
	)
}

func createOllamaChatModel(ctx context.Context) *ollama.ChatModel {
	chatModel, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
		BaseURL: "http://host.docker.internal:11434",
		Model:   "qwen2.5:7b",
	})
	if err != nil {
		panic(err)
	}
	return chatModel
}
