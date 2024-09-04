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

	installments, err := cli.CheckInstallments(&requests.InstallmentRequest{
		Locale:         "tr",
		ConversationId: "1",
		BinNumber:      "454359",
		Price:          "2380.0",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(installments)

	binCheck, err := cli.CheckBin(&requests.BinCheckRequest{
		Locale:         "tr",
		ConversationId: "1",
		BinNumber:      "589283",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(binCheck)
}
