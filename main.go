package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/lehoon/core_dump_upload_server/v2/library/config"
	"github.com/lehoon/core_dump_upload_server/v2/library/logger"
	"github.com/lehoon/core_dump_upload_server/v2/routes"
	md "github.com/lehoon/core_dump_upload_server/v2/routes/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	route := chi.NewRouter()
	route.Use(middleware.RequestID)
	route.Use(middleware.Logger)
	route.Use(middleware.Recoverer)
	route.Use(middleware.URLFormat)
	//route.Use(middleware.Compress(6, "gzip"))
	route.Use(md.RequestLoggerFilter)
	//route.Use(render.SetContentType(render.ContentTypeJSON))

	route.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})

	route.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("错误,未找到请求路径"))
	})

	route.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("非法的method请求"))
	})

	route.Route("/api/v1", func(r chi.Router) {
		r.Mount("/", routes.Routes())
	})

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		fmt.Printf("%s %s\n", method, route)
		routes.PushRoute(method, route)
		return nil
	}

	if err := chi.Walk(route, walkFunc); err != nil {
		fmt.Printf("Logging error: %s\n", err.Error())
	}

	logger.Log().Info("dump文件服务器启动成功.")
	fmt.Printf("服务准备启动,本地监听地址:%s\n", config.GetLocalAddress())
	http.ListenAndServe(config.GetLocalAddress(), route)
}
