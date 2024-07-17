package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.ifoodcorp.com.br/easyzap"
	"main.go/model"
)

type SlugRepositoryInterface interface {
	InsertSlug(ctx context.Context, slug model.Slug) error
	ListSlugs(ctx context.Context) ([]model.Slug, error)
	SearchSlug(ctx context.Context, id string) (model.Slug, error)
}

type SlugRepository struct {
	pgxpool *pgxpool.Pool
}

func NewSlugRepository(pgxpool *pgxpool.Pool) *SlugRepository {
	return &SlugRepository{pgxpool: pgxpool}
}

func (pg SlugRepository) InsertSlug(ctx context.Context, slug model.Slug) error {
	sql := `INSERT INTO slug(id, name, status, cost, created_by, updated_by, created_at, updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8)`

	err := pg.pgxpool.QueryRow(ctx, sql, slug.Id, slug.Name, slug.Status, slug.Cost, slug.CreatedBy, slug.UpdatedBy, slug.CreatedAt, slug.UpdatedAt).Scan()

	if err != nil {
		easyzap.Error(context.Background(), err, "error on insert slug")
		return err
	}
	return nil
}

func (pg SlugRepository) ListSlugs(ctx context.Context) ([]model.Slug, error) {
	sql := `SELECT id, name, status, cost, created_by, updated_by, created_at, updated_at FROM slug`

	rows, err := pg.pgxpool.Query(ctx, sql)

	if err != nil {
		easyzap.Error(context.Background(), err, "error on list slugs")
		return nil, err
	}
	defer rows.Close()
	var slugs []model.Slug
	for rows.Next() {
		var slug model.Slug
		err := rows.Scan(&slug.Id, &slug.Name, &slug.Status, &slug.Cost, &slug.CreatedBy, &slug.UpdatedBy, &slug.CreatedAt, &slug.UpdatedAt)
		if err != nil {
			easyzap.Error(context.Background(), err, "error on list slugs")
			return nil, err
		}
		slugs = append(slugs, slug)
	}
	return slugs, nil
}

func (pg SlugRepository) SearchSlug(ctx context.Context, id string) (model.Slug, error) {
	sql := `SELECT id, name, status, cost, created_by, updated_by, created_at, updated_at FROM slug WHERE id=$1`

	var slug model.Slug
	err := pg.pgxpool.QueryRow(ctx, sql, id).Scan(&slug.Id, &slug.Name, &slug.Status, &slug.Cost, &slug.CreatedBy, &slug.UpdatedBy, &slug.CreatedAt, &slug.UpdatedAt)
	if err != nil {
		easyzap.Error(context.Background(), err, "error on search slug")
		return slug, err
	}
	return slug, nil
}
