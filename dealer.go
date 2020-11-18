package moka

type (
	PaymentDealerAuthentication struct {
		Code     string `json:"DealerCode"`
		Username string `json:"Username"`
		Password string `json:"Password"`
		checkKey string `json:"CheckKey"`
	}
)
