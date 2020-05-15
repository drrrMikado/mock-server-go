package main

import (
	"context"
	"log"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/drrrMikado/mock-server-go/conf"
	"github.com/drrrMikado/mock-server-go/servers/http"
	"github.com/drrrMikado/mock-server-go/services"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}
	s := services.New(conf.Conf)
	http.Init(s)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	q := <-quit
	log.Println("Signal: ", q)
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := http.Close(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
