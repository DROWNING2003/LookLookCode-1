package workflow

import (
	"context"
)

type AnalysisAgent struct {
}

func NewAnalysisAgent(ctx context.Context, baseURL, model string) (*AnalysisAgent, error) {
	// chatModel, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
	// 	BaseURL: baseURL,
	// 	Model:   model,

	// 	Options: &api.Options{
	// 		Temperature: 0.000001,
	// 	},
	// })
	// if err != nil {
	// 	return nil, err
	// }
	//identifynode
	// chain := compose.NewChain[[]*schema.Message, *schema.Message]()
	// chain.AppendChatTemplate()
	// chain.AppendChatModel(chatModel)
	// identifyAbstractionsNode, err := node.NewIdentifyAbstractionsNode(ctx)
	// chain.AppendChatModel(identifyAbstractionsNode)
	// chain.AppendLambda(compose.InvokableLambda(func(ctx context.Context, input []*schema.Message) ([]*schema.Message, error) {
	// 	systemMsg := &schema.Message{
	// 		Role:    schema.System,
	// 		Content: "You are responsible for preparing the user query for insertion into journal. The user's query is expected to contain the actual text the user want to write to journal, as well as convey the intention that this query should be written to journal. You job is to remove that intention from the user query, while preserving as much as possible the user's original query, and output ONLY the text to be written into journal",
	// 	}
	// 	return append([]*schema.Message{systemMsg}, input...), nil
	// }))
	return &AnalysisAgent{}, nil
}
