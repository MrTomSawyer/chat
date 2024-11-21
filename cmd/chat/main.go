package main

import (
	"context"
	"flag"
	"os/signal"
	"syscall"
	"time"

	"github.com/MrTomSawyer/chat/internal/app"
	"github.com/MrTomSawyer/chat/internal/app/config"
	"github.com/gorilla/mux"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	var path string
	flag.StringVar(&path, "cfg", "./config/config.yml", "config file path")

	cfg := config.MustLoadConfig(path)
	r := mux.NewRouter()
	server := app.New(cfg.Server.Addr, r)

	go server.MustListenHTTP()

	<-ctx.Done()

	cwt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.ShutDownHTTP(cwt)
}
