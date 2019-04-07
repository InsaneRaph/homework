package dto

import (
	"time"
)

type PaymentStatsDto struct {
	Date         time.Time `json:"date,omitempty"`
	Amount       float64   `json:"amount,omitempty"`
	Seller       string    `json:"seller"`
}