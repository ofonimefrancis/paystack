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
)

func (ps PaystackApi) GetCutomers(u url.Values) (*CustomerResponse, error) {
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
func (ps PaystackApi) GetCustomerById(id int32)               {}
func (ps PaystackApi) GetCustomerByCode(customer_code string) {}

func (ps PaystackApi) CreateCustomer()          {}
func (ps PaystackApi) UpdateCustomer()          {}
func (ps PaystackApi) DeactivateAuthorization() {}

func getCustomers(body []byte) (*CustomerResponse, error) {
	var c = new(CustomerResponse)
	err := json.Unmarshal(body, &c)
	if err != nil {
		fmt.Println("whoops: ", err)
	}
	return c, err
}
