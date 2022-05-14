package main

import (
	"fmt"
	"log"
	"os"
	"restapi-oauth2-go/controller"

	_ "restapi-oauth2-go/docs"

	echoSwagger "restapi-oauth2-go/echo-swagger"

	echoserver "github.com/dasjott/oauth2-echo-server"
	"github.com/joho/godotenv"
	echoold "github.com/labstack/echo"
	middlewareold "github.com/labstack/echo/middleware"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

type Usercred struct {
	Username string
	Password string
	Id       string
}

// @title REST API PROJECT Golang
// @version version(1.0)
// @description This is restful api project using golang.

// @contact.name Tripang
// @contact.url https://www.instagram.com/tripang_panggang/
// @contact.email rifalalfa1702@gmail.com

// @license.name ahmadrifal

// @host localhost:1323
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in query
// @name access_token

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// end

	manager := manage.NewDefaultManager()

	// token store
	manager.MustTokenStorage(store.NewFileTokenStore("data.db"))

	userCredList := []Usercred{
		{
			Username: os.Getenv("username1"),
			Password: os.Getenv("password1"),
			Id:       os.Getenv("userid1"),
		},
		{
			Username: os.Getenv("username2"),
			Password: os.Getenv("password2"),
			Id:       os.Getenv("userid2"),
		},
	}

	// client store
	clientStore := store.NewClientStore()
	clientStore.Set("admin", &models.Client{
		ID:     "admin",
		Secret: "adminxyz",
		Domain: "http://localhost",
		UserID: "admin",
	})
	manager.MapClientStorage(clientStore)

	// Initialize the oauth2 service
	echoserver.InitServer(manager)
	echoserver.SetAllowGetAccessRequest(true)
	echoserver.SetClientInfoHandler(server.ClientFormHandler)
	echoserver.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		for _, uc := range userCredList {
			if uc.Username == username && uc.Password == password {
				userID = uc.Id
				return
			}
		}
		err = fmt.Errorf("invalid username and password")
		return

	})

	eold := echoold.New()

	// Middleware
	eold.Use(middlewareold.LoggerWithConfig(middlewareold.LoggerConfig{
		Format: "[${method}] ${time_rfc3339} | ${remote_ip} | ${latency_human} | ${uri} | ${status}\n",
	}))

	eold.Use(middlewareold.Recover())

	ctrx := controller.NewController()
	server := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	database := os.Getenv("DB_NAME")

	ctrx.SetParamEnv(server, port, user, password, database)

	eold.GET("/swagger/*", echoSwagger.WrapHandler)

	auth := eold.Group("/oauth2")
	{
		auth.GET("/token", echoserver.HandleTokenRequest)
	}

	api := eold.Group("/api")
	{
		api.Use(echoserver.TokenHandler())
		api.GET("/cakes", ctrx.GetCake)
		api.GET("/cakes/:id", ctrx.GetCakebyId)
		api.POST("/cakes", ctrx.CreatePostCake)
		api.PATCH("/cakes/:id", ctrx.UpdatePatchCake)
		api.DELETE("/cakes/:id", ctrx.DeleteCake)

	}

	// Start server
	// e.Logger.Fatal(e.Start("localhost:1323"))
	eold.Logger.Fatal(eold.Start("localhost:1323"))
}
