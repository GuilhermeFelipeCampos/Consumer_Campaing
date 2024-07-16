package model

import (
	"time"

	"github.com/google/uuid"
)

type Merchant struct {
	Id        uuid.UUID `json:"id"`
	OwnerId   uuid.UUID `json:"owner_id"`
	RegionId  uuid.UUID `json:"region_id"`
	Slug      []string  `json:"slug"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
