package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alokMIPL/students-api/internal/config"
)

func main() {

	fmt.Println("Welcome to students api")

	// Load config
	cfg := config.MustLoad()
	log.Println("Environment:", cfg.Env)

	// database setup

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students api"))
	})

	// setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

}
