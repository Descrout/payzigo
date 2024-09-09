package requests

type CheckPWIRequest struct {
	Locale         string `json:"locale"`
	ConversationID string `json:"conversationId"`
	Token          string `json:"token"`
}
