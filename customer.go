package paystack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func (ps PaystackAPI) GetCutomers(u url.Values) (*CustomerResponse, error) {
	v := cleanValues(u)
	r, err := http.NewRequest(http.MethodGet, ps.baseUrl+"/customer", bytes.NewBufferString(v.Encode()))
	ps.setBasicAuth(r)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	resp, err := ps.HttpClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return getCustomers([]byte(body))

	//decodeResponse
	/**TODO: Check response code based on paystack API and return appropriate error **/

}

func (ps PaystackAPI) GetCustomerById(id int) {
	r, err := http.NewRequest(http.MethodGet, ps.baseUrl+"/customer/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	ps.setBasicAuth(r)
	res, err := ps.HttpClient.Do(r)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer res.Body.Close()

	/**TODO: Decode into the appropriate Customer object **/

}

func (ps PaystackAPI) CreateCustomer(u url.Values) {
	v := cleanValues(u)
	r, err := http.NewRequest(http.MethodPost, ps.baseUrl+"/customer", bytes.NewBufferString(v.Encode()))
	ps.setBasicAuth(r)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		log.Fatal(err.Error())
	}

	resp, err := ps.HttpClient.Do(r)
	//TODO: handle response code and spit out the error
	defer resp.Body.Close()

	fmt.Println(resp.Request)
	fmt.Println(resp.Status)

	/**body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return getNewCustomer([]byte(body))**/

}

func getCustomers(body []byte) (*CustomerResponse, error) {
	var c = new(CustomerResponse)
	err := json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}
	return c, err
}

func getNewCustomer(body []byte) (*NewCustomerResponse, error) {
	var c = new(NewCustomerResponse)
	err := json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}
	return c, err
}

func (ps PaystackAPI) UpdateCustomer()                        {}
func (ps PaystackAPI) GetCustomerByCode(customer_code string) {}

func (ps PaystackAPI) DeactivateAuthorization() {}
