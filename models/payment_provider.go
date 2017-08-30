package models

import "time"

type Provider interface {
	GetFreePaymentAddr() (PaymentAddr, error)
	LockPaymentAddr(PaymentAddr, time.Duration) error
}

type PaymentAddr interface {
	String() string
}
