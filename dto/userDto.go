package dto

import (
	"homeworkprojet/model"
	"time"
)

type UserDto struct {
	ID        uint      `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	Token     string    `json:"token,omitempty"`
}

func FromUser(user *model.User) *UserDto {
	return &UserDto{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
		Password:  user.Password,
	}
}

func (dto UserDto) ToUser() model.User {
	return model.User{
		Email:     dto.Email,
		Password:  dto.Password,
	}
}
