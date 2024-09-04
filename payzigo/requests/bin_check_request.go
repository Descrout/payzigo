package requests

type BinCheckRequest struct {
	Locale         string `json:"locale"`
	ConversationId string `json:"conversationId"`
	BinNumber      string `json:"binNumber"`
}
