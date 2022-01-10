package routes

import (
	"grpc-rest/api-protocol/rest/movies"

	"github.com/labstack/echo/v4"

	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func HandlerApi(e *echo.Echo, cont *movies.Controller) {
	usersRoutes := e.Group("/movies")

	usersRoutes.Use(echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	usersRoutes.GET("", cont.GetMovies)
	usersRoutes.GET("/id", cont.GetMoviesById)

}
