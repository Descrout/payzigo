package requests

import "github.com/Descrout/payzigo/payzigo/common"

type Init3dsRequest struct {
	Locale         string `json:"locale"`
	ConversationID string `json:"conversationId"`
	Price          string `json:"price"`
	PaidPrice      string `json:"paidPrice"`
	Installment    int    `json:"installment"`

	BasketID string `json:"basketId"`

	// PRODUCT, LISTING, SUBSCRIPTION
	PaymentGroup string `json:"paymentGroup"`

	PaymentCard common.PaymentCard `json:"paymentCard"`

	Buyer           common.Buyer        `json:"buyer"`
	ShippingAddress common.Address      `json:"shippingAddress"`
	BillingAddress  common.Address      `json:"billingAddress"`
	BasketItems     []common.BasketItem `json:"basketItems"`

	// SHOPIFY, MAGENTO, PRESTASHOP, WOOCOMMERCE, OPENCART
	PaymentSource string `json:"paymentSource"`
	Currency      string `json:"currency"`

	CallbackURL string `json:"callbackUrl"`
}
