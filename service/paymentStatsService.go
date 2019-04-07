package service

import (
	"homeworkprojet/dto"
	"homeworkprojet/external"
	"homeworkprojet/model"
)

type PaymentStatsSvc interface {
	GetPayments(userId uint, cardId uint) []dto.PaymentStatsDto
}
type paymentStatsSvc struct {
}

var psSvc *paymentStatsSvc

func GetPaymentStatsService() PaymentStatsSvc {
	if psSvc == nil {
		psSvc = &paymentStatsSvc{}
	}
	return psSvc
}


func (svc paymentStatsSvc) GetPayments(userId uint, cardId uint) []dto.PaymentStatsDto {
	card := model.GetCardSchemeForUser(userId, cardId)
	return external.GetExternalService(card.Type).GetPayments(card)
}
