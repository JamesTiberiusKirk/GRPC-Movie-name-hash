package main

import (
	"github.com/JamesTiberiusKirk/moviehash/apiservice/controllers"
	"github.com/JamesTiberiusKirk/moviehash/common/hashmovieservice"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func runServer(port string, grpcClient hashmovieservice.HashMovieNameClient) {
	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.CORS(),
	)

	hmsController := controllers.NewHashMovieNameController(grpcClient)
	e.POST("/hash-movie-name", hmsController.HashMovieNameHandler)

	e.Logger.Fatal(e.Start(port))
}
