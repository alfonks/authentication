package mongodb

import (
	"context"
	"sync"
	"time"

	"deall-alfon/pkg/config"
	"deall-alfon/pkg/util/converter"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoDBSlaveOnce sync.Once
	mongoDBSlave     *mongo.Database

	mongoDBMasterOnce sync.Once
	mongoDBMaster     *mongo.Database
)

func GetMongoSlave() *mongo.Database {
	mongoDBSlaveOnce.Do(func() {
		ctx := context.Background()
		cfg := config.GetConfig()
		mongoDBSlave = connectDB(
			ctx,
			dbConfig{
				ConnectionURI: cfg.Mongo.Slave.URI,
				AppName:       cfg.Mongo.AppName,
				Timeout:       time.Duration(cfg.Mongo.Slave.Timeout) * time.Second,
				MaxConnection: converter.ToUInt64(cfg.Mongo.Slave.MaxConnection),
				MaxPoolSize:   converter.ToUInt64(cfg.Mongo.Slave.MaxPoolSize),
				MaxIdleTime:   time.Duration(cfg.Mongo.Slave.MaxIdleTime) * time.Second,
			},
		)
	})
	return mongoDBSlave
}

func GetMongoMaster() *mongo.Database {
	mongoDBMasterOnce.Do(func() {
		ctx := context.Background()
		cfg := config.GetConfig()
		mongoDBMaster = connectDB(
			ctx,
			dbConfig{
				ConnectionURI: cfg.Mongo.Master.URI,
				AppName:       cfg.Mongo.AppName,
				Timeout:       time.Duration(cfg.Mongo.Master.Timeout) * time.Second,
				MaxConnection: converter.ToUInt64(cfg.Mongo.Master.MaxConnection),
				MaxPoolSize:   converter.ToUInt64(cfg.Mongo.Master.MaxPoolSize),
				MaxIdleTime:   time.Duration(cfg.Mongo.Master.MaxIdleTime) * time.Second,
			},
		)
	})

	return mongoDBMaster
}

func GetMongoDBinstance() MongoDBInstance {
	return MongoDBInstance{
		Slave:  GetMongoSlave(),
		Master: GetMongoMaster(),
	}
}
