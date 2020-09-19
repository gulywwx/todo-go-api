package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gulywwx/todo-go-api/conf"
	"github.com/gulywwx/todo-go-api/routers"
)

func main() {
	cfg, err := conf.Load()
	if err != nil {
		log.Fatal(err)
	}

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:        router,
		ReadTimeout:    cfg.Server.ReadTimeout * time.Second,
		WriteTimeout:   cfg.Server.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	errChan := make(chan error)

	go func() {
		errChan <- s.ListenAndServe()
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	error := <-errChan
	log.Fatal(error)

}
