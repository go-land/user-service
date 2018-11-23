package main

import (
	"algorithms/config"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users/{name}", GetUserInfo)
	http.Handle("/", router)

	envConfig := config.NewEnvConfig()

	log.Printf("Using %s profile\n", envConfig.Profile)

	srv := &http.Server{
		Addr: "0.0.0.0:" + envConfig.Port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	log.Printf("user-service started at %s\n", envConfig.Port)

	go func() {
		srv.ListenAndServe()
	}()

	channelSignal := make(chan os.Signal)

	signal.Notify(channelSignal, os.Interrupt)

	signalValue := <-channelSignal

	log.Printf("Signal %v received\n", signalValue)

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
