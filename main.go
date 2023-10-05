package main

import (
	"backend/json-server/handler"
	"backend/json-server/helper"
	"backend/json-server/repository"
	"backend/json-server/service"
	"context"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//SetStartTime
	startTimeServer := time.Now()
	//Set Main Context
	ctx := context.Background()

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}

	db := helper.ConnectMongo(ctx)

	//Repository
	serverRepository := repository.NewServerRepository(db)

	//Service
	serverService := service.NewServerService(serverRepository)

	//Handler
	serverHandler := handler.NewServiceHandler(serverService)

	//Router
	//Init Router App
	app := gin.Default()

	app.Use(gin.Recovery())
	app.Use(gin.Logger())

	// cors	config
	cfg := cors.DefaultConfig()
	cfg.AllowOrigins = []string{"*"}
	cfg.AllowCredentials = true
	cfg.AllowMethods = []string{"*"}
	cfg.AllowHeaders = []string{"*"}

	app.Use(cors.New(cfg))
	app.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root: "views",
	})

	app.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
		return
	})

	//Route
	apiRoute := app.Group("/api")
	apiRoute.POST("/", serverHandler.InsertJSON)
	apiRoute.GET("/:id", serverHandler.FindByID)
	//========> Run App
	//Init Server
	srv := &http.Server{
		Addr:    ":" + viper.GetString("PORT"),
		Handler: app,
	}

	// graceful shutdown
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ... ", time.Since(startTimeServer).Seconds(), " s")

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()

	log.Println("Server exiting")
}
