package controller

import (
	"example/chi_router"
	"example/usecase"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/go-chi/chi/v5"
)

type setKeyBcontroller struct {
	router     *chi.Mux
	updateBKey usecase.IUpdateBKey
}

var setKeyBcontrollerDependency = ioc.InboundAdapter[setKeyBcontroller](func() (setKeyBcontroller, error) {
	controller := setKeyBcontroller{
		router:     chi_router.Mux.Dependency,
		updateBKey: usecase.UpdateBKey.Dependency,
	}
	controller.router.Get("/b/{bParam}", controller.handle)
	return controller, nil
})

func (ctrl setKeyBcontroller) handle(w http.ResponseWriter, r *http.Request) {
	ctrl.updateBKey(chi.URLParam(r, "bParam"))
	w.Write([]byte("b key updated"))
}
