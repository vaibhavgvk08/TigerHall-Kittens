package main

import (
	"github.com/vaibhavgvk08/tigerhall-kittens/services/auth"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vaibhavgvk08/tigerhall-kittens/graph"
)

const defaultIPAddress = "127.0.0.1"
const defaultPort = "8080"

func main() {
	ip := os.Getenv("IP")
	if ip == "" {
		ip = defaultIPAddress
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", auth.AuthMiddleware(srv))
	// todo : Add user validation middleware as well later.

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(ip+":"+port, nil))
}
