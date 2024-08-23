package service

import (
	"context"

	"main.go/model"
	"main.go/repository"
)

type SlugServiceInterface interface {
	AddNewSlug(ctx context.Context, slug model.Slug) error
}

type SlugService struct {
	slugRepository repository.SlugRepositoryInterface
}

func NewSlugService(slugRepository repository.SlugRepositoryInterface) SlugService {

	return SlugService{
		slugRepository: slugRepository,
	}
}

func (s SlugService) AddNewSlug(ctx context.Context, slug model.Slug) error {

	err := s.slugRepository.InsertSlug(ctx, model.Slug{
		Id:        slug.Id,
		Name:      slug.Name,
		Status:    slug.Status,
		Cost:      slug.Cost,
		CreatedBy: slug.CreatedBy,
		UpdatedBy: slug.UpdatedBy,
		CreatedAt: slug.CreatedAt,
		UpdatedAt: slug.UpdatedAt,
	})

	if err != nil {
		return err
	}

	return nil
}
