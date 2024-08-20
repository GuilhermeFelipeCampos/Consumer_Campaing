package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.ifoodcorp.com.br/easyzap"
	"main.go/model"
)

type CampaingRepositoryInterface interface {
	InsertCampaing(ctx context.Context, campaing model.Campaing) error
	ListCampaings(ctx context.Context) ([]model.Campaing, error)
	GetCampaing(ctx context.Context, id string) (model.Campaing, error)
}

type CampaingRepository struct {
	pgxpool *pgxpool.Pool
}

func NewCampaingRepository(pgxpool *pgxpool.Pool) *CampaingRepository {
	return &CampaingRepository{pgxpool: pgxpool}
}

func (pg CampaingRepository) InsertCampaing(ctx context.Context, campaing model.Campaing) error {

	sql := `INSERT INTO campaing(id, merchant_id, status, latitude, longitude, budget,created_by,updated_by, created_at, updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`

	err := pg.pgxpool.QueryRow(ctx, sql, campaing.Id, campaing.MerchantId, campaing.Status, campaing.Latitude,
		campaing.Longitude, campaing.Budget, campaing.CreatedBy, campaing.UpdatedBy, campaing.CreatedAt, campaing.UpdatedAt).Scan()

	if err != nil {
		easyzap.Error(context.Background(), err, "error on insert campaing")
		return err
	}
	return nil
}

func (pg CampaingRepository) ListCampaings(ctx context.Context) ([]model.Campaing, error) {
	sql := `SELECT id, merchant_id, status, latitude, longitude, budget,created_by,updated_by, created_at, updated_at FROM campaing`
	rows, err := pg.pgxpool.Query(ctx, sql)
	if err != nil {
		easyzap.Error(context.Background(), err, "error on list campaings")
		return nil, err
	}
	defer rows.Close()
	var campaings []model.Campaing
	for rows.Next() {
		var campaing model.Campaing
		err := rows.Scan(&campaing.Id, &campaing.MerchantId, &campaing.Status, &campaing.Latitude, &campaing.Longitude,
			&campaing.Budget, &campaing.CreatedBy, &campaing.UpdatedBy, &campaing.CreatedAt, &campaing.UpdatedAt)
		if err != nil {
			easyzap.Error(context.Background(), err, "error on list campaings")
			return nil, err
		}
		campaings = append(campaings, campaing)
	}
	return campaings, nil

}

func (pg CampaingRepository) GetCampaing(ctx context.Context, id string) (model.Campaing, error) {
	sql := `SELECT id, merchant_id, status, latitude, longitude, budget, created_by, updated_by, created_at, updated_at FROM campaing WHERE id = $1`
	var campaing model.Campaing
	err := pg.pgxpool.QueryRow(ctx, sql, id).Scan(&campaing.Id, &campaing.MerchantId, &campaing.Status, &campaing.Latitude, &campaing.Longitude,
		&campaing.Budget, &campaing.CreatedBy, &campaing.UpdatedBy, &campaing.CreatedAt, &campaing.UpdatedAt)
	if err != nil {
		easyzap.Error(context.Background(), err, "error on get campaing")
		return campaing, err
	}
	return campaing, nil
}
