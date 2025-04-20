package internal

type ChatRequestBody struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

type RequestBody struct {
	ID       string     `json:"id"`
	Messages []Messages `json:"messages"`
}
type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Parts   []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"parts"`
}
