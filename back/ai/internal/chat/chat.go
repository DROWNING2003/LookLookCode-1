package chat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/cloudwego/eino-ext/components/model/ollama"
	"github.com/cloudwego/eino-ext/components/tool/duckduckgo"
	"github.com/cloudwego/eino-ext/components/tool/duckduckgo/ddgsearch"
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
)

type ChatService struct {
	ragent    *react.Agent
	agent     compose.Runnable[[]*schema.Message, []*schema.Message]
	chatModel *ollama.ChatModel
	template  *prompt.DefaultChatTemplate
}

func NewChatService(ctx context.Context) (*ChatService, error) {
	chatModel := createOllamaChatModel(ctx)

	// 创建 duckduckgo Search 工具
	ducktool, err := duckduckgo.NewTool(ctx, &duckduckgo.Config{
		ToolName:   "duckduckgo_search",                        // 工具名称
		ToolDesc:   "search web for information by duckduckgo", // 工具描述
		Region:     ddgsearch.RegionWT,                         // 搜索地区
		MaxResults: 10,                                         // 每页结果数量
		SafeSearch: ddgsearch.SafeSearchOff,                    // 安全搜索级别
		TimeRange:  ddgsearch.TimeRangeAll,                     // 时间范围
		DDGConfig:  &ddgsearch.Config{},                        // DuckDuckGo 配置
	})
	// 初始化 tools
	todoTools := []tool.BaseTool{
		ducktool, // 官方封装的工具
	}
	// 获取工具信息并绑定到 ChatModel
	toolInfos := make([]*schema.ToolInfo, 0, len(todoTools))
	for _, tool := range todoTools {
		info, err := tool.Info(ctx)
		if err != nil {
			log.Fatal(err)
		}
		toolInfos = append(toolInfos, info)
	}
	err = chatModel.BindTools(toolInfos)
	if err != nil {
		log.Fatal(err)
	}

	// 创建 tools 节点
	todoToolsNode, err := compose.NewToolNode(context.Background(), &compose.ToolsNodeConfig{
		Tools: todoTools,
	})
	if err != nil {
		log.Fatal(err)
	}
	ragent, err := react.NewAgent(ctx, &react.AgentConfig{
		ToolCallingModel: chatModel,
		ToolsConfig: compose.ToolsNodeConfig{
			Tools: todoTools,
		},
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	chain := compose.NewChain[[]*schema.Message, []*schema.Message]()
	chain.
		AppendChatModel(chatModel, compose.WithNodeName("chat_model")).
		AppendToolsNode(todoToolsNode, compose.WithNodeName("tools"))
		// 编译并运行 chain
	agent, err := chain.Compile(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return &ChatService{
		ragent:    ragent,
		agent:     agent,
		chatModel: chatModel,
		template:  createTemplate(),
	}, nil
}

func (s *ChatService) GetResponse(ctx context.Context, messages []Message, userInput string) (string, error) {
	messages = append(messages, Message{Role: "user", Content: userInput})

	// Prepare chat history
	var chatHistory []*schema.Message
	for _, msg := range messages[:len(messages)-1] {
		if msg.Role == "user" {
			chatHistory = append(chatHistory, schema.UserMessage(msg.Content))
		} else {
			chatHistory = append(chatHistory, schema.AssistantMessage(msg.Content, nil))
		}
	}

	// Format template messages
	msgs, err := s.template.Format(ctx, map[string]any{
		"role":         "程序员鼓励师",
		"style":        "暧昧、积极、温暖且专业",
		"question":     userInput,
		"chat_history": chatHistory,
	})
	if err != nil {
		return "", err
	}

	result, err := s.ragent.Stream(ctx, msgs)
	if err != nil {
		return "", err
	}
	defer result.Close()

	var response strings.Builder
	for {
		message, err := result.Recv()
		// fmt.Printf("message: %v\n", message)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		// 实时输出消息内容
		fmt.Print(message)
		// response.WriteString(message.Content)
	}

	return response.String(), nil
}

func createOllamaChatModel(ctx context.Context) *ollama.ChatModel {
	chatModel, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
		BaseURL: "http://localhost:11434",
		Model:   "qwen2.5:7b",
	})
	if err != nil {
		panic(err)
	}
	return chatModel
}

func createTemplate() *prompt.DefaultChatTemplate {
	return prompt.FromMessages(schema.FString,
		schema.SystemMessage("你是一个{role}。你需要用{style}的语气回答问题。你的目标是帮助程序员保持积极乐观的心态，提供技术建议的同时也要关注他们的心理健康。"),
		schema.MessagesPlaceholder("chat_history", true),
		schema.UserMessage("问题: {question}"),
	)
}

type LoggerCallback struct {
	callbacks.HandlerBuilder // 可以用 callbacks.HandlerBuilder 来辅助实现 callback
}

func (cb *LoggerCallback) OnStart(ctx context.Context, info *callbacks.RunInfo, input callbacks.CallbackInput) context.Context {
	fmt.Println("==================")
	inputStr, _ := json.MarshalIndent(input, "", "  ") // nolint: byted_s_returned_err_check
	fmt.Printf("[OnStart] %s\n", string(inputStr))
	return ctx
}

func (cb *LoggerCallback) OnEnd(ctx context.Context, info *callbacks.RunInfo, output callbacks.CallbackOutput) context.Context {
	fmt.Println("=========[OnEnd]=========")
	outputStr, _ := json.MarshalIndent(output, "", "  ") // nolint: byted_s_returned_err_check
	fmt.Println(string(outputStr))
	return ctx
}

func (cb *LoggerCallback) OnError(ctx context.Context, info *callbacks.RunInfo, err error) context.Context {
	fmt.Println("=========[OnError]=========")
	fmt.Println(err)
	return ctx
}

func (cb *LoggerCallback) OnEndWithStreamOutput(ctx context.Context, info *callbacks.RunInfo,
	output *schema.StreamReader[callbacks.CallbackOutput]) context.Context {

	var graphInfoName = react.GraphName

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("[OnEndStream] panic err:", err)
			}
		}()

		defer output.Close() // remember to close the stream in defer

		fmt.Println("=========[OnEndStream]=========")
		for {
			frame, err := output.Recv()
			if errors.Is(err, io.EOF) {
				// finish
				break
			}
			if err != nil {
				fmt.Printf("internal error: %s\n", err)
				return
			}

			s, err := json.Marshal(frame)
			if err != nil {
				fmt.Printf("internal error: %s\n", err)
				return
			}

			if info.Name == graphInfoName { // 仅打印 graph 的输出, 否则每个 stream 节点的输出都会打印一遍
				fmt.Printf("%s: %s\n", info.Name, string(s))
			}
		}

	}()
	return ctx
}

func (cb *LoggerCallback) OnStartWithStreamInput(ctx context.Context, info *callbacks.RunInfo,
	input *schema.StreamReader[callbacks.CallbackInput]) context.Context {
	defer input.Close()
	return ctx
}
