package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"ai/internal/chat"
	"ai/pkg/utils"
)

func main() {
	ctx := context.Background()
	chatService, err := chat.NewChatService(ctx)
	if err != nil {
		fmt.Printf("Failed to initialize chat service: %v\n", err)
		return
	}

	messages := []chat.Message{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(utils.UserStyle.Render("You: "))
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if userInput == "" {
			continue
		}
		if userInput == "exit" {
			break
		}

		fmt.Print(utils.AIStyle.Render("AI: "))
		response, err := chatService.GetResponse(ctx, messages, userInput)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		messages = append(messages,
			chat.Message{Role: "user", Content: userInput},
			chat.Message{Role: "assistant", Content: response},
		)
		fmt.Println()
	}
}
