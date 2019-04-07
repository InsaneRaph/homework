package external

import (
	"homeworkprojet/dto"
	"homeworkprojet/model"
)

type mastercardSvc struct {
}

func (svc mastercardSvc) Validate(cardSchemeDto *dto.CardSchemeDto) bool {
	return true
}

func (svc mastercardSvc) GetPayments(cardScheme *model.CardScheme) []dto.PaymentStatsDto {
	return mockedPayments()
}
