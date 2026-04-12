package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTrade = `-- name: CreateTrade :one
INSERT INTO trades (pair, price, amount, trade_type)
VALUES ($1, $2, $3, $4)
RETURNING id, pair, price, amount, trade_type, executed_at, fee
`

type CreateTradeParams struct {
	Pair      string         `json:"pair"`
	Price     pgtype.Numeric `json:"price"`
	Amount    pgtype.Numeric `json:"amount"`
	TradeType string         `json:"trade_type"`
}

func (q *Queries) CreateTrade(ctx context.Context, arg CreateTradeParams) (Trade, error) {
	row := q.db.QueryRow(ctx, createTrade,
		arg.Pair,
		arg.Price,
		arg.Amount,
		arg.TradeType,
	)
	var i Trade
	err := row.Scan(
		&i.ID,
		&i.Pair,
		&i.Price,
		&i.Amount,
		&i.TradeType,
		&i.ExecutedAt,
		&i.Fee,
	)
	return i, err
}

const getTradeByID = `-- name: GetTradeByID :one
SELECT id, pair, price, amount, trade_type, executed_at, fee FROM trades
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTradeByID(ctx context.Context, id pgtype.UUID) (Trade, error) {
	row := q.db.QueryRow(ctx, getTradeByID, id)
	var i Trade
	err := row.Scan(
		&i.ID,
		&i.Pair,
		&i.Price,
		&i.Amount,
		&i.TradeType,
		&i.ExecutedAt,
		&i.Fee,
	)
	return i, err
}

const listRecentTrades = `-- name: ListRecentTrades :many
SELECT id, pair, price, amount, trade_type, executed_at, fee FROM trades
WHERE pair = $1
ORDER BY executed_at DESC
LIMIT $2
`

type ListRecentTradesParams struct {
	Pair  string `json:"pair"`
	Limit int32  `json:"limit"`
}

func (q *Queries) ListRecentTrades(ctx context.Context, arg ListRecentTradesParams) ([]Trade, error) {
	rows, err := q.db.Query(ctx, listRecentTrades, arg.Pair, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Trade
	for rows.Next() {
		var i Trade
		if err := rows.Scan(
			&i.ID,
			&i.Pair,
			&i.Price,
			&i.Amount,
			&i.TradeType,
			&i.ExecutedAt,
			&i.Fee,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
