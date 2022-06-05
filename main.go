package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	router := mux.NewRouter()
	router.StrictSlash(true)

	server, err := NewConfigServer()
	if err != nil {
		log.Fatal(err)
		return
	}

	router.HandleFunc("/config/", countPostConfig(server.createConfigHandler)).Methods("POST")
	router.HandleFunc("/config/{id}/", countGetConfigVer(server.getConfigVersionsHandler)).Methods("GET")
	router.HandleFunc("/config/{id}", countPostConfigVer(server.putNewVersion)).Methods("POST")
	router.HandleFunc("/config/{id}/{ver}", countGetConfig(server.getConfigHandler)).Methods("GET")
	router.HandleFunc("/config/{id}/{ver}", countDelConfig(server.delConfigHandler)).Methods("DELETE")
	// router.HandleFunc("/config/{id}/{ver}", server.getConfigHandler).Methods("DELETE")
	// router.HandleFunc("/config/{id}/", server.getAllConfigsHandler).Methods("GET")
	// router.HandleFunc("/config/{id}/{ver}/", server.getConfigHandler).Methods("GET")
	// router.HandleFunc("/config/{id}/{ver}/", server.delConfigHandler).Methods("DELETE")
	router.HandleFunc("/group/", countPostGroup(server.createGroupHandler)).Methods("POST")
	router.HandleFunc("/group/{id}", countPostGroupVer(server.putNewGroupVersion)).Methods("POST")
	// router.HandleFunc("/group/{id}/", server.getGroupVersionsHandler).Methods("GET")
	router.HandleFunc("/group/{id}/{ver}/", countGetGroup(server.getGroupHandler)).Methods("GET")
	// router.HandleFunc("/group/{id}/{ver}/", server.getLabelsHandler).Methods("GET")
	router.HandleFunc("/group/{id}/{ver}/", countDelGroup(server.delGroupHandler)).Methods("DELETE")
	router.HandleFunc("/group/{id}/{ver}/config/", countGetGroupConfigs(server.getConfigFromGroup)).Methods("GET")
	router.HandleFunc("/group/{id}/{ver}/config/", countAddGroupConfig(server.addConfigToGroupHandler)).Methods("POST")
	router.Path("/metrics").Handler(metricsHandler())
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
