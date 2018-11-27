package main

import (
	"context"
	"github.com/go-land/user-service/config"
	"github.com/go-land/user-service/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	envConfig := config.NewEnvConfig()

	log.Printf("Using %s profile\n", envConfig.Profile)

	srv := &http.Server{
		Addr: "0.0.0.0:" + envConfig.Port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 60,
		Handler:      createRouter(),
	}

	log.Printf("user-service started at %s\n", envConfig.Port)

	go func() {
		srv.ListenAndServe()
	}()

	waitForInterruptionSignal()

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
}

func createRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/users/{name}", handlers.GetUserInfo)
	http.Handle("/", router)
	return router
}

func waitForInterruptionSignal() {
	channelSignal := make(chan os.Signal)

	signal.Notify(channelSignal, os.Interrupt)

	signalValue := <-channelSignal

	log.Printf("Signal %v received\n", signalValue)
}
