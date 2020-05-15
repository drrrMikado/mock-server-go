package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/drrrMikado/mock-server-go/conf"
	"github.com/drrrMikado/mock-server-go/handlers"
	"github.com/drrrMikado/mock-server-go/middlewares"
	"github.com/drrrMikado/mock-server-go/services"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}
	s := services.New(conf.Conf)
	handlers.Init(s)
	r := gin.Default()
	r.Use(middlewares.Cors())

	mockGroup := r.Group("/mock")
	{
		mockGroup.GET("/*uri", handlers.Handler)
		mockGroup.POST("/*uri", handlers.Handler)
	}
	adminGroup := r.Group("/admin")
	{
		adminGroup.GET("/mock/:id", handlers.Get)
		adminGroup.GET("/mock", handlers.List)
		adminGroup.POST("/mock", handlers.Add)
		adminGroup.PUT("/mock", handlers.Update)
		adminGroup.DELETE("/mock/:id", handlers.Delete)
	}
	go func() { // pprof
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	log.Fatalln(r.Run())
}
