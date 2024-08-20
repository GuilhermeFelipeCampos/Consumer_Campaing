package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.ifoodcorp.com.br/easyzap"
	"main.go/model"
)

type MerchantRepositoryInterface interface {
	InsertMerchant(ctx context.Context, merchant model.Merchant) error
	ListMerchants(ctx context.Context) ([]model.Merchant, error)
	GetMerchant(ctx context.Context, id string) (model.Merchant, error)
}

type MerchantRepository struct {
	pgxpool *pgxpool.Pool
}

func NewMerchantRepository(pgxpool *pgxpool.Pool) *MerchantRepository {
	return &MerchantRepository{pgxpool: pgxpool}
}

func (pg MerchantRepository) InsertMerchant(ctx context.Context, merchant model.Merchant) error {
	sql := `INSERT INTO merchant(id, owner_id, region_id, slug, name, status,created_by,updated_by, created_at, updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`

	err := pg.pgxpool.QueryRow(ctx, sql, merchant.Id, merchant.OwnerId, merchant.RegionId, merchant.Name, merchant.Slug,
		merchant.Status, merchant.CreatedBy, merchant.UpdatedBy, merchant.CreatedAt, merchant.UpdatedAt).Scan()
	if err != nil {
		easyzap.Error(context.Background(), err, "error on insert merchant")
		return err
	}

	return nil

}

func (pg MerchantRepository) ListMerchants(ctx context.Context) ([]model.Merchant, error) {

	sql := `SELECT id, owner_id, region_id, slug, name, status,created_by,updated_by, created_at, updated_at FROM merchant`

	rows, err := pg.pgxpool.Query(ctx, sql)
	if err != nil {
		easyzap.Error(context.Background(), err, "error on list merchants")
		return nil, err
	}

	defer rows.Close()
	var merchants []model.Merchant
	for rows.Next() {
		var merchant model.Merchant
		err := rows.Scan(&merchant.Id, &merchant.OwnerId, &merchant.RegionId, &merchant.Slug, &merchant.Name,
			&merchant.Status, &merchant.CreatedBy, &merchant.UpdatedBy, &merchant.CreatedAt, &merchant.UpdatedAt)
		if err != nil {
			easyzap.Error(context.Background(), err, "error on list merchants")
			return nil, err
		}
		merchants = append(merchants, merchant)
	}

	return merchants, nil

}

func (pg MerchantRepository) GetMerchant(ctx context.Context, id string) (model.Merchant, error) {

	sql := `SELECT id, owner_id, region_id, slug, name, status,created_by,updated_by, created_at, updated_at FROM merchant where id = $1`

	var merchant model.Merchant
	err := pg.pgxpool.QueryRow(ctx, sql, id).Scan(&merchant.Id, &merchant.OwnerId, &merchant.RegionId, &merchant.Slug, &merchant.Name,
		&merchant.Status, &merchant.CreatedBy, &merchant.UpdatedBy, &merchant.CreatedAt, &merchant.UpdatedAt)
	if err != nil {
		easyzap.Error(context.Background(), err, "error on get merchant")
		return merchant, err
	}
	return merchant, nil

}
