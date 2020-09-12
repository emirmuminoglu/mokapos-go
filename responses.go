package moka

import (
	jwriter "github.com/mailru/easyjson/jwriter"
	jlexer "github.com/mailru/easyjson/jlexer"
)

type (
	
	ThreeDResponse struct {
		Data string `json:"Data"`
		ResultCode string `json:"ResultCode"`
		ResultMessage string `json:"ResultMessage"`
		Exception string `json:"Exception"`
	}

	NonTDResponse struct {
		Data struct {
			IsSuccessful bool `json:"IsSuccessful"`
			ResultCode string `json:"ResultCode"`
			ResultMessage string `json:"ResultMessage"`
			VirtualPostOrderID string `json:"VirtualPostOrderId"` 
		} `json:"Data"`
		ResultCode string `json:"ResultCode"`
		ResultMessage string `json:"ResultMessage"`
		Exception string `json:"Exception"`
	}

	Response interface{
		Error() error
		MarshalJSON()([]byte, error)
		MarshalEasyJSON(w *jwriter.Writer)
		UnmarshalEasyJSON(l *jlexer.Lexer)
		UnmarshalJSON([]byte) error
	}
)

func (r ThreeDResponse) Error() (err error) {
	err = errorSwitch(r.ResultCode)
	return 
}

func (r NonTDResponse) Error() (err error) {
	err = errorSwitch(r.ResultCode)
	return 
}

func errorSwitch(errCode string) (err error){
	switch errCode{
	case "success":
		err = nil 
	case "PaymentDealer.CheckPaymentDealerAuthentication.InvalidRequest":
		err = ErrInvalidRequest
	case "PaymentDealer.CheckPaymentDealerAuthentication.InvalidAccount":
		err = ErrInvalidAccount
	case "PaymentDealer.CheckPaymentDealerAuthentication.VirtualPosNotFound":
		err = ErrVirtualPosNotFound
	case "PaymentDealer.CheckDealerPaymentLimits.DailyDealerLimitExceeded":
		err = ErrDailyDealerLimitExceeded
	case "PaymentDealer.CheckDealerPaymentLimits.DailyCardLimitExceeded":
		err = ErrDailyCardLimitExceeded
	case "PaymentDealer.CheckCardInfo.InvalidCardInfo":
		err = ErrInvalidCardInfo
	case "PaymentDealer.DoDirectPayment.ThreeDRequired":
		err = ErrThreeDRequired
	case "PaymentDealer.DoDirectPayment.InstallmentNotAvailableForForeignCurrencyTransaction":
		err = ErrInstallmentWithForeignCurrency
	case "PaymentDealer.DoDirectPayment.ThisInstallmentNumberNotAvailableForDealer":
		err = ErrInstallmentNumberForDealer
	case "PaymentDealer.DoDirectPayment.InvalidInstallmentNumber":
		err = ErrInvalidInstallmentNumber
	}

	return
}