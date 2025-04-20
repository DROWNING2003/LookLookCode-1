package sse

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseWriter 封装SSE响应的写入器
type ResponseWriter struct {
	ctx     *gin.Context
	flusher http.Flusher
}

// NewResponseWriter 创建一个新的SSE响应写入器
func NewResponseWriter(ctx *gin.Context) (*ResponseWriter, error) {
	// 设置SSE headers
	ctx.Writer.Header().Set("Content-Type", "text/event-stream")
	ctx.Writer.Header().Set("Cache-Control", "no-cache")
	ctx.Writer.Header().Set("Connection", "keep-alive")
	ctx.Writer.WriteHeader(http.StatusOK)

	flusher, ok := ctx.Writer.(http.Flusher)
	if !ok {
		return nil, fmt.Errorf("streaming not supported")
	}

	return &ResponseWriter{
		ctx:     ctx,
		flusher: flusher,
	}, nil
}

// SendStart 发送开始消息
func (w *ResponseWriter) SendStart(requestID string) {
	messageID := fmt.Sprintf("msg-%s", requestID)
	startMsg := fmt.Sprintf("f:{\"messageId\":\"%s\"}\n", messageID)
	fmt.Fprint(w.ctx.Writer, startMsg)
	w.flusher.Flush()
}

// SendContent 发送内容消息
func (w *ResponseWriter) SendContent(content string) {
	fmt.Fprintf(w.ctx.Writer, "0:%#v\n", content)
	w.flusher.Flush()
}

// Usage 用于统计token使用情况
type Usage struct {
	PromptTokens     int `json:"promptTokens"`
	CompletionTokens int `json:"completionTokens"`
}

// SendFinish 发送完成消息
func (w *ResponseWriter) SendFinish(usage Usage) error {
	finishInfo := map[string]interface{}{
		"finishReason": "stop",
		"usage":        usage,
		"isContinued":  false,
	}

	finishInfoBytes, err := json.Marshal(finishInfo)
	if err != nil {
		return err
	}

	fmt.Fprintf(w.ctx.Writer, "e:%s\n", string(finishInfoBytes))
	fmt.Fprintf(w.ctx.Writer, "d:%s\n", string(finishInfoBytes))
	w.flusher.Flush()
	return nil
}

// PipeDataStreamToResponse 将数据流写入响应
func (w *ResponseWriter) PipeDataStreamToResponse(requestID string, dataStream <-chan string) error {
	// 发送开始消息
	w.SendStart(requestID)

	// 处理数据流
	for data := range dataStream {
		w.SendContent(data)
	}

	// 发送完成消息
	return w.SendFinish(Usage{
		PromptTokens:     0,
		CompletionTokens: 0,
	})
}
