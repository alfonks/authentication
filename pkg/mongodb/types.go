package mongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type dbConfig struct {
	ConnectionURI string
	AppName       string
	Timeout       time.Duration
	MaxConnection uint64
	MaxPoolSize   uint64
	MaxIdleTime   time.Duration
}

type MongoDBInstance struct {
	Slave  *mongo.Database
	Master *mongo.Database
}
