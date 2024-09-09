package responses

type InitPWIResponse struct {
	Status               string `json:"status"`
	Locale               string `json:"locale"`
	SystemTime           int64  `json:"systemTime"`
	ConversationID       string `json:"conversationId"`
	Token                string `json:"token"`
	TokenExpireTime      int    `json:"tokenExpireTime"`
	PayWithIyzicoPageURL string `json:"payWithIyzicoPageUrl"`
}
