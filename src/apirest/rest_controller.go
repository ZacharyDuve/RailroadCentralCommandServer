package apirest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("all systems online"))
}

func ServeRESTAPIs() http.Handler {
	router := mux.NewRouter()

	router.Methods(http.MethodGet).HandlerFunc(healthHandler)

	return router
}
