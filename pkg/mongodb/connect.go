package mongodb

import (
	"context"
	"log"

	"deall-alfon/pkg/util/fn"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func connectDB(ctx context.Context, config dbConfig) *mongo.Database {
	op := fn.Name()
	client, err := mongo.Connect(
		ctx,
		options.Client().SetConnectTimeout(config.Timeout),
		options.Client().SetAppName(config.AppName),
		options.Client().SetMaxConnecting(config.MaxConnection),
		options.Client().SetMaxPoolSize(config.MaxPoolSize),
		options.Client().SetMaxConnIdleTime(config.MaxIdleTime),
		options.Client().ApplyURI(config.ConnectionURI),
	)
	if err != nil {
		log.Fatalf("[%v] error connect to db: %v", op, err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("[%v] fail ping to db", op)
	}

	return client.Database(deallDatabase)
}
