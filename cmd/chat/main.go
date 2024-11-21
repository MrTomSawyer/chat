package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/MrTomSawyer/chat/internal/app"
	"github.com/MrTomSawyer/chat/internal/app/config"
	"github.com/MrTomSawyer/chat/internal/app/domain/user"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	var path string
	flag.StringVar(&path, "cfg", "./config/config.yml", "config file path")
	cfg := config.MustLoadConfig(path)

	conn, err := pgx.Connect(ctx, cfg.DB.DSN)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := conn.Close(ctx); err != nil {
			log.Printf("failed to close db connection: %s\n", err.Error())
		}
	}()

	r := mux.NewRouter()

	userRepo := user.NewUserRepository(conn, "user")
	userServ := user.NewUserService(userRepo)
	user.NewUserController(r, userServ).Init()

	server := app.New(cfg.Server.Addr, r)
	go server.MustListenHTTP()

	fmt.Println("=====================")
	fmt.Println("Chat server v1.0")
	fmt.Println("Status: Running")
	fmt.Printf("Environment: %s\n", cfg.Env)
	fmt.Printf("Port: %s\n", cfg.Server.Addr)
	fmt.Println("=====================")

	<-ctx.Done()

	cwt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.ShutDownHTTP(cwt)
}
