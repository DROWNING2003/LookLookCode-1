package node

import (
	"context"
	"testing"

	"github.com/cloudwego/eino/schema"
)

func TestNewIdentifyAbstractionsNode(t *testing.T) {
	ctx := context.Background()

	// 创建Specialist实例
	r, err := NewIdentifyAbstractionsNode(ctx)

	// 验证返回值
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	prepInput := PrepInput{
		Files: map[string]string{
			"python.py": "print('hello')",
		},
		Project_name: "ptrts",
		Language:     "中文",
	}
	invoke, err := r.Invoke(ctx, &schema.Message{
		Extra: map[string]any{
			"PrepInput": prepInput,
		},
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	print(invoke)
}
