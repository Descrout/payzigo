package requests

type InstallmentRequest struct {
	Locale         string `json:"locale"`
	ConversationId string `json:"conversationId"`
	BinNumber      string `json:"binNumber"`
	Price          string `json:"price"`
}
