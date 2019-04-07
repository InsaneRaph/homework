package service

import (
	"errors"
	"homeworkprojet/dto"
	"homeworkprojet/external"
	"homeworkprojet/model"
	"strings"
)

type CardSchemeSvc interface {
	Create(cardSchemeDto *dto.CardSchemeDto, userId uint) (*dto.CardSchemeDto, error)
	GetCardSchemesForUser(userId uint) []*dto.CardSchemeDto
	DeleteCardSchemeForUser(userId uint, id uint) error
}
type carSchemeSvc struct {
}

var csSvc *carSchemeSvc

func GetCardSchemeService() CardSchemeSvc {
	if csSvc == nil {
		csSvc = &carSchemeSvc{}
	}
	return csSvc
}



func (svc carSchemeSvc) Create(cardSchemeDto *dto.CardSchemeDto, userId uint) (*dto.CardSchemeDto, error) {
	schemeType := strings.ToUpper(cardSchemeDto.Type)
	if !(schemeType == external.Visa || schemeType == external.MasterCard) {
		return nil, errors.New("unsupported scheme type")
	}

	if !external.GetExternalService(schemeType).Validate(cardSchemeDto){
		return nil, errors.New("invalid card information")
	}

	cardScheme := cardSchemeDto.ToCardScheme()
	cardScheme.UserID = userId
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
