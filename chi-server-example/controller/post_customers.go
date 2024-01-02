package controller

import (
	"example/chi_router"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/go-chi/chi/v5"
)

type postCustomers struct {
	router *chi.Mux
}

var postCustomersDependency = ioc.InboundAdapter[postCustomers](func() (postCustomers, error) {
	controller := postCustomers{
		router: chi_router.Mux.Dependency,
	}
	controller.router.Post("/person", controller.handle)
	return controller, nil
})

func (ctrl postCustomers) handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from post customer controller"))
}
