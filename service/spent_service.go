package service

import (
	"context"

	"main.go/model"
	"main.go/repository"
)

type SpentServiceInterface interface {
	AddNewSpent(ctx context.Context, spent model.Spent) error
}

type SpentService struct {
	spentRepository repository.SpentRepositoryInterface
}

func NewSpentService(spentRepository repository.SpentRepositoryInterface) SpentService {
	return SpentService{
		spentRepository: spentRepository,
	}
}

func (s SpentService) AddNewSpent(ctx context.Context, spent model.Spent) error {

	err := s.spentRepository.InsertSpent(ctx, model.Spent{
		Id:               spent.Id,
		CampaingId:       spent.CampaingId,
		Bucket:           spent.Bucket,
		TotalSpent:       spent.TotalSpent,
		TotalClicks:      spent.TotalClicks,
		TotalImpressions: spent.TotalImpressions,
	})
	if err != nil {
		return err
	}
	return nil
}
