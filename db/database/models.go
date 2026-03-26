package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Trade struct {
	ID         pgtype.UUID        `json:"id"`
	Pair       string             `json:"pair"`
	Price      pgtype.Numeric     `json:"price"`
	Amount     pgtype.Numeric     `json:"amount"`
	TradeType  string             `json:"trade_type"`
	ExecutedAt pgtype.Timestamptz `json:"executed_at"`
	Fee        pgtype.Numeric     `json:"fee"`
}
