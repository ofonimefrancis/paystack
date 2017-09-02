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
	BaseURL = "https://api.paystack.co" //BaseURL paystack API URL
)

type credentials struct {
	CustomerKey string
	SecretKey   string
}

// PaystackAPI struct holding the needful
type PaystackAPI struct {
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

func (ps PaystackAPI) getSecretKey() string {
	return ps.Credentials.SecretKey
}

func (ps PaystackAPI) setBasicAuth(r *http.Request) {
	r.Header.Set("Authorization", "Bearer "+ps.Credentials.SecretKey)
	r.Header.Add("Content-Type", "application/json")
}

//NewPaystackAPI initialize a paystack app
func NewPaystackAPI(CustomerKey string, CustomerSecret string) *PaystackAPI {
	c := &PaystackAPI{
		Credentials: &credentials{
			CustomerKey: CustomerKey,
			SecretKey:   CustomerSecret,
		},
		HttpClient: http.DefaultClient,
		baseUrl:    BaseURL,
	}
	return c
}
