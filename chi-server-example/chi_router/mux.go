package chi_router

import (
	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

var Mux = ioc.Installation[*chi.Mux](func() (*chi.Mux, error) {
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	return r, nil
})
