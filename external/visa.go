package external

import (
	"homeworkprojet/dto"
	"homeworkprojet/model"
)

type visaSvc struct {

}

func (svc visaSvc) Validate(cardSchemeDto *dto.CardSchemeDto) bool {
	return true
}

func (svc visaSvc) GetPayments(cardScheme *model.CardScheme) []dto.PaymentStatsDto {
	return mockedPayments()
}