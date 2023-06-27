package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/lehoon/core_dump_upload_server/v2/api"
)

type RouteInfo struct {
	Method string
	Path   string
}

var routes = []RouteInfo{}

func PushRoute(method, path string) {
	routeInfo := RouteInfo{
		Method: method,
		Path:   path,
	}

	routes = append(routes, routeInfo)
}

func GetRoutes(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, api.SuccessBizResultWithData(routes))
}

func Routes() http.Handler {
	r := chi.NewRouter()

	r.Route("/dump", func(r chi.Router) {
		r.Post("/upload*", api.UploadDumpFile)
	})

	r.Route("/", func(r chi.Router) {
		r.Get("/", GetRoutes)
	})

	return r
}
