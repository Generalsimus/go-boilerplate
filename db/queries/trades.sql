-- name: CreateTrade :one
INSERT INTO trades (pair, price, amount, trade_type)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetTradeByID :one
SELECT * FROM trades
WHERE id = $1 LIMIT 1;

-- name: ListRecentTrades :many
SELECT * FROM trades
WHERE pair = $1
ORDER BY executed_at DESC
LIMIT $2;