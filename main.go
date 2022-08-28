package main

import (
	"app/config"
	"app/routers"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var e *echo.Echo

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Arquivo .env não encontrado")
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	client, ctx := config.ConnectionDB()
	defer client.Disconnect(*ctx)

	routers.RoutersConfig(e, client, ctx)

	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%s", config.GetPort())))
}
