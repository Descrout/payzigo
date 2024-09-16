package requests

type Auth3dsRequest struct {
	Locale           string `json:"locale"`
	ConversationID   string `json:"conversationId"`
	PaymentID        string `json:"paymentId"`
	ConversationData string `json:"conversationData"`
}
