package handlers

import (
	"github.com/Generalsimus/go-monolith-boilerplate/config"
	"github.com/Generalsimus/go-monolith-boilerplate/db/database"
	"github.com/Generalsimus/go-monolith-boilerplate/internal/handlers/user"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(db *database.Queries) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	humaConfig := huma.DefaultConfig("Go Boilerplate", "1.0.0")
	humaConfig.DocsRenderer = huma.DocsRendererScalar
	if config.Cfg.ENV == "prod" {
		humaConfig.DocsPath = ""
		humaConfig.OpenAPIPath = ""
	}
	api := humachi.New(r, humaConfig)

	huma.AutoRegister(huma.NewGroup(api, "/user"), &user.Handler{Db: db})

	// r.Get("/ws/", delivery.HandleWebSocket)

	return r
}
