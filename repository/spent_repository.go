package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.ifoodcorp.com.br/easyzap"
	"main.go/model"
)

type SpentRepositoryInterface interface {
	InsertSpent(ctx context.Context, spent model.Spent) error
	ListSpents(ctx context.Context) ([]model.Spent, error)
	GetSpent(ctx context.Context, id string) (model.Spent, error)
}

type SpentRepository struct {
	pgxpool *pgxpool.Pool
}

func NewSpentRepository(pgxpool *pgxpool.Pool) *SpentRepository {
	return &SpentRepository{pgxpool: pgxpool}
}

func (pg SpentRepository) InsertSpent(ctx context.Context, spent model.Spent) error {
	sql := `INSERT INTO merchant(id, campaing_id,bucket, total_spent, total_clicks, total_impressions) VALUES($1,$2,$3,$4,$5,$6)`

	err := pg.pgxpool.QueryRow(ctx, sql, spent.Id, spent.CampaingId, spent.Bucket, spent.TotalSpent, spent.TotalClicks, spent.TotalImpressions).Scan()
	if err != nil {
		easyzap.Error(context.Background(), err, "error on insert spent")
		return err
	}
	return nil
}

func (pg SpentRepository) ListSpents(ctx context.Context) ([]model.Spent, error) {

	sql := `SELECT id, campaing_id,bucket, total_spent, total_clicks, total_impressions FROM spent`

	rows, err := pg.pgxpool.Query(ctx, sql)
	if err != nil {
		easyzap.Error(context.Background(), err, "error on list spents")
		return nil, err
	}

	defer rows.Close()
	var spents []model.Spent
	for rows.Next() {
		var spent model.Spent
		err := rows.Scan(&spent.Id, &spent.CampaingId, &spent.Bucket, &spent.TotalSpent, &spent.TotalClicks, &spent.TotalImpressions)
		if err != nil {
			easyzap.Error(context.Background(), err, "error on list spents")
			return nil, err
		}
		spents = append(spents, spent)
	}

	return spents, nil
}

func (pg SpentRepository) GetSpent(ctx context.Context, id string) (model.Spent, error) {

	sql := `SELECT id, campaing_id,bucket, total_spent, total_clicks, total_impressions FROM spent where id = $1`

	var spent model.Spent
	err := pg.pgxpool.QueryRow(ctx, sql, id).Scan(&spent.Id, &spent.CampaingId, &spent.Bucket, &spent.TotalSpent, &spent.TotalClicks, &spent.TotalImpressions)
	if err != nil {
		easyzap.Error(context.Background(), err, "error on get merchant")
		return spent, err
	}
	return spent, nil
}
