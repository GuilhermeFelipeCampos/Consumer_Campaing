package service

import (
	"context"

	"main.go/model"
	"main.go/repository"
)

type OnwerServiceInterface interface {
	AddNewOnwer(ctx context.Context, owner model.Owner) error
}

type OnwerService struct {
	onwerRepository repository.OwnerRepository
}

func NewOwnerService(onwerRepository repository.OwnerRepository) OnwerService {
	return OnwerService{
		onwerRepository: onwerRepository,
	}
}

func (o OnwerService) AddNewOnwer(ctx context.Context, owner model.Owner) error {

	err := o.onwerRepository.InsertOwner(ctx, model.Owner{
		Id:        owner.Id,
		Email:     owner.Email,
		Status:    owner.Status,
		CreatedBy: owner.CreatedBy,
		UpdatedBy: owner.UpdatedBy,
		CreatedAt: owner.CreatedAt,
		UpdatedAt: owner.UpdatedAt,
	})
	if err != nil {

		return err
	}
	return nil
}
