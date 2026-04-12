package user

import (
	"github.com/Generalsimus/go-monolith-boilerplate/db/database"
)

type Handler struct {
	Db *database.Queries
}
