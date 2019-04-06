package service

import (
	"errors"
	"homeworkprojet/dto"
	"homeworkprojet/model"
	"strings"
)

type cardSchemeSvc interface {
	Create(cardSchemeDto *dto.CardSchemeDto) (*dto.CardSchemeDto, error)
	GetCardSchemesForUser(userId uint) []*dto.CardSchemeDto
	DeleteCardSchemeForUser(userId uint, id uint) error
}
type carSchemeSvc struct {
}

var csSvc *carSchemeSvc

func GetCardSchemeService() cardSchemeSvc {
	if csSvc == nil {
		csSvc = &carSchemeSvc{}
	}
	return csSvc
}

const visa = "VISA"
const masterCard = "MASTERCARD"

func (svc carSchemeSvc) Create(cardSchemeDto *dto.CardSchemeDto) (*dto.CardSchemeDto, error) {
	schemeType := strings.ToUpper(cardSchemeDto.Type)
	if !(schemeType == visa || schemeType == masterCard) {
		return nil, errors.New("unsupported scheme type")
	}

	if schemeType == visa {
		//TODO validate with visa external service
	} else {
		//TODO validate with mastercard external service
	}

	cardScheme := cardSchemeDto.ToCardScheme()
	err := cardScheme.Create()

	if err != nil {
		return nil, errors.New("failed to create cad Scheme, connection error")
	}

	return dto.FromCardScheme(&cardScheme), nil
}

func (svc carSchemeSvc) GetCardSchemesForUser(userId uint) []*dto.CardSchemeDto {

	var cards = make([]*dto.CardSchemeDto, 0)
	for _, c := range model.GetCardSchemesForUser(userId) {
		cards = append(cards, dto.FromCardScheme(c))
	}

	return cards
}

func (svc carSchemeSvc) DeleteCardSchemeForUser(userId uint, id uint) error {
	return model.DeleteCardSchemeForUser(userId, id)
}
