package external

import (
	"fmt"
	"homeworkprojet/dto"
	"homeworkprojet/model"
	"math/rand"
	"strconv"
	"time"
)

type CardSchemeExternalSvc interface {
	Validate(cardSchemeDto *dto.CardSchemeDto) bool
	GetPayments(cardScheme *model.CardScheme) []dto.PaymentStatsDto
}

var vSvc *visaSvc
var mSvc *mastercardSvc

const Visa = "VISA"
const MasterCard = "MASTERCARD"

func GetExternalService(schemeType string) CardSchemeExternalSvc {
	if schemeType == Visa {
		if vSvc == nil {
			vSvc = &visaSvc{}
		}
		return vSvc
	} else if schemeType == MasterCard {
		if mSvc == nil {
			mSvc = &mastercardSvc{}
		}
		return mSvc
	}
	return nil
}

var sellers = []string{"Mcdo", "Auchan", "BioCop", "Décathlon", "Irish pub", "Carefour market"}

func mockedPayments() []dto.PaymentStatsDto {
	payments := make([]dto.PaymentStatsDto, 0)
	for i := 0; i < 15; i++ {
		amount, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", rand.Float64()*100), 64)
		p := dto.PaymentStatsDto{
			Date:   randate(),
			Amount: amount,
			Seller: sellers[rand.Intn(len(sellers))],
		}
		payments = append(payments, p)
	}
	return payments
}

func randate() time.Time {
	min := time.Date(2019, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
