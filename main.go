package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	port := ":8888"

	pwiInit, err := cli.InitPayWithIyzico(&requests.InitPWIRequest{
		Locale:         "tr",
		ConversationID: "2",
		Price:          "119.98",
		BasketID:       "2",
		PaymentGroup:   "PRODUCT",
		CallbackURL:    "http://localhost" + port + "/payconfirm",
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
	log.Println(pwiInit)

	server := &http.Server{
		Addr:    port,
		Handler: initRoutes(cli),
	}
	defer server.Shutdown(context.TODO())

	// Serve on a seperate goroutine
	log.Println("Listening on port:", port)
	go func() {
		server.ListenAndServe()
		log.Println("Server shutdown gracefully.")
	}()

	// Wait for any closing signals
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, os.Interrupt, syscall.SIGQUIT)
	<-s
	log.Println("Shutting down...")
}

func initRoutes(cli *payzigo.Payzigo) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/payconfirm", func(w http.ResponseWriter, r *http.Request) {
		token := r.FormValue("token")

		pwiCheck, err := cli.CheckPayWithIyzico(&requests.CheckPWIRequest{
			ConversationID: "",
			Locale:         "tr",
			Token:          token,
		})
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		log.Println(pwiCheck)

		json.NewEncoder(w).Encode(pwiCheck)
	})

	mux.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]any{}
		json.NewDecoder(r.Body).Decode(&data)

		for key, value := range data {
			log.Println(key, value)
		}
	})

	return mux
}
