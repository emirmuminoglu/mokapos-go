package moka

import (
	"errors"
)

var (
	ErrInvalidRequest = errors.New("Invalid CheckKey. Please check your CheckKey")
	ErrInvalidAccount = errors.New("Invalid Account. Please check your dealer code, username and password.")
	ErrVirtualPosNotFound = errors.New("There is no Virtual Pos associated with your account.")
	ErrDailyDealerLimitExceeded = errors.New("Dealer exceeded one of the daily limits.")
	ErrDailyCardLimitExceeded = errors.New("This Credit Card is delimitated.")
	ErrInvalidCardInfo = errors.New("Invalid Card Info.")
	ErrThreeDRequired = errors.New("Dealer can't do Non-3D payment.")
	ErrInstallmentWithForeignCurrency = errors.New("Can't make installments wiht foreign currency.")
	ErrInstallmentNumberForDealer = errors.New("This number of installments is not available for dealer.")
	ErrInvalidInstallmentNumber = errors.New("Installment Number must be between 2 and 12")
)