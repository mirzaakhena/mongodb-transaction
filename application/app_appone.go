package application

import (
	"mongodb-trx/domain_mytrx/controller/restapi"
	"mongodb-trx/domain_mytrx/gateway/withgorm"
	"mongodb-trx/domain_mytrx/usecase/createorder"
	"mongodb-trx/shared/gogen"
	"mongodb-trx/shared/infrastructure/config"
	"mongodb-trx/shared/infrastructure/logger"
	"mongodb-trx/shared/infrastructure/server"
)

type appone struct{}

func NewAppone() gogen.Runner {
	return &appone{}
}

func (appone) Run() error {

	const appName = "appone"

	cfg := config.ReadConfig()

	appData := gogen.NewApplicationData(appName)

	log := logger.NewSimpleJSONLogger(appData)

	datasource := withgorm.NewGateway(log, appData, cfg)

	httpHandler := server.NewGinHTTPHandler(log, ":8080", appData)

	x := restapi.NewGinController(log, cfg)
	x.AddUsecase(
		//
		createorder.NewUsecase(datasource),
	)
	x.RegisterRouter(httpHandler.Router)

	httpHandler.RunWithGracefullyShutdown()

	return nil
}
