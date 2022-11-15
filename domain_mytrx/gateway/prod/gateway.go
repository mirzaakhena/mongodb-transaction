package prod

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"mongodb-trx/domain_mytrx/model/entity"
	"mongodb-trx/shared/driver"
	"mongodb-trx/shared/infrastructure/config"
	"mongodb-trx/shared/infrastructure/database"
	"mongodb-trx/shared/infrastructure/logger"
	// "github.com/ostafen/clover"
)

type gateway struct {
	*database.MongoWithTransaction

	log     logger.Logger
	appData driver.ApplicationData
	config  *config.Config
}

// NewGateway ...
func NewGateway(log logger.Logger, appData driver.ApplicationData, cfg *config.Config) *gateway {

	cl := database.NewMongoDefault(cfg)

	prepareCollection(cl)

	return &gateway{
		log:                  log,
		appData:              appData,
		config:               cfg,
		MongoWithTransaction: database.NewMongoWithTransaction(cl),
	}
}

func prepareCollection(cl *mongo.Client) {
	db := cl.Database("cobadb")

	collectionNames := []string{
		"order",
		"person",
	}

	existingCollectionNames, err := db.ListCollectionNames(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}

	mapCollName := map[string]int{}
	for _, name := range existingCollectionNames {
		mapCollName[name] = 1
	}

	for _, name := range collectionNames {
		_, exist := mapCollName[name]
		if !exist {
			createCollection(db.Collection(name), db)
		}
	}
}

func createCollection(coll *mongo.Collection, db *mongo.Database) {
	createCmd := bson.D{{"create", coll.Name()}}
	res := db.RunCommand(context.Background(), createCmd)
	err := res.Err()
	if err != nil {
		panic(err)
	}
}

func (r *gateway) SaveOrder(ctx context.Context, obj *entity.Order) error {
	r.log.Info(ctx, "called SaveOrder")

	coll := r.MongoClient.Database("cobadb").Collection("order")

	res, err := coll.InsertOne(ctx, obj)
	if err != nil {
		r.log.Error(ctx, err.Error())
		return err
	}

	r.log.Info(ctx, "Order %v Inserted", res.InsertedID)

	return nil
}

func (r *gateway) SavePerson(ctx context.Context, obj *entity.Person) error {
	r.log.Info(ctx, "called SavePerson")

	coll := r.MongoClient.Database("cobadb").Collection("person")

	res, err := coll.InsertOne(ctx, obj)
	if err != nil {
		r.log.Error(ctx, err.Error())
		return err
	}

	r.log.Info(ctx, "Person %v Inserted", res.InsertedID)

	//return nil
	return fmt.Errorf("Uppss")
}
