package entity

import "errors"

type Order struct {
	ID        string     `json:"id"`
	PersonId  int32      `json:"person_id"`
	ProductId int32      `json:"product_id"`
	Qty       int32      `json:"qty"`
	State     OrderState `json:"state"`
}

type OrderState int

const (
	StateWaitForPayment OrderState = iota
	StateProcessing
	StateDelivery
	StateFinish
	StateCanceled
)

var ErrPaid = errors.New("order has been paid")
var ErrNotPaid = errors.New("order has not been paid")
var ErrProcessed = errors.New("order has been processed")
var ErrNotProcessed = errors.New("order has not been processed")
var ErrDelivered = errors.New("order has been delivered")
var ErrNotDelivered = errors.New("order has not been delivered")
var ErrFinished = errors.New("order has been finished")
var ErrNotFinished = errors.New("order has not been finished")
var ErrCanceled = errors.New("order has been canceled")
var ErrNotCanceled = errors.New("order has not been canceled")
