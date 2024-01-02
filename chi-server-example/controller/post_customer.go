package controller

import (
	"example/chi_router"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/go-chi/chi/v5"
)

type postCustomer struct {
	router *chi.Mux
}

var postCustomerDependency = ioc.InboundAdapter[postCustomer](func() (postCustomer, error) {
	controller := postCustomer{
		router: chi_router.Mux.Dependency, // Installation Dependency Injected Here!
	}
	controller.router.Post("/customers", controller.handle)
	return controller, nil
})

func (ctrl postCustomer) handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from post customer controller"))
}
