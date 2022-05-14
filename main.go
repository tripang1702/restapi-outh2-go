package main

import (
	"log"
	"os"
	"restapi-oauth2-go/controller"

	_ "restapi-oauth2-go/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title REST API PROJECT Golang
// @version version(1.0)
// @description This is restful api project using golang.

// @contact.name Tripang
// @contact.url https://www.instagram.com/tripang_panggang/
// @contact.email rifalalfa1702@gmail.com

// @license.name ahmadrifal

// @host localhost:1323
// @BasePath /
func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${method}] ${time_rfc3339} | ${remote_ip} | ${latency_human} | ${uri} | ${status}\n",
	}))

	e.Use(middleware.Recover())

	ctrx := controller.NewController()

	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	database := os.Getenv("DB_NAME")

	ctrx.SetParamEnv(server, port, user, password, database)
	// end

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/cakes", ctrx.GetCake)
	e.GET("/cakes/:id", ctrx.GetCakebyId)
	e.POST("/cakes", ctrx.CreatePostCake)
	e.PATCH("/cakes/:id", ctrx.UpdatePatchCake)
	e.DELETE("/cakes/:id", ctrx.DeleteCake)

	// Start server
	e.Logger.Fatal(e.Start("localhost:1323"))
}
