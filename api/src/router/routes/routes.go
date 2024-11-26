package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Rotas da API
type Rota struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	NeedAuth bool
}

func Configure(r *mux.Router) *mux.Router {
	for _,route := range userRoutes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
	return r
}
