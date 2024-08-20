package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.ifoodcorp.com.br/easyzap"
	"main.go/model"
)

type LedgerRepositoryInterface interface {
	InsertLedger(ctx context.Context, ledger model.Ledger) error
	ListLedgers(ctx context.Context) ([]model.Ledger, error)
	GetLedger(ctx context.Context, id string) (model.Ledger, error)
}

type LedgerRepository struct {
	pgxpool *pgxpool.Pool
}

func NewLedgerRepository(pgxpool *pgxpool.Pool) *LedgerRepository {
	return &LedgerRepository{pgxpool: pgxpool}
}

func (pg SpentRepository) InsertLedger(ctx context.Context, ledger model.Ledger) error {
	sql := `INSERT INTO ledger(id, campaing_id,spent_id,slug_id,user_id,event_type,latitude,longitude,cost,created_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`

	err := pg.pgxpool.QueryRow(ctx, sql, ledger.Id, ledger.CampaingId, ledger.SpentId, ledger.SlugId, ledger.UserId,
		ledger.EventType, ledger.Latitude, ledger.Longitude, ledger.Cost, ledger.CreatedAt).Scan()
	if err != nil {
		easyzap.Error(context.Background(), err, "error on insert ledger")
		return err
	}
	return nil
}

func (pg LedgerRepository) ListLedgers(ctx context.Context) ([]model.Ledger, error) {

	sql := `SELECT id, campaing_id,spent_id,slug_id,user_id,event_type,latitude,longitude,cost,created_at FROM ledger`

	rows, err := pg.pgxpool.Query(ctx, sql)
	if err != nil {
		easyzap.Error(context.Background(), err, "error on list ledgers")
		return nil, err
	}

	defer rows.Close()
	var ledgers []model.Ledger
	for rows.Next() {
		var ledger model.Ledger
		err := rows.Scan(&ledger.Id, &ledger.CampaingId, &ledger.SpentId, &ledger.SlugId, &ledger.UserId,
			&ledger.EventType, &ledger.Latitude, &ledger.Longitude, &ledger.Cost, &ledger.CreatedAt)
		if err != nil {
			easyzap.Error(context.Background(), err, "error on list ledgers")
			return nil, err
		}
		ledgers = append(ledgers, ledger)
	}

	return ledgers, nil

}

func (pg LedgerRepository) GetLedger(ctx context.Context, id string) (model.Ledger, error) {

	sql := `SELECT id, campaing_id,spent_id,slug_id,user_id,event_type,latitude,longitude,cost,created_at FROM ledger where id = $1`

	var ledger model.Ledger
	err := pg.pgxpool.QueryRow(ctx, sql, id).Scan(ledger.Id, &ledger.CampaingId, &ledger.SpentId, &ledger.SlugId, &ledger.UserId,
		&ledger.EventType, &ledger.Latitude, &ledger.Longitude, &ledger.Cost, &ledger.CreatedAt)
	if err != nil {
		easyzap.Error(context.Background(), err, "error on get merchant")
		return ledger, err
	}
	return ledger, nil

}
