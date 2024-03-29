package server

import (
	"fmt"
	"mongodb-trx/shared/gogen"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"mongodb-trx/shared/infrastructure/config"
	"mongodb-trx/shared/infrastructure/logger"
)

func NewGinHTTPHandlerDefault(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) GinHTTPHandler {
	return NewGinHTTPHandler(log, fmt.Sprintf(":%d", cfg.Server.Port), appData)
}

// GinHTTPHandler will define basic HTTP configuration with gracefully shutdown
type GinHTTPHandler struct {
	GracefullyShutdown
	Router *gin.Engine
}

func NewGinHTTPHandler(log logger.Logger, address string, appData gogen.ApplicationData) GinHTTPHandler {

	router := gin.Default()

	// PING API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, appData)
	})

	// CORS
	router.Use(cors.New(cors.Config{
		ExposeHeaders:   []string{"Data-Length"},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"},
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Content-Type", "Authorization"},
		MaxAge:          12 * time.Hour,
	}))

	return GinHTTPHandler{
		GracefullyShutdown: NewGracefullyShutdown(log, router, address),
		Router:             router,
	}
}

// RunApplication is implementation of RegistryContract.RunApplication()
func (r *GinHTTPHandler) RunApplication() {
	r.RunWithGracefullyShutdown()
}
