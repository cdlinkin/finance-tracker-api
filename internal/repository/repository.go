package repository

import "github.com/cdlinkin/finance-tracker-api/internal/domain"

type Repository interface {
	Create(tx domain.Transaction) (domain.Transaction, error)
	GetAll() ([]domain.Transaction, error)
	Delete(id int) error
	Summary() (float64, float64, error)
}
