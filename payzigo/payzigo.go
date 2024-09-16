package payzigo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Descrout/payzigo/payzigo/requests"
	"github.com/Descrout/payzigo/payzigo/responses"
	"github.com/Descrout/payzigo/payzigo/utils"
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

	reqBody, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, p.baseUrl+endpoint, bytes.NewReader(reqBody))
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

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	errResp := &responses.ErrorResponse{}
	unmarshalErr := json.Unmarshal(rawBody, errResp)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	if errResp.HasError() {
		return nil, errResp.Error()
	}

	return rawBody, nil
}

func (p *Payzigo) CheckInstallments(req *requests.InstallmentRequest) (*responses.InstallmentResponse, error) {
	rawData, err := p.makeRequest("POST", "/payment/iyzipos/installment", *req)
	if err != nil {
		return nil, err
	}

	resp := &responses.InstallmentResponse{}
	err = json.Unmarshal(rawData, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *Payzigo) CheckBin(req *requests.BinCheckRequest) (*responses.BinCheckResponse, error) {
	rawData, err := p.makeRequest("POST", "/payment/bin/check", *req)
	if err != nil {
		return nil, err
	}

	resp := &responses.BinCheckResponse{}
	err = json.Unmarshal(rawData, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *Payzigo) InitPayWithIyzico(req *requests.InitPWIRequest) (*responses.InitPWIResponse, error) {
	rawData, err := p.makeRequest("POST", "/payment/pay-with-iyzico/initialize", *req)
	if err != nil {
		return nil, err
	}

	resp := &responses.InitPWIResponse{}
	err = json.Unmarshal(rawData, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *Payzigo) CheckPayWithIyzico(req *requests.CheckPWIRequest) (*responses.PayCompleteResponse, error) {
	rawData, err := p.makeRequest("POST", "/payment/iyzipos/checkoutform/auth/ecom/detail", *req)
	if err != nil {
		return nil, err
	}

	resp := &responses.PayCompleteResponse{}
	err = json.Unmarshal(rawData, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *Payzigo) Init3ds(req *requests.Init3dsRequest) (*responses.Init3dsResponse, error) {
	rawData, err := p.makeRequest("POST", "/payment/3dsecure/initialize", *req)
	if err != nil {
		return nil, err
	}

	resp := &responses.Init3dsResponse{}
	err = json.Unmarshal(rawData, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *Payzigo) Auth3ds(req *requests.Auth3dsRequest) (*responses.PayCompleteResponse, error) {
	rawData, err := p.makeRequest("POST", "/payment/3dsecure/auth", *req)
	if err != nil {
		return nil, err
	}

	resp := &responses.PayCompleteResponse{}
	err = json.Unmarshal(rawData, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
