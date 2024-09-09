package responses

type InstallmentResponse struct {
	Status             string              `json:"status"`
	Locale             string              `json:"locale"`
	SystemTime         int                 `json:"systemTime"`
	ConversationId     string              `json:"conversationId"`
	InstallmentDetails []InstallmentDetail `json:"installmentDetails"`
}

type InstallmentDetail struct {
	BinNumber string  `json:"binNumber"`
	Price     float32 `json:"price"`

	// * CREDIT_CARD,
	// * DEBIT_CARD,
	// * PREPAID_CARD
	CardType string `json:"cardType"`

	// * TROY,
	// * VISA,
	// * MASTER_CARD,
	// * AMERICAN_EXPRESS
	CardAssociation string `json:"cardAssociation"`

	// * Bonus,
	// * Axess,
	// * World,
	// * Maximum,
	// * Paraf,
	// * CardFinans,
	// * Advantage
	CardFamilyName string `json:"cardFamilyName"`

	Force3ds          int                `json:"force3ds"`
	BankCode          int                `json:"bankCode"`
	BankName          string             `json:"bankName"`
	ForceCvc          int                `json:"forceCvc"`
	Commercial        int                `json:"commercial"`
	InstallmentPrices []InstallmentPrice `json:"installmentPrices"`
}

type InstallmentPrice struct {
	InstallmentPrice  float32 `json:"installmentPrice"`
	TotalPrice        float32 `json:"totalPrice"`
	InstallmentNumber float32 `json:"installmentNumber"`
}
