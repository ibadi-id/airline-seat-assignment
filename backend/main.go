package main

import (
	"net/http"

	"github.com/ibadi-id/backend-bc/config"
	_ "github.com/ibadi-id/backend-bc/docs"
	"github.com/ibadi-id/backend-bc/internal/handler"
	"github.com/ibadi-id/backend-bc/internal/repository"
	"github.com/ibadi-id/backend-bc/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Airline Voucher API
// @version 1.0
// @description API for generating and checking airline voucher seats
// @host localhost:8080
// @BasePath /
func main() {
	db := config.InitDB()
	repo := repository.NewVoucherRepository(db)
	use := usecase.NewVoucherUsecase(repo)
	h := handler.NewHandler(use)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Swagger route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/api/check", h.Check)
	e.POST("/api/generate", h.Generate)

	e.Logger.Fatal(e.Start(":8080"))
}
