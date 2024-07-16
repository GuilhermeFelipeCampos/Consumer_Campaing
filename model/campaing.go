package model

import (
	"time"

	"github.com/google/uuid"
)

type Campaing struct {
	Id         uuid.UUID `json:"id"`
	MerchantId uuid.UUID `json:"merchant_id"`
	Status     string    `json:"status"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Budget     float32   `json:"budget"`
	CreatedBy  string    `json:"created_by"`
	UpdatedBy  string    `json:"updated_by"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
