package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"key-value-store/api/v1/store"
	_ "key-value-store/docs"
)

// @title Key-Value Store API
// @description Key-Value Store
// @version 1.0
// @termsOfService https://swagger.io/terms/

// @contact.name Key-Value Store API Support
// @contact.url https://www.swagger.io/support
// @contact.email ce.emreerinc@gmail.com

// @license.name Apache 2.0
// @license.url https://www.apache.org/licenses/LICENSE-2.0.html

// @host google.com
// @BasePath /
func main() {
	// Echo instance
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	store.Handler(e.Group("/store"))

	store.FetchLastPersistedData() // restore data
	store.Routine() // start store persistence routine

	e.Logger.Fatal(e.Start(":9090"))
}
