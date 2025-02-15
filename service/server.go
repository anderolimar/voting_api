package service

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"votingapi/bootstrap"
	"votingapi/service/handlers"
)

func Start() {

	bootstrap.Start()

	router := http.NewServeMux()
	handlers.RegisterRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)

	log.Printf("Server started on %s\n", addr)

	srv := http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Fatal(srv.ListenAndServe())
}
