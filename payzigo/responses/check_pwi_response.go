package responses

type CheckPWIResponse struct {
	// success, failure
	Status                       string            `json:"status"`
	Locale                       string            `json:"locale"`
	SystemTime                   int               `json:"systemTime"`
	ConversationID               string            `json:"conversationId"`
	Price                        float64           `json:"price"`
	PaidPrice                    float64           `json:"paidPrice"`
	Installment                  int               `json:"installment"`
	PaymentID                    string            `json:"paymentId"`
	FraudStatus                  int               `json:"fraudStatus"`
	MerchantCommissionRate       float64           `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount float64           `json:"merchantCommissionRateAmount"`
	IyziCommissionRateAmount     float64           `json:"iyziCommissionRateAmount"`
	IyziCommissionFee            float64           `json:"iyziCommissionFee"`
	BinNumber                    string            `json:"binNumber"`
	LastFourDigits               string            `json:"lastFourDigits"`
	BasketID                     string            `json:"basketId"`
	Currency                     string            `json:"currency"`
	ItemTransactions             []ItemTransaction `json:"itemTransactions"`
	Phase                        string            `json:"phase"`
	Token                        string            `json:"token"`
	CallbackURL                  string            `json:"callbackUrl"`

	// SUCCESS, FAILURE, INIT_THREEDS, CALLBACK_THREEDS, BKM_POS_SELECTED, CALLBACK_PECCO
	PaymentStatus string `json:"paymentStatus"`
}

type ConvertedPayout struct {
	PaidPrice                     float64 `json:"paidPrice"`
	IyziCommissionRateAmount      float64 `json:"iyziCommissionRateAmount"`
	IyziCommissionFee             float64 `json:"iyziCommissionFee"`
	BlockageRateAmountMerchant    float64 `json:"blockageRateAmountMerchant"`
	BlockageRateAmountSubMerchant float64 `json:"blockageRateAmountSubMerchant"`
	SubMerchantPayoutAmount       float64 `json:"subMerchantPayoutAmount"`
	MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
	IyziConversionRate            float64 `json:"iyziConversionRate"`
	IyziConversionRateAmount      float64 `json:"iyziConversionRateAmount"`
	Currency                      string  `json:"currency"`
}

type ItemTransaction struct {
	ItemID                        string          `json:"itemId"`
	PaymentTransactionID          string          `json:"paymentTransactionId"`
	TransactionStatus             int             `json:"transactionStatus"`
	Price                         float64         `json:"price"`
	PaidPrice                     float64         `json:"paidPrice"`
	MerchantCommissionRate        float64         `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount  float64         `json:"merchantCommissionRateAmount"`
	IyziCommissionRateAmount      float64         `json:"iyziCommissionRateAmount"`
	IyziCommissionFee             float64         `json:"iyziCommissionFee"`
	BlockageRate                  float64         `json:"blockageRate"`
	BlockageRateAmountMerchant    float64         `json:"blockageRateAmountMerchant"`
	BlockageRateAmountSubMerchant float64         `json:"blockageRateAmountSubMerchant"`
	BlockageResolvedDate          string          `json:"blockageResolvedDate"`
	SubMerchantPrice              float64         `json:"subMerchantPrice"`
	SubMerchantPayoutRate         float64         `json:"subMerchantPayoutRate"`
	SubMerchantPayoutAmount       float64         `json:"subMerchantPayoutAmount"`
	MerchantPayoutAmount          float64         `json:"merchantPayoutAmount"`
	ConvertedPayout               ConvertedPayout `json:"convertedPayout"`
}
