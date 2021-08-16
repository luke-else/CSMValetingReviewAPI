package main

import (
	"CSMValetingReviewAPI/handlers"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	//middlewares

	//endpoints
	e.POST("/upload", handlers.Upload)
	e.POST("/view", handlers.View)

	//start server
	e.Logger.Fatal(e.Start(":80"))

}
