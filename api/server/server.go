package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/FernandoH-G/gw2-items-server/graph"
	"github.com/FernandoH-G/gw2-items-server/graph/generated"
)

const DEFAULT_PORT = "80"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Print("connect to http://localhost:5001/ or http://{digitalOceanIP}:5001 for GraphQL playground")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
