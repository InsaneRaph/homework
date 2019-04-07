package dto

import (
	"homeworkprojet/model"
	"strings"
	"time"
)

const mask = " **** **** ****"

type CardSchemeDto struct {
	ID              uint       `json:"id,omitempty"`
	CreatedAt       time.Time  `json:"createdAt,omitempty"`
	UpdatedAt       time.Time  `json:"updatedAt,omitempty"`
	Number          string     `json:"number"`
	Holder          string     `json:"holder"`
	Crypto          string     `json:"crypto,omitempty"`
	ExpirationMonth time.Month `json:"expirationMonth"`
	ExpirationYear  int        `json:"expirationYear"`
	Type            string     `json:"type,omitempty"`
}

func FromCardScheme(card *model.CardScheme) *CardSchemeDto {

	return &CardSchemeDto{
		ID:              card.ID,
		CreatedAt:       card.CreatedAt,
		UpdatedAt:       card.UpdatedAt,
		Number:          card.Number[:4] + mask,
		Holder:          card.Holder,
		ExpirationMonth: card.ExpirationMonth,
		ExpirationYear:  card.ExpirationYear,
		Type:            card.Type,
	}
}

func (dto CardSchemeDto) ToCardScheme() model.CardScheme {
	return model.CardScheme{
		Number:          dto.Number,
		Holder:          dto.Holder,
		ExpirationMonth: dto.ExpirationMonth,
		ExpirationYear:  dto.ExpirationYear,
		Type:            strings.ToUpper(dto.Type),
	}
}
