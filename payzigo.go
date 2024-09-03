package payzigo

import (
	"bytes"
	"io"
	"net/http"

	"github.com/Descrout/payzigo/utils"
)

const (
	PROD_URL    = "https://api.iyzipay.com"
	SANDBOX_URL = "https://sandbox-api.iyzipay.com"
)

type PayzigoOptions struct {
	BaseUrl   string
	ApiKey    string
	SecretKey string
}

type Payzigo struct {
	baseUrl   string
	apiKey    string
	secretKey string
	client    *http.Client
}

func WithOptions(options *PayzigoOptions) *Payzigo {
	return &Payzigo{
		baseUrl:   options.BaseUrl,
		apiKey:    options.ApiKey,
		secretKey: options.SecretKey,
		client:    &http.Client{},
	}
}

func (p *Payzigo) makeRequest(method string, endpoint string, data any) ([]byte, error) {
	requestString := utils.GenerateRequestString(data)
	randomString := utils.GenerateRandomString(8)
	pkiString := utils.GeneratePKIString(p.apiKey, randomString, p.secretKey, requestString)
	hashedPki := utils.HashSha1(pkiString)

	req, err := http.NewRequest(method, p.baseUrl+endpoint, bytes.NewBufferString(requestString))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Expect", "100-continue")
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("Cache-Control", "no-cache")

	req.Header.Set("Authorization", utils.GenerateAuthorizationHeader(p.apiKey, hashedPki))
	req.Header.Set("x-iyzi-rnd", randomString)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return rawBody, nil
}
