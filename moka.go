package moka

import (
	"encoding/hex"
	"crypto/sha256"
	"bytes"
	"net/http"
)

type (
	Moka struct {
		baseURL string
		Dealer PaymentDealerAuthentication
		Client *http.Client
	}
)

func New(code, username, password string, isTest bool) (m *Moka) {
	m = new(Moka)

	url := "https://service.moka.com"		

	if isTest{
		url = "https://service.testmoka.com" 
	}
	
	m.baseURL = url
	m.Dealer = PaymentDealerAuthentication{
		Code: code,
		Username: username,
		Password: password,
	}

	m.createCheckKey()
	return
}

func (m *Moka) createCheckKey() (str string) {
	var buffer bytes.Buffer

	buffer.WriteString(m.Dealer.Code)
	buffer.WriteString("MK")
	buffer.WriteString(m.Dealer.Username)
	buffer.WriteString("PD")
	buffer.WriteString(m.Dealer.Password)

	b := buffer.Bytes()
	
	hash := sha256.Sum256(b)

	str = hex.EncodeToString(hash[:])

	return
}

