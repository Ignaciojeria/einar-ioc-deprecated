package controller

import (
	"example/chi_router"
	"example/usecase"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/go-chi/chi/v5"
)

type getCurrentKeyController struct {
	router        *chi.Mux
	getCurrentKey usecase.IgetCurrentKey
}

var getCurrentKeyControllerDependency = ioc.InjectInboundAdapter[getCurrentKeyController](func() (getCurrentKeyController, error) {
	controller := getCurrentKeyController{
		router:        chi_router.Mux.Dependency,
		getCurrentKey: usecase.GetCurrentKey.Dependency,
	}
	controller.router.Get("/getKeys", controller.handle)
	return controller, nil
})

func (ctrl getCurrentKeyController) handle(w http.ResponseWriter, r *http.Request) {
	w.Write(
		[]byte("current key value for a : " + ctrl.getCurrentKey("a") +
			"\n" + "current key value for b : " + ctrl.getCurrentKey("b")))
}
