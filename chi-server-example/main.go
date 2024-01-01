package main

import (
	"example/chi_router"
	_ "example/controller"
	_ "example/kvdatabase"
	_ "example/usecase"
	"fmt"
	"net/http"
	"os"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

func main() {
	if err := ioc.LoadDependencies(); err != nil {
		os.Exit(0)
	}
	fmt.Println("server starting at port 3000")
	http.ListenAndServe(":3000", chi_router.Mux.Dependency)
}
