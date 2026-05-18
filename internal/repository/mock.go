package repository

import "github.com/cdlinkin/finance-tracker-api/internal/domain"

type MockRepository struct {
	CreateFunc  func(tx domain.Transaction) (domain.Transaction, error)
	GetAllFunc  func() ([]domain.Transaction, error)
	DeleteFunc  func(id int) error
	SummaryFunc func() (float64, float64, error)
}

func (m *MockRepository) Create(tx domain.Transaction) (domain.Transaction, error) {
	return m.CreateFunc(tx)
}

func (m *MockRepository) GetAll() ([]domain.Transaction, error) {
	return m.GetAllFunc()
}

func (m *MockRepository) Delete(id int) error {
	return m.DeleteFunc(id)
}

func (m *MockRepository) Summary() (float64, float64, error) {
	return m.SummaryFunc()
}
