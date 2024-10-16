package bootstrap

import (
	"context"
	"fmt"
	"gqlgen-todos/api/repository"
	"gqlgen-todos/api/services"
	"gqlgen-todos/graph"
	"gqlgen-todos/infrastructure"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	services.Module,
	repository.Module,
	infrastructure.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	router infrastructure.Router,
	redisClient infrastructure.RedisClient,
	userService services.UserService,
) {

	graphQLServer := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{
			Resolvers: &graph.Resolver{
				UserService: userService,
			}}))

	// Define a route for the GraphQL Playground and query endpoint
	router.Gin.GET("/playground", func(c *gin.Context) {
		playground.Handler("GraphQL", "/query")(c.Writer, c.Request)
	})

	router.Gin.POST("/query", func(c *gin.Context) {
		graphQLServer.ServeHTTP(c.Writer, c.Request)
	})
	appStop := func(context.Context) error {
		fmt.Println("Stopping Application")
		_ = redisClient.RDB.Close()
		return nil
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting Application")
			fmt.Println("------------------------")

			go func() {
				router.Gin.Run(":" + "8000")
			}()
			return nil
		},
		OnStop: appStop,
	})
}
