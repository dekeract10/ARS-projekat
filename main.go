package main

import (
	"context"
	cs "github.com/dekeract10/ARS-projekat/configstore"
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

	store, err := cs.New()
	if err != nil {
		log.Fatal(err)
	}

	server := Service{
		store: store,
	}

	router.HandleFunc("/config/", server.createConfigHandler).Methods("POST")
	// router.HandleFunc("/config/{id}/", server.getConfigVersions).Methods("GET")
	// router.HandleFunc("/config/{id}/", server.createConfigHandler).Methods("PUT")
	router.HandleFunc("/config/{id}/{ver}", server.getConfigHandler).Methods("GET")
	// router.HandleFunc("/config/{id}/{ver}", server.getConfigHandler).Methods("DELETE")
	// router.HandleFunc("/config/{id}/", server.getAllConfigsHandler).Methods("GET")
	// router.HandleFunc("/config/{id}/{ver}/", server.getConfigHandler).Methods("GET")
	// router.HandleFunc("/config/{id}/{ver}/", server.delConfigHandler).Methods("DELETE")
	// router.HandleFunc("/group/{ver}/", server.createGroupHandler).Methods("POST")
	// router.HandleFunc("/group/{id}/", server.getAllGroupsHandler).Methods("GET")
	// router.HandleFunc("/group/{id}/{ver}/", server.getGroupHandler).Methods("GET")
	// router.HandleFunc("/group/{id}/{ver}/", server.delGroupHandler).Methods("DELETE")
	// router.HandleFunc("/group/{id}/configs/{ver}/", server.putConfigHandler).Methods("POST")

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
