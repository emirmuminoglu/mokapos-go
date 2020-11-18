package moka

type (
	//PaymentDealerRequest represents a part of request that you send to moka services.
	//For detailed documentation: https://developer.moka.com
	PaymentDealerRequest struct {
		VirtualPosOrderID   string              `json:"VirtualPosOrderId,omitempty"`
		CardHolderFullName  string              `json:"CardHolderFullName,omitempty"`
		CardNumber          string              `json:"CardNumber,omitempty"`
		ExpMonth            string              `json:"ExpMonth,omitempty"`
		ExpYear             string              `json:"ExpYear,omitempty"`
		CvcNumber           string              `json:"CvcNumber,omitempty"`
		Amount              uint                `json:"Amount,omitempty"`
		Currency            string              `json:"Currency,omitempty"`
		InstallmentNumber   int8                `json:"InstallmentNumber,omitempty"`
		ClientIP            string              `json:"ClientIP,omitempty"`
		RedirectURL         string              `json:"RedirectUrl,omitempty"`
		RedirectType        int8                `json:"RedirectType,omitempty"`
		OtherTrxCode        string              `json:"OtherTrxCode,omitempty"`
		IsPreAuth           int8                `json:"IsPreAuth,omitempty"`
		IsPoolPayment       int8                `json:"IsPoolPayment,omitempty"`
		IsTokenized         int8                `json:"IsTokenized,omitempty"`
		IntegratorID        int8                `json:"IntegratorId,omitempty"`
		Software            string              `json:"Software,omitempty"`
		SubMerchantName     string              `json:"SubMerchantName,omitempty"`
		Description         string              `json:"Description,omitempty"`
		BuyerInformation    BuyerInformation    `json:"BuyerInformation,omitempty"`
		CustomerInformation CustomerInformation `json:"CustomerInformation,omitempty"`
	}

	//BuyerInformation represents
	BuyerInformation struct {
		FullName  string `json:"BuyerFullName,omitempty"`
		EMail     string `json:"BuyerEmail,omitempty"`
		GSMNumber string `json:"BuyerGsmNumber,omitempty"`
		Address   string `json:"BuyerAddress,omitempty"`
	}

	CustomerInformation struct {
		DealerCustomerID int    `json:"DealerCustomerId,omitempty"`
		CustomerCode     string `json:"CustomerCode,omitempty"`
		FirstName        string `json:"FirstName,omitempty"`
		LastName         string `json:"LastName,omitempty"`
		Gender           int8   `json:"Gender,omitempty"`
		BirthDate        string `json:"BirthDate,omitempty"`
		GsmNumber        string `json:"GsmNumber,omitempty"`
		EMail            string `json:"Email,omitempty"`
		Address          string `json:"Address,omitempty"`
	}

	Request struct {
		PaymentDealerAuthentication PaymentDealerAuthentication `json:"PaymentDealrAuthentication"`
		PaymentDealerRequest        PaymentDealerRequest        `json:"PaymentDealerRequest"`
	}
)
