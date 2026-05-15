package service

import (
	"fmt"
	"time"

	"github.com/cdlinkin/finance-tracker-api/internal/domain"
	"github.com/cdlinkin/finance-tracker-api/internal/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepo
}

func NewTransactionService(repo *repository.TransactionRepo) *TransactionService {
	return &TransactionService{
		repo: repo,
	}
}

func (s *TransactionService) Create(r domain.CreateTransactionRequest) (domain.Transaction, error) {
	if r.Amount <= 0 {
		return domain.Transaction{}, fmt.Errorf("amount must be positive")
	}

	if r.Type != domain.Income && r.Type != domain.Expense {
		return domain.Transaction{}, fmt.Errorf("invalid type")
	}

	tx := domain.Transaction{
		Type:      r.Type,
		Amount:    r.Amount,
		Category:  r.Category,
		CreatedAt: time.Now(),
	}

	return s.repo.Create(tx)
}

func (s *TransactionService) GetAll() ([]domain.Transaction, error) {
	return s.repo.GetAll()
}

func (s *TransactionService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *TransactionService) Summary() (income, expense, balance float64, err error) {
	income, expense, err = s.repo.Summary()
	return income, expense, income - expense, err
}
