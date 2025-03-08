package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"dictionary-app/db"
	"dictionary-app/graphql"

	"github.com/graphql-go/handler"
)

func main() {
	db.InitDB()

	schema, err := graphql.CreateSchema()
	if err != nil {
		log.Fatalf("Error creating schema: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema: schema,
		Pretty: true,
	})

	http.Handle("/dictionary", h)
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("GraphQL server running on http://localhost:" + port + "/dictionary")
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		if opErr, ok := err.(*net.OpError); ok && opErr.Op == "listen" {
			log.Fatalf("Port %s is already in use", port)
		} else {
			log.Fatalf("Error starting server: %v", err)
		}
	}
}
