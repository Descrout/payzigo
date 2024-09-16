package main

import "github.com/Descrout/payzigo/payzigo/common"

var Buyers = []common.Buyer{
	{
		ID:                  "b12345",
		Name:                "John",
		Surname:             "Doe",
		IdentityNumber:      "13063839722",
		Email:               "john.doe@example.com",
		GsmNumber:           "5437842590",
		RegistrationAddress: "123 Elm Street",
		City:                "Metropolis",
		Country:             "Countryland",
		IP:                  "192.168.1.1",
	},
	{
		ID:                  "b67890",
		Name:                "Alice",
		Surname:             "Smith",
		IdentityNumber:      "13063839722",
		Email:               "alice.smith@example.com",
		GsmNumber:           "5437842590",
		RegistrationAddress: "456 Oak Avenue",
		City:                "Gotham",
		Country:             "Countryland",
		IP:                  "192.168.1.2",
	},
}

var Cards = []common.PaymentCard{
	{
		CardHolderName: "Adil Basar",
		CardNumber:     "5890040000000016",
		ExpireYear:     "2028",
		ExpireMonth:    "02",
		Cvc:            "200",
	},
	{
		CardHolderName: "Adil Basar",
		CardNumber:     "5526080000000006",
		ExpireYear:     "2028",
		ExpireMonth:    "02",
		Cvc:            "200",
	},
}

var Addresses = []common.Address{
	{
		Address:     "789 Maple Road",
		ContactName: "John Doe",
		City:        "Metropolis",
		Country:     "Countryland",
	},
	{
		Address:     "101 Pine Street",
		ContactName: "Alice Smith",
		City:        "Gotham",
		Country:     "Countryland",
	},
}

var BasketItems = []common.BasketItem{
	{
		ID:        "i98765",
		Price:     "29.99",
		Name:      "Wireless Mouse",
		Category1: "Electronics",
		ItemType:  "PHYSICAL",
	},
	{
		ID:        "i54321",
		Price:     "89.99",
		Name:      "%20 Discount Coupon",
		Category1: "Games",
		ItemType:  "VIRTUAL",
	},
}
