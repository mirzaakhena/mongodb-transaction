package withgorm

import (
	"context"
	"fmt"
	"mongodb-trx/domain_mytrx/model/entity"
	"mongodb-trx/shared/gogen"
	"mongodb-trx/shared/infrastructure/config"
	"mongodb-trx/shared/infrastructure/database"
	"mongodb-trx/shared/infrastructure/logger"
)

type gateway struct {
	*database.GormWithTransactionImpl

	log     logger.Logger
	appData gogen.ApplicationData
	config  *config.Config
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *gateway {

	db := database.NewSQLiteDefault()

	err := db.AutoMigrate(entity.Order{}, entity.Person{})
	if err != nil {
		panic(err)
	}

	return &gateway{
		log:                     log,
		appData:                 appData,
		config:                  cfg,
		GormWithTransactionImpl: database.NewGormWithTransactionImpl(db),
	}
}

func (r *gateway) SaveOrder(ctx context.Context, obj *entity.Order) error {
	r.log.Info(ctx, "called")

	err := r.ExtractDB(ctx).Save(obj).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gateway) SavePerson(ctx context.Context, obj *entity.Person) error {
	r.log.Info(ctx, "called")

	err := r.ExtractDB(ctx).Save(obj).Error
	if err != nil {
		return err
	}

	return fmt.Errorf("Error here...")
}
