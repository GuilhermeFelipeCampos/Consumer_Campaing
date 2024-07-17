package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.ifoodcorp.com.br/easyzap"
	"main.go/model"
)

type OwnerRepository interface {
	InsertOwner(ctx context.Context, owner model.Owner) error
	ListOwner(ctx context.Context) ([]model.Owner, error)
	GetOwner(ctx context.Context, id string) (model.Owner, error)
}

type OnwerRepository struct {
	pgxpool *pgxpool.Pool
}

func NewOwnerRepository(pgxpool *pgxpool.Pool) *OnwerRepository {
	return &OnwerRepository{pgxpool: pgxpool}
}

func (pg OnwerRepository) InsertOwner(ctx context.Context, owner model.Owner) error {

	sql := `INSERT INTO owner(id, email, status, created_by, updated_by, created_at, updated_at) VALUES($1,$2,$3,$4,$5,$6,$7)`

	err := pg.pgxpool.QueryRow(ctx, sql, owner.Id, owner.Email, owner.Status, owner.CreatedBy, owner.UpdatedBy, owner.CreatedAt, owner.UpdatedAt).Scan()

	if err != nil {
		easyzap.Error(context.Background(), err, "error on insert owner")
		return err
	}
	return nil
}

func (pg OnwerRepository) ListOwners(ctx context.Context) ([]model.Owner, error) {
	sql := `SELECT id, name, status, cost, created_by, updated_by, created_at, updated_at FROM owner`
	rows, err := pg.pgxpool.Query(ctx, sql)
	if err != nil {
		easyzap.Error(context.Background(), err, "error on list owners")
		return nil, err
	}
	defer rows.Close()
	var owners []model.Owner
	for rows.Next() {
		var owner model.Owner
		err := rows.Scan(&owner.Id, &owner.Email, &owner.Status, &owner.CreatedBy, &owner.UpdatedBy, &owner.CreatedAt, &owner.UpdatedAt)
		if err != nil {
			easyzap.Error(context.Background(), err, "error on list owners")
			return nil, err
		}
		owners = append(owners, owner)
	}
	return owners, nil

}

func (pg OnwerRepository) GetOwner(ctx context.Context, id string) (model.Owner, error) {
	sql := `SELECT id, email, status, created_by, updated_by, created_at, updated_at FROM owner WHERE id = $1`
	var owner model.Owner
	err := pg.pgxpool.QueryRow(ctx, sql, id).Scan(&owner.Id, &owner.Email, &owner.Status, &owner.CreatedBy, &owner.UpdatedBy, &owner.CreatedAt, &owner.UpdatedAt)
	if err != nil {
		easyzap.Error(context.Background(), err, "error on get owner")
		return owner, err
	}
	return owner, nil
}
