package main

import (
	"log"
	"os"

	"github.com/Descrout/payzigo/payzigo"
	"github.com/Descrout/payzigo/payzigo/requests"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env could not be loaded: ", err)
	}

	cli := payzigo.WithOptions(&payzigo.PayzigoOptions{
		BaseUrl:   payzigo.SANDBOX_URL,
		ApiKey:    os.Getenv("API_KEY"),
		SecretKey: os.Getenv("SECRET_KEY"),
	})

	// installments, err := cli.CheckInstallments(&requests.InstallmentRequest{
	// 	Locale:         "tr",
	// 	ConversationId: "1",
	// 	BinNumber:      "454359",
	// 	Price:          "2380.0",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(installments)

	// binCheck, err := cli.CheckBin(&requests.BinCheckRequest{
	// 	Locale:         "tr",
	// 	ConversationId: "1",
	// 	BinNumber:      "589283",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(binCheck)

	pwiCheck, err := cli.InitPayWithIyzico(&requests.InitPWIRequest{
		Locale:         "tr",
		ConversationID: "2",
		Price:          "119.98",
		BasketID:       "2",
		PaymentGroup:   "PRODUCT",
		CallbackURL:    "https://webhook.site/baf81438-06d8-4dba-ad6d-22e94c6ce3b8",
		Currency:       "TRY",
		PaidPrice:      "119.98",
		EnabledInstallments: []int{
			2,
		},
		Buyer:           Buyers[0],
		ShippingAddress: Addresses[0],
		BillingAddress:  Addresses[0],
		BasketItems:     BasketItems,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(pwiCheck)
}
