package main

import (
	"gqlgen-todos/bootstrap"

	"go.uber.org/fx"
)

func main() {
	fx.New(bootstrap.Module).Run()
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = defaultPort
	// }

	// srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))
}
