package application

import (
	"mongodb-trx/domain_belajar/controller/restapi"
	"mongodb-trx/domain_belajar/gateway/prod"
	"mongodb-trx/domain_belajar/usecase/createorder"
	"mongodb-trx/shared/driver"
	"mongodb-trx/shared/infrastructure/config"
	"mongodb-trx/shared/infrastructure/logger"
	"mongodb-trx/shared/infrastructure/server"
	"mongodb-trx/shared/infrastructure/util"
)

type appOne struct {
	httpHandler *server.GinHTTPHandler
	controller  driver.Controller
}

func (c appOne) RunApplication() {
	c.controller.RegisterRouter()
	c.httpHandler.RunApplication()
}

func NewAppOne() func() driver.RegistryContract {
	return func() driver.RegistryContract {

		cfg := config.ReadConfig()

		appID := util.GenerateID(4)

		appData := driver.NewApplicationData("appOne", appID)

		log := logger.NewSimpleJSONLogger(appData)

		httpHandler := server.NewGinHTTPHandlerDefault(log, appData, cfg)

		datasource := prod.NewGateway(log, appData, cfg)

		return &appOne{
			httpHandler: &httpHandler,
			controller: &restapi.Controller{
				Log:               log,
				Config:            cfg,
				Router:            httpHandler.Router,
				CreateOrderInport: createorder.NewUsecase(datasource),
			},
		}

	}
}
