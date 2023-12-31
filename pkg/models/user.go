package models

import (
	"time"
)

type User struct {
	ID           uint       `json:"id" gorm:"primary_key"`
	Name         string     `json:"name" validate:"required"`
	Email        *string    `json:"email,omitempty" validate:"required"`
	Age          uint8      `json:"age" validate:"required"`
	Birthday     *time.Time `json:"birthday,omitempty"`
	MemberNumber *string    `json:"member_number,omitempty"`
	ActivatedAt  *time.Time `json:"activated_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
