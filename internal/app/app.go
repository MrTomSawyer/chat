package app

import (
	"context"
	"errors"
	"log"
	"net/http"
)

type App struct {
	httpServer http.Server
}

func New(addr string, handler http.Handler) *App {
	return &App{
		httpServer: http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

func (a *App) MustListenHTTP() {
	if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

func (a *App) ShutDownHTTP(ctx context.Context) {
	if err := a.httpServer.Shutdown(ctx); err != nil {
		log.Println(err.Error())
	}
}
