package model

import "github.com/google/uuid"

type Spent struct {
	Id               uuid.UUID `json:"id"`
	CampaingId       uuid.UUID `json:"campaing_id"`
	Bucket           string    `json:"bucket"`
	TotalSpent       float64   `json:"total_spent"`
	TotalClicks      int       `json:"total_clicks"`
	TotalImpressions int       `json:"total_impressions"`
}
