package repository

import (
	"fmt"

	"github.com/cdlinkin/finance-tracker-api/internal/domain"
	"github.com/jmoiron/sqlx"
)

type PostgresRepo struct {
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (p *PostgresRepo) Create(tx domain.Transaction) (domain.Transaction, error) {
	query := `INSERT INTO transactions (type, amount, category, created_at)
			VALUES ($1, $2, $3, $4) RETURNING id, type, amount, category, created_at`
	err := p.db.QueryRowx(query, tx.Type, tx.Amount, tx.Category, tx.CreatedAt).StructScan(&tx)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("failed to create transaction in database")
	}
	return tx, nil
}

func (p *PostgresRepo) GetAll() ([]domain.Transaction, error) {
	query := `SELECT id, type, amount, category, created_at FROM transactions`

	var trs []domain.Transaction
	err := p.db.Select(&trs, query)
	if err != nil {
		return []domain.Transaction{}, fmt.Errorf("failed to get all transaction in database")
	}
	return trs, nil
}

func (p *PostgresRepo) Delete(id int) error {
	query := `DELETE FROM transactions WHERE id = $1`
	result, err := p.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete transaction in database")
	}
	check, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to rows affected in database")
	}
	if check == 0 {
		return fmt.Errorf("transaction not found")
	}
	return nil
}

func (p *PostgresRepo) Summary() (float64, float64, error) {
	query := `SELECT 
    			COALESCE(SUM(amount) FILTER (WHERE type = 'income'), 0) as income,
    			COALESCE(SUM(amount) FILTER (WHERE type = 'expense'), 0) as expense
			FROM transactions`
	var result struct {
		Income  float64 `db:"income"`
		Expense float64 `db:"expense"`
	}
	err := p.db.QueryRowx(query).StructScan(&result)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get summary: %w", err)
	}
	return result.Income, result.Expense, nil
}
