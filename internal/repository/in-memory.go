package repository

import (
	"fmt"

	"github.com/cdlinkin/finance-tracker-api/internal/domain"
)

type InMemoryRepo struct {
	transaction map[int]domain.Transaction
	nextID      int
}

func NewTransactionRepo() *InMemoryRepo {
	return &InMemoryRepo{
		transaction: map[int]domain.Transaction{},
		nextID:      1,
	}
}

func (t *InMemoryRepo) Create(tx domain.Transaction) (domain.Transaction, error) {
	tx.ID = t.nextID
	t.nextID++
	t.transaction[tx.ID] = tx
	return tx, nil
}

func (t *InMemoryRepo) GetAll() ([]domain.Transaction, error) {
	transactions := make([]domain.Transaction, 0)
	for _, tr := range t.transaction {
		transactions = append(transactions, tr)
	}
	return transactions, nil
}

func (t *InMemoryRepo) Delete(id int) error {
	if _, exists := t.transaction[id]; !exists {
		return fmt.Errorf("transaction not found")
	}
	delete(t.transaction, id)
	return nil
}

func (t *InMemoryRepo) Summary() (income, expense float64, err error) {
	for _, tr := range t.transaction {
		if tr.Type == domain.Income {
			income += tr.Amount
		} else {
			expense += tr.Amount
		}
	}
	return income, expense, nil
}
