package paystack

import (
	"net/http"
	"net/url"
)

const (
	_GET    = iota
	_POST   = iota
	_DELETE = iota
	_PUT    = iota
	BaseUrl = "https://api.paystack.co"
)

type credentials struct {
	CustomerKey string
	SecretKey   string
}

type Response struct {
	status  string
	message string
	data    interface{}
}

type PaystackApi struct {
	Credentials *credentials
	HttpClient  *http.Client
	baseUrl     string
}

type query struct {
	url    string
	form   url.Values
	data   interface{}
	method int
}

func cleanValues(v url.Values) url.Values {
	if v == nil {
		return url.Values{}
	}
	return v
}

func (ps PaystackApi) NewPaystackApi(CustomerKey string, CustomerSecret string) *PaystackApi {
	c := &PaystackApi{
		Credentials: &credentials{
			CustomerKey: CustomerKey,
			SecretKey:   CustomerSecret,
		},
		HttpClient: http.DefaultClient,
		baseUrl:    BaseUrl,
	}

	return c
}
