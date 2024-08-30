package service

import (
	"context"

	"main.go/model"
	"main.go/repository"
)

type MerchantServiceInterface interface {
	AddNewMerchant(ctx context.Context, merchant model.Merchant) error
}

type MerchantService struct {
	merchantRepository repository.MerchantRepositoryInterface
}

func NewMerchantService(merchantRepository repository.MerchantRepositoryInterface) MerchantService {

	return MerchantService{
		merchantRepository: merchantRepository,
	}
}

func (m MerchantService) AddNewMerchant(ctx context.Context, merchant model.Merchant) error {

	err := m.merchantRepository.InsertMerchant(ctx, model.Merchant{
		Id:        merchant.Id,
		OwnerId:   merchant.OwnerId,
		RegionId:  merchant.RegionId,
		Slug:      merchant.Slug,
		Name:      merchant.Name,
		Status:    merchant.Status,
		CreatedBy: merchant.CreatedBy,
		UpdatedBy: merchant.UpdatedBy,
		CreatedAt: merchant.CreatedAt,
		UpdatedAt: merchant.UpdatedAt,
	})

	if err != nil {
		return err
	}
	return nil
}
