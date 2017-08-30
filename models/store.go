package models

import "time"

type Store interface {
	CreateInvoice(addr PaymentAddr, user User, amount float64) (Invoice, error)
	GetUserInvoice() error
}

type Invoice interface {
	PaymentAddr() PaymentAddr
	User() User
	Amount() float64
	Created() time.Time
}
