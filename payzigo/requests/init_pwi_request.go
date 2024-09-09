package requests

import "github.com/Descrout/payzigo/payzigo/common"

type InitPWIRequest struct {
	Locale         string `json:"locale"`
	ConversationID string `json:"conversationId"`
	Price          string `json:"price"`
	BasketID       string `json:"basketId"`

	// PRODUCT, LISTING, SUBSCRIPTION
	PaymentGroup string `json:"paymentGroup"`

	Buyer               common.Buyer        `json:"buyer"`
	ShippingAddress     common.Address      `json:"shippingAddress"`
	BillingAddress      common.Address      `json:"billingAddress"`
	BasketItems         []common.BasketItem `json:"basketItems"`
	CallbackURL         string              `json:"callbackUrl"`
	Currency            string              `json:"currency"`
	PaidPrice           string              `json:"paidPrice"`
	EnabledInstallments []int               `json:"enabledInstallments"`
}
