package domain

import "time"

type Type string

const (
	Income  Type = "income"
	Expense Type = "expense"
)

type Transaction struct {
	ID        int       `json:"id" db:"id"`
	Type      Type      `json:"type" db:"type"`
	Amount    float64   `json:"amount" db:"amount"`
	Category  string    `json:"category" db:"category"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CreateTransactionRequest struct {
	Type     Type    `json:"type"`
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
}
