package service

import (
	"context"

	"main.go/model"
	"main.go/repository"
)

type CampaingServiceInterface interface {
	AddNewCampaing(ctx context.Context, campaing model.Campaing) error
}

type CampaingService struct {
	campaingRepository repository.CampaingRepositoryInterface
}

func NewCampaingService(campaingRepository repository.CampaingRepository) CampaingService {

	return CampaingService{
		campaingRepository: campaingRepository,
	}
}

func (c CampaingService) AddNewCampaing(ctx context.Context, campaing model.Campaing) error {

	err := c.campaingRepository.InsertCampaing(ctx, model.Campaing{
		Id:         campaing.Id,
		MerchantId: campaing.MerchantId,
		Status:     campaing.Status,
		Latitude:   campaing.Latitude,
		Longitude:  campaing.Longitude,
		Budget:     campaing.Budget,
		CreatedBy:  campaing.CreatedBy,
		UpdatedBy:  campaing.UpdatedBy,
		CreatedAt:  campaing.CreatedAt,
		UpdatedAt:  campaing.UpdatedAt,
	})

	if err != nil {
		return err
	}
	return nil
}
