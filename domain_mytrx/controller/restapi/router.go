package restapi

import (
	"github.com/gin-gonic/gin"
	"mongodb-trx/shared/gogen"
	"mongodb-trx/shared/infrastructure/config"
	"mongodb-trx/shared/infrastructure/logger"
)

type selectedRouter = gin.IRouter

type ginController struct {
	*gogen.BaseController
	log logger.Logger
	cfg *config.Config
}

func NewGinController(log logger.Logger, cfg *config.Config) gogen.RegisterRouterHandler[selectedRouter] {
	return &ginController{
		BaseController: gogen.NewBaseController(),
		log:            log,
		cfg:            cfg,
	}
}

func (r *ginController) RegisterRouter(router selectedRouter) {

	//resource := router.Group("/api/v1", r.authentication())
	router.GET("/trx", r.authorization(), r.createOrderHandler())

}
