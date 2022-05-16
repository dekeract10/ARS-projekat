package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	router := mux.NewRouter()
	router.StrictSlash(true)

	server := Service{
		versions: make(map[string]map[string][]*Config),
	}

	router.HandleFunc("/{ver}/config/", server.createConfigHandler).Methods("POST")
	router.HandleFunc("/{ver}/group/", server.createGroupHandler).Methods("POST")
	router.HandleFunc("/{ver}/configs/", server.getAllConfigsHandler).Methods("GET")
	router.HandleFunc("/{ver}/groups/", server.getAllGroupsHandler).Methods("GET")
	router.HandleFunc("/{ver}/group/{id}", server.getGroupHandler).Methods("GET")
	router.HandleFunc("/{ver}/config/{id}", server.getConfigHandler).Methods("GET")
	router.HandleFunc("/{ver}/config/{id}/", server.delConfigHandler).Methods("DELETE")
	router.HandleFunc("/{ver}/group/{id}/", server.delGroupHandler).Methods("DELETE")
	router.HandleFunc("/{ver}/group/{id}/configs", server.putConfigHandler).Methods("POST")

	// start server
	srv := &http.Server{Addr: "0.0.0.0:8000", Handler: router}
	go func() {
		log.Println("server starting")
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-quit

	log.Println("service shutting down ...")

	// gracefully stop server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("server stopped")
}
