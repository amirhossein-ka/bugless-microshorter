package http

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"ush/internal/config"
	"ush/internal/gateway/controller"
	"ush/internal/gateway/service"
)

// Server is used to implement Controller interface
type handler struct {
	srv service.Service
}

type rest struct {
	handler *handler
	server  *http.Server
	router  *mux.Router
}

func (r *rest) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r.router.ServeHTTP(w, req)
}

func (r *rest) Start(s string) error {
	r.routing()
	r.server = &http.Server{
		Addr:         s,
		Handler:      r,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	return r.server.ListenAndServe()
}

func (r *rest) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return r.server.Shutdown(ctx)
}

func New(cfg *config.GatewayConfig, s service.Service) controller.Controller {
	return &rest{
		handler: &handler{srv: s},
		router:  mux.NewRouter(),
	}
}
