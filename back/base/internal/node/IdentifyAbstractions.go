package node

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudwego/eino-ext/components/model/ollama"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

// type IdentifyAbstractionsNode struct {
// 	AgentMeta    *ollama.ChatModel
// 	template *prompt.DefaultChatTemplate
// }

func NewIdentifyAbstractionsNode(ctx context.Context) (compose.Runnable[*schema.Message, *schema.Message], error) {
	chain := compose.NewChain[*schema.Message, *schema.Message]()
	chain.AppendLambda(prep()).
		AppendLambda(compose.InvokableLambda(func(ctx context.Context, input *schema.Message) ([]*schema.Message, error) {
			format, err := createTemplate().Format(ctx, map[string]any{
				"Project_name":            input.Extra["Project_name"],
				"context":                 input.Content,
				"Language":                input.Extra["Language"],
				"file_listing_for_prompt": input.Extra["file_listing_for_prompt"],
				//"chat_history":            input[0].Extra["chat_history"],
			})
			if err != nil {
				return nil, err
			}
			return format, nil
		})).
		AppendChatModel(createOllamaChatModel(ctx))
	//.AppendLambda(test())
	// .AppendLambda(prepar())
	// chatModel := createOllamaChatModel(ctx)
	// return &IdentifyAbstractionsNode{
	// 	model:    chatModel,
	// 	template: createTemplate(),
	// }, nil
	r, err := chain.Compile(ctx)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// input 和 output 类型为自定义的任何类型
func prep() *compose.Lambda {
	lambda := compose.InvokableLambda(func(ctx context.Context, input *schema.Message) (output *schema.Message, err error) {
		index := 0
		c := ""
		fileinfo := []struct {
			idx  int
			path string
		}{}

		//if input[0].Extra["PrepInput"] == "" {
		//	input[0].Extra["Language"] = "中文"
		//}

		for path, content := range input.Extra["PrepInput"].(PrepInput).Files {
			c += fmt.Sprintf("--- File Index %d: %s ---\n%s\n\n", index, path, content)
			fileinfo = append(fileinfo, struct {
				idx  int
				path string
			}{idx: index, path: path})
			index++
		}
		file_listing_for_prompt := buildFileListing(fileinfo)

		return &schema.Message{
			Role:         schema.System,
			Content:      c,
			MultiContent: nil,
			Name:         "文件预处理",
			ToolCalls:    nil,
			ToolCallID:   "",
			ResponseMeta: nil,
			Extra: map[string]any{
				"file_listing_for_prompt": file_listing_for_prompt,
				"Language":                "中文",
				"use_cache":               false,
				"max_abstraction_num":     1024,
			},
		}, nil
	})
	return lambda
}

// 测试
func test() *compose.Lambda {
	lambda := compose.InvokableLambda(func(ctx context.Context, input []*schema.Message) (*schema.Message, error) {
		fmt.Printf("input messages: %v\n", input)
		return &schema.Message{
			Role:    schema.Assistant,
			Content: "Test message response",
		}, nil
	})
	return lambda
}

func buildFileListing(fileInfo []struct {
	idx  int
	path string
}) string {
	var builder strings.Builder
	for _, info := range fileInfo {
		builder.WriteString(fmt.Sprintf("- %d # %s\n", info.idx, info.path))
	}
	return strings.TrimSuffix(builder.String(), "\n") // 移除最后多余的换行
}

func createTemplate() *prompt.DefaultChatTemplate {
	template := `
使用{Language}回答
For the project {Project_name}:

Codebase Context:{context}

Analyze the codebase context.
Identify the top 5-10 core most important abstractions...

For each abstraction, provide:
1. A concise "name".
2. A beginner-friendly "description" explaining what it is with a simple analogy...
3. A list of relevant "file_indices"...

List of file indices and paths present in the context:
{file_listing_for_prompt}

Format the output as a YAML list of dictionaries:
` + "```yaml\n" +
		`- name: |
    Abstraction Name
  description: |
    A simple explanation with an analogy.
  file_indices:
    - 0 # path/to/file1.py
    - 3 # path/to/related.py
# ... more abstractions ...
` + "```"
	return prompt.FromMessages(schema.FString,
		//schema.SystemMessage("你是一个{role}。你需要用{style}的语气回答问题。你的目标是帮助程序员保持积极乐观的心态，提供技术建议的同时也要关注他们的心理健康。"),
		schema.SystemMessage(template),
		schema.MessagesPlaceholder("chat_history", true),
		//schema.UserMessage("问题: {question}"),
	)
}

func createOllamaChatModel(ctx context.Context) *ollama.ChatModel {
	chatModel, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
		BaseURL: "http://127.0.0.1:11434",
		Model:   "qwen2.5:7b",
	})
	if err != nil {
		panic(err)
	}
	return chatModel
}

// IdentifyAbstractions的 prep
type prepData struct {
	context                 string
	project_name            string
	file_listing_for_prompt string
	language                string
	use_cache               bool
	max_abstraction_num     int
}
type PrepInput struct {
	Files        map[string]string `json:"files,omitempty"`
	Project_name string            `json:"projectName,omitempty"`
	Language     string            `json:"language,omitempty"`
}
