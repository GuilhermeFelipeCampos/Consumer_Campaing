package service

import (
	"context"

	"main.go/model"
	"main.go/repository"
)

type RegionServiceInterface interface {
	AddNewRegion(ctx context.Context, region model.Region) error
}

type RegionService struct {
	regionRepository repository.RegionRepositoryInterface
}

func NewRegionService(regionRepository repository.RegionRepositoryInterface) RegionService {
	return RegionService{
		regionRepository: regionRepository,
	}
}

func (r RegionService) AddNewRegion(ctx context.Context, region model.Region) error {

	err := r.regionRepository.InsertRegion(ctx, model.Region{
		Id:        region.Id,
		Name:      region.Name,
		Status:    region.Status,
		Latitude:  region.Latitude,
		Longitude: region.Longitude,
		Cost:      region.Cost,
		CreatedBy: region.CreatedBy,
		UpdatedBy: region.UpdatedBy,
		CreatedAt: region.CreatedAt,
		UpdatedAt: region.UpdatedAt,
	})

	if err != nil {
		return err
	}
	return nil
}
