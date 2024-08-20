package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.ifoodcorp.com.br/easyzap"
	"main.go/model"
)

type RegionRepositoryInterface interface {
	InsertRegion(ctx context.Context, region model.Region) error
	ListRegions(ctx context.Context) ([]model.Region, error)
	GetRegion(ctx context.Context, id string) (model.Region, error)
}

type RegionRepository struct {
	pgxpool *pgxpool.Pool
}

func NewRegionRepository(pgxpool *pgxpool.Pool) *MerchantRepository {
	return &MerchantRepository{pgxpool: pgxpool}
}

func (pg RegionRepository) InsertRegion(ctx context.Context, region model.Region) error {

	sql := `INSERT INTO region(id, name, status, latitude, longitude, cost,created_by,updated_by, created_at, updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`

	err := pg.pgxpool.QueryRow(ctx, sql, region.Id, region.Name, region.Status, region.Latitude,
		region.Longitude, region.Cost, region.CreatedBy, region.UpdatedBy, region.CreatedAt, region.UpdatedAt).Scan()
	if err != nil {
		easyzap.Error(context.Background(), err, "error on insert region")
		return err
	}

	return nil

}

func (pg RegionRepository) ListRegions(ctx context.Context) ([]model.Region, error) {

	sql := `SELECT id, name, status, latitude, longitude, cost,created_by,updated_by, created_at, updated_at FROM region`

	rows, err := pg.pgxpool.Query(ctx, sql)
	if err != nil {
		easyzap.Error(context.Background(), err, "error on list regions")
		return nil, err
	}

	defer rows.Close()
	var regions []model.Region
	for rows.Next() {
		var region model.Region
		err := rows.Scan(&region.Id, &region.Name, &region.Status, &region.Latitude,
			&region.Longitude, &region.Cost, &region.CreatedBy, &region.UpdatedBy, &region.CreatedAt, &region.UpdatedAt)
		if err != nil {
			easyzap.Error(context.Background(), err, "error on list regions")
			return nil, err
		}
		regions = append(regions, region)
	}

	return regions, nil
}

func (pg RegionRepository) GetRegion(ctx context.Context, id string) (model.Region, error) {
	sql := `SELECT id, name, status, latitude, longitude, cost,created_by,updated_by, created_at, updated_at FROM region where id = $1`

	var region model.Region
	err := pg.pgxpool.QueryRow(ctx, sql, id).Scan(&region.Id, &region.Name, &region.Status, &region.Latitude, &region.Longitude,
		&region.Cost, &region.CreatedBy, &region.UpdatedBy, &region.CreatedAt, &region.UpdatedAt)
	if err != nil {
		easyzap.Error(context.Background(), err, "error on get region")
		return region, err
	}
	return region, nil
}
