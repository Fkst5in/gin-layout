package data

import (
	"context"
	"fmt"

	c "github.com/wannanbigpig/gin-layout/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database

func initMongo() {
	// TODO

	url := fmt.Sprintf("mongodb://%s:%s@%s:%d",
		c.Config.Mongo.Username,
		c.Config.Mongo.Password,
		c.Config.Mongo.Host,
		c.Config.Mongo.Port,
	)

	clientOptions := options.Client().ApplyURI(url)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic("Mongo open connection failed：" + err.Error())
	}

	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic("Mongo ping connection failed：" + err.Error())
	}

	MongoDB = client.Database(c.Config.Mongo.Database)
}
