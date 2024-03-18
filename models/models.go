package models

import (
	"time"
)

type User struct {
	ID        uint
	Username  string
	Email     string
	Age       uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}
