package domain

import "time"

type Type string

const (
	Income  Type = "income"
	Expense Type = "expense"
)

type Transaction struct {
	ID        int       `json:"id"`
	Type      Type      `json:"type"`
	Amount    float64   `json:"amount"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateTransactionRequest struct {
	Type     Type    `json:"type"`
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
}
