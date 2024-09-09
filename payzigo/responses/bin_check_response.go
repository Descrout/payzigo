package responses

import "time"

type BinCheckResponse struct {
	Status         string `json:"status"`
	Locale         string `json:"locale"`
	SystemTime     int64  `json:"systemTime"`
	ConversationId string `json:"conversationId"`
	BinNumber      string `json:"binNumber"`

	// CREDIT_CARD, DEBIT_CARD, PREPAID_CARD
	CardType string `json:"cardType"`

	// VISA, MASTER_CARD, AMERICAN_EXPRESS, TROY
	CardAssociation string `json:"cardAssociation"`

	// Bonus, Axess, World, Maximum, Paraf, CardFinans, Advantage
	CardFamily string `json:"cardFamily"`

	BankName   string `json:"bankName"`
	BankCode   int    `json:"bankCode"`
	Commercial int    `json:"commercial"`
}

func (b *BinCheckResponse) GetSystemTime() time.Time {
	return time.Unix(0, b.SystemTime*int64(time.Millisecond))
}
