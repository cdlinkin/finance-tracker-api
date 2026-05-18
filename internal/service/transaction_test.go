package service

import (
	"testing"

	"github.com/cdlinkin/finance-tracker-api/internal/domain"
	"github.com/cdlinkin/finance-tracker-api/internal/repository"
)

func TestCreate_Succes(t *testing.T) {
	mock := &repository.MockRepository{
		CreateFunc: func(tx domain.Transaction) (domain.Transaction, error) {
			tx.ID = 1
			return tx, nil
		},
	}

	svc := NewTransactionService(mock)

	res, err := svc.Create(domain.CreateTransactionRequest{
		Type:     domain.Income,
		Amount:   1000,
		Category: "salary",
	})
	if err != nil {
		t.Errorf("expected no error, get: %v", err)
	}
	if res.Amount != 1000 {
		t.Errorf("expected amount 1000, got: %v", err)
	}
}

func TestCreate_InvalidAmount(t *testing.T) {
	mock := &repository.MockRepository{
		CreateFunc: func(tx domain.Transaction) (domain.Transaction, error) {
			tx.ID = 1
			return tx, nil
		},
	}

	svc := NewTransactionService(mock)

	res, err := svc.Create(domain.CreateTransactionRequest{
		Type:     domain.Income,
		Amount:   -100,
		Category: "salary",
	})
	if err == nil {
		t.Errorf("expected error for negative amount, got nil: %v", err)
	}
	if res.Amount < 0 {
		t.Errorf("amount is bad: %v", err)
	}
}

func TestCreate_InvalidType(t *testing.T) {
	mock := &repository.MockRepository{
		CreateFunc: func(tx domain.Transaction) (domain.Transaction, error) {
			tx.ID = 1
			return tx, nil
		},
	}

	svc := NewTransactionService(mock)

	_, err := svc.Create(domain.CreateTransactionRequest{
		Type:     "unknown",
		Amount:   1000,
		Category: "salary",
	})
	if err == nil {
		t.Errorf("expected error for invalid type, got nil: %v", err)
	}
}
