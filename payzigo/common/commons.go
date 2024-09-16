package common

type Buyer struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	Surname             string `json:"surname"`
	IdentityNumber      string `json:"identityNumber"`
	Email               string `json:"email"`
	GsmNumber           string `json:"gsmNumber"`
	RegistrationAddress string `json:"registrationAddress"`
	City                string `json:"city"`
	Country             string `json:"country"`
	IP                  string `json:"ip"`
}

type Address struct {
	Address     string `json:"address"`
	ContactName string `json:"contactName"`
	City        string `json:"city"`
	Country     string `json:"country"`
}

type BasketItem struct {
	ID        string `json:"id"`
	Price     string `json:"price"`
	Name      string `json:"name"`
	Category1 string `json:"category1"`
	Category2 string `json:"category2,omitempty"`

	// PHYSICAL, VIRTUAL
	ItemType string `json:"itemType"`
}

type PaymentCard struct {
	CardHolderName string `json:"cardHolderName"`
	CardNumber     string `json:"cardNumber"`
	ExpireYear     string `json:"expireYear"`
	ExpireMonth    string `json:"expireMonth"`
	Cvc            string `json:"cvc"`
}
