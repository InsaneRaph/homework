package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

const cardSchemeTable = "card_schemes"

type CardScheme struct {
	gorm.Model
	Number          string
	Holder          string
	UserID          uint
	ExpirationMonth time.Month
	ExpirationYear  int
	Type            string
	User            User
}

func (card *CardScheme) Create() error {

	return DB.Create(card).Error
}

func GetCardSchemesForUser(userId uint) []*CardScheme {

	var cards = make([]*CardScheme, 0)
	DB.Table(cardSchemeTable).Where("user_id = ?", userId).Find(&cards)

	return cards
}

func GetCardSchemeForUser(userId uint, id uint) *CardScheme {
	card := &CardScheme{}
	DB.Table(cardSchemeTable).Where("user_id = ? and id = ?", userId, id).First(card)
	return card
}

func DeleteCardSchemeForUser(userId uint, id uint) error {

	return DB.Table(cardSchemeTable).Where("user_id = ? and id = ?", userId, id).Delete(CardScheme{}).Error
}
