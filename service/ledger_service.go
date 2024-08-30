package service

import (
	"context"

	"main.go/model"
	"main.go/repository"
)

type LedgerServiceInterface interface {
	AddNewLedger(ctx context.Context, ledger model.Ledger) error
}

type LedgerService struct {
	ledgerRepository repository.LedgerRepositoryInterface
}

func NewLedgerService(ledgerRepository repository.LedgerRepositoryInterface) LedgerService {
	return LedgerService{
		ledgerRepository: ledgerRepository,
	}
}

func (l LedgerService) AddNewLedger(ctx context.Context, ledger model.Ledger) error {

	err := l.ledgerRepository.InsertLedger(ctx, model.Ledger{
		Id:         ledger.Id,
		CampaingId: ledger.CampaingId,
		SpentId:    ledger.SpentId,
		SlugId:     ledger.SpentId,
		UserId:     ledger.UserId,
		EventType:  ledger.EventType,
		Latitude:   ledger.Latitude,
		Longitude:  ledger.Longitude,
		Cost:       ledger.Cost,
		CreatedAt:  ledger.CreatedAt,
	})
	if err != nil {
		return err
	}
	return nil
}
