package main

import (
	"log"
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
	http.ListenAndServe(":"+port, nil)
}
