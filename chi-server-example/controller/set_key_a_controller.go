package controller

import (
	"example/chi_router"
	"example/usecase"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/go-chi/chi/v5"
)

type setKeyAcontroller struct {
	router     *chi.Mux
	updateAKey usecase.IUpdateAKey
}

var setKeyAcontrollerDependency = ioc.InjectInboundAdapter[setKeyAcontroller](func() (setKeyAcontroller, error) {
	controller := setKeyAcontroller{
		router:     chi_router.Mux.Dependency,
		updateAKey: usecase.UpdateAKey.Dependency,
	}
	controller.router.Get("/a/{aParam}", controller.handle)
	return controller, nil
})

func (ctrl setKeyAcontroller) handle(w http.ResponseWriter, r *http.Request) {
	ctrl.updateAKey(chi.URLParam(r, "aParam"))
	w.Write([]byte("a key updated"))
}
