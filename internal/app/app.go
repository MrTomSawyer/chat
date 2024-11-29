package app

import (
	"context"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
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

func (a *App) ListenHTTP() error {
	return a.httpServer.ListenAndServe()
}

func (a *App) StartWS() error {
	return nil
}

func (a *App) MustStart() {
	var g errgroup.Group
	g.Go(a.ListenHTTP)
	g.Go(a.StartWS)

	if err := g.Wait(); err != nil {
		panic(err)
	}
}

func (a *App) ShutDownHTTP(ctx context.Context) {
	if err := a.httpServer.Shutdown(ctx); err != nil {
		log.Println(err.Error())
	}
}
