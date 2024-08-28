package handler

import (
	_ "SpellNote/docs"
	"SpellNote/pkg/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json"),
	))
	router.Route("/api", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", h.loginUser)
			r.Post("/register", h.registerUser)
		})

		r.Route("/note", func(r chi.Router) {
			r.Use(h.userIdentity)
			r.Get("/all", h.getAllNotes)

			r.Post("/", h.createNote)
			r.Get("/", h.getUserNotes)
		})
	})

	return router
}
