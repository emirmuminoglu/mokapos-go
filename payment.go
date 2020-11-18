package moka

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func (m *Moka) PaymentWithThreeD(pdr PaymentDealerRequest) (Response, error) {
	return m.doRequest(pdr, "/PaymentDealer/DoDirectPaymentThreeD")
}

func (m *Moka) PaymentNonThreeD(pdr PaymentDealerRequest) (Response, error) {
	return m.doRequest(pdr, "/PaymentDealer/DoDirectPayment")
}

func (m *Moka) Capture(pdr PaymentDealerRequest) (Response, error) {
	return m.doRequest(pdr, "/PaymentDealer/DoCaputre")
}

func (m *Moka) doRequest(pdr PaymentDealerRequest, url string) (response Response, err error) {
	pdr.Software = "github.com/emirmuminoglu/mokapos-go"
	r := &Request{
		PaymentDealerRequest:        pdr,
		PaymentDealerAuthentication: m.Dealer,
	}

	json, err := r.MarshalJSON()
	if err != nil {
		return
	}

	reader := bytes.NewReader(json)

	req, err := http.NewRequest("POST", m.baseURL+url, reader)
	if err != nil {
		return
	}

	resp, err := m.Client.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = response.UnmarshalJSON(b)
	if err != nil {
		return
	}

	err = response.Error()

	return
}
