package restapi

import (
	"github.com/gin-gonic/gin"

	"mongodb-trx/domain_belajar/usecase/createorder"
	"mongodb-trx/shared/infrastructure/config"
	"mongodb-trx/shared/infrastructure/logger"
)

type Controller struct {
	Router            gin.IRouter
	Config            *config.Config
	Log               logger.Logger
	CreateOrderInport createorder.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.GET("/createorder", r.authorized(), r.createOrderHandler(r.CreateOrderInport))
}
