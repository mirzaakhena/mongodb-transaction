package database

import (
	"context"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteDefault() (db *gorm.DB) {

	db, err := gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}

// func NewPostgresDefault() (db *gorm.DB) {
//
// 	cfg, err := config.ReadConfig()
// 	if err != nil {
// 		panic(err.Error())
// 	}
//
// 	if cfg.User == "" || cfg.Password == "" || cfg.Database == "" {
// 		panic(fmt.Errorf("user or password ord databaseName is empty"))
// 	}
//
// 	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%v", cfg.Host, cfg.Port, cfg.User, cfg.Database, cfg.Password, cfg.SSLMode)
//
// 	loggerMode := logger.Silent
//
// 	if cfg.LogMode {
// 		loggerMode = logger.Info
// 	}
//
// 	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		Logger: logger.Default.LogMode(loggerMode),
// 	})
// 	if err != nil {
// 		panic(err.Error())
// 	}
//
// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		panic(err.Error())
// 	}
//
// 	sqlDB.SetMaxIdleConns(10)
//
// 	sqlDB.SetMaxOpenConns(10)
//
// 	sqlDB.SetConnMaxLifetime(10 * time.Second)
//
// 	return db
// }

type contextDBType string

var ContextDBValue contextDBType = "gormDB"

type gormWrapper struct {
	db *gorm.DB
}

// ExtractDB is used by other repo to extract the databasex from context
func (r *gormWrapper) ExtractDB(ctx context.Context) *gorm.DB {

	db, ok := ctx.Value(ContextDBValue).(*gorm.DB)
	if !ok {
		return r.db
	}

	return db
}

type GormWithoutTransactionImpl struct {
	*gormWrapper
}

func NewGormWithoutTransactionImpl(db *gorm.DB) *GormWithoutTransactionImpl {
	return &GormWithoutTransactionImpl{
		gormWrapper: &gormWrapper{db: db},
	}
}

func (r *GormWithoutTransactionImpl) GetDatabase(ctx context.Context) (context.Context, error) {
	trxCtx := context.WithValue(ctx, ContextDBValue, r.db)
	return trxCtx, nil
}

func (r *GormWithoutTransactionImpl) Close(ctx context.Context) error {
	return nil
}

// ---------------------------------------------------------------------------------------------

type GormWithTransactionImpl struct {
	*gormWrapper
}

func NewGormWithTransactionImpl(db *gorm.DB) *GormWithTransactionImpl {
	return &GormWithTransactionImpl{
		gormWrapper: &gormWrapper{db: db},
	}
}

func (r *GormWithTransactionImpl) BeginTransaction(ctx context.Context) (context.Context, error) {
	dbTrx := r.db.Begin()
	trxCtx := context.WithValue(ctx, ContextDBValue, dbTrx)
	return trxCtx, nil
}

func (r *GormWithTransactionImpl) CommitTransaction(ctx context.Context) error {
	return r.ExtractDB(ctx).Commit().Error
}

func (r *GormWithTransactionImpl) RollbackTransaction(ctx context.Context) error {
	return r.ExtractDB(ctx).Rollback().Error
}
