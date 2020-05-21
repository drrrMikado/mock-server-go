package http

import (
	"context"
	"log"
	"net/http"

	"github.com/drrrMikado/mock-server-go/middlewares"
	"github.com/drrrMikado/mock-server-go/services"
	"github.com/gin-gonic/gin"
)

var (
	svc *services.Service
	svr []*http.Server
)

func Init(s *services.Service) {
	gin.SetMode(gin.ReleaseMode)
	svc = s
	// mock
	mockEngine := gin.Default()
	mockEngine.Use(middlewares.Cors())
	mockGroup := mockEngine.Group("/")
	{
		mockGroup.GET("/*uri", mockHandler)
		mockGroup.POST("/*uri", mockHandler)
	}
	// admin
	adminEngine := gin.Default()
	adminEngine.Use(middlewares.Cors())
	adminGroup := adminEngine.Group("/")
	{
		adminGroup.GET("/mock/:id", get)
		adminGroup.GET("/mock", list)
		adminGroup.POST("/mock", add)
		adminGroup.PUT("/mock", update)
		adminGroup.DELETE("/mock/:id", del)
	}

	svr = append(svr, &http.Server{Addr: ":8080", Handler: adminEngine})
	svr = append(svr, &http.Server{Addr: ":8081", Handler: mockEngine})
	svr = append(svr, &http.Server{Addr: ":6060", Handler: nil}) // pprof

	run()
}

func run() {
	for _, s := range svr {
		go func(svr *http.Server) {
			if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("Server listen: %s\n", err)
			}
		}(s)
	}
}

func Close(ctx context.Context) (err error) {
	for _, s := range svr {
		if err = s.Shutdown(ctx); err != nil {
			return
		}
	}
	return
}
