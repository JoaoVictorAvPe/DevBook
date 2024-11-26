package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

//Retorna um router com as rotas configuradas
func Create() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
