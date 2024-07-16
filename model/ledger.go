package model

import (
	"time"

	"github.com/google/uuid"
)

type Ledger struct {
	Id         uuid.UUID `json:"id"`
	CampaingId uuid.UUID `json:"campaing_id"`
	SpentId    uuid.UUID `json:"spent_id"`
	SlugId     uuid.UUID `json:"slug_id"`
	UserId     uuid.UUID `json:"user_id"`
	EventType  string    `json:"event_type"`
	Latitude   float32   `json:"latitude"`
	Longitude  float32   `json:"longitude"`
	Cost       float32   `json:"cost"`
	CreatedAt  time.Time `json:"created_at"`
}
