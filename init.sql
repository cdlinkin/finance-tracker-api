CREATE TABLE IF NOT EXISTS transactions (
    id         SERIAL PRIMARY KEY,
    type       VARCHAR(10) NOT NULL CHECK (type IN ('income', 'expense')),
    amount     NUMERIC(12, 2) NOT NULL CHECK (amount > 0),
    category   VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);