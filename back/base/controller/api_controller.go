package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiController struct{}

func NewApiController() *ApiController {
	return &ApiController{}
}

type RequestBody struct {
	ID       string `json:"id"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
		Parts   []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"messages"`
}

func (c *ApiController) ChatHandler(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodPost {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}

	var request RequestBody
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	prompt := `
	- 你是资深全栈开发工程师兼代码架构师，能够对各种编程语言的代码进行深入分析。
	- 你接收编程语言、代码片段或文件路径以及分析维度后，会在 5 秒内给出分析结果。
	- 技能：
		- 讲解代码原理，解释代码的设计意图和实现思路。
		- 提供代码优化建议，包括重构、性能优化、代码规范等方面。
		- 提供代码重构建议，包括变量命名、函数命名、代码结构优化等方面。
		- 提供代码性能优化建议，包括算法优化、数据结构优化、并发优化等方面。
		- 提供代码可读性优化建议，包括代码格式优化、注释优化、变量命名优化等方面。
		- 提供代码安全性优化建议，包括输入验证、输出编码、异常处理等方面。
		- 提供代码可维护性优化建议，包括代码注释、代码结构优化、代码命名优化等方面。
		- 提供代码可扩展性优化建议，包括代码模块化、代码解耦、代码复用等方面。
		- 提供代码可测试性优化建议，包括单元测试、集成测试、端到端测试等方面。
		- 提供代码可移植性优化建议，包括代码兼容性、代码可移植性、代码可移植性等方面。
		- 静态代码结构分析，可分析模块依赖图、函数调用链、类继承关系、数据流向追踪、潜在风险点等维度。
		- 解析语法树（AST），建立符号索引库，绘制依赖关系图并生成人类可读摘要。
	- 输出规范：明确展示文件间的依赖关系、调用路径和影响范围等信息，并对潜在风险给出警告。
`
	prompt += "讲解项目框架:\n"
	for _, msg := range request.Messages {
		prompt += msg.Content + "\n"
	}

	payload := map[string]interface{}{
		"model":  "qwen2.5:7b",
		"prompt": prompt,
		"stream": true,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal payload"})
		return
	}

	resp, err := http.Post("http://host.docker.internal:11434/api/generate", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to send request to Ollama: %v", err)})
		return
	}
	defer resp.Body.Close()

	ctx.Writer.Header().Set("Content-Type", "text/event-stream")
	ctx.Writer.Header().Set("Cache-Control", "no-cache")
	ctx.Writer.Header().Set("Connection", "keep-alive")
	ctx.Writer.WriteHeader(http.StatusOK)
	flusher, ok := ctx.Writer.(http.Flusher)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Streaming unsupported!"})
		return
	}

	messageID := fmt.Sprintf("msg-%s", request.ID)
	startMsg := fmt.Sprintf("f:{\"messageId\":\"%s\"}\n", messageID)
	fmt.Fprint(ctx.Writer, startMsg)
	flusher.Flush()

	var usage struct {
		PromptTokens     int `json:"promptTokens"`
		CompletionTokens int `json:"completionTokens"`
	}
	usage.PromptTokens = 47

	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error reading response from Ollama: %v", err)})
			return
		}
		part := string(buf[:n])
		var ollamaResp map[string]interface{}
		if err := json.Unmarshal([]byte(part), &ollamaResp); err == nil {
			if text, ok := ollamaResp["response"].(string); ok {
				fmt.Printf("text: %#v\n", text)
				fmt.Fprintf(ctx.Writer, "0:%#v\n", text)
				usage.CompletionTokens++
				flusher.Flush()
			}
		}
	}

	finishInfo := map[string]interface{}{
		"finishReason": "stop",
		"usage":        usage,
		"isContinued":  false,
	}
	finishInfoBytes, _ := json.Marshal(finishInfo)
	fmt.Fprintf(ctx.Writer, "e:%s\n", string(finishInfoBytes))

	finishData := map[string]interface{}{
		"finishReason": "stop",
		"usage":        usage,
	}
	finishDataBytes, _ := json.Marshal(finishData)
	fmt.Fprintf(ctx.Writer, "d:%s\n", string(finishDataBytes))
	flusher.Flush()
}
