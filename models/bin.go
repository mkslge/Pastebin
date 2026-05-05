package models

import (
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Contents  string    `json:"contents"`
	CreatedAt time.Time `json:"createdAt"`
}
