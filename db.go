package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"os"
	"sync"
)



var (
	MONGODB_URI string

	clientInstance    *mongo.Client
	clientInstanceErr error
	mongoOnce         sync.Once
	mongoDBName       string
)

func GetDB() (*MongoDB, error) {
	mongoOnce.Do(func() {
		_ = godotenv.Load()

		uri := os.Getenv("MONGODB_URI")
		mongoDBName = os.Getenv("MONGODB_DB_NAME")

		mode := os.Getenv("MODE")

		if mode == "DEBUG" || mode == "TEST" {
			mongoDBName = fmt.Sprintf("%s_TEST", mongoDBName)
		}
		if uri == "" || mongoDBName == "" {
			clientInstanceErr = fmt.Errorf("variáveis de ambiente MONGODB_URI ou MONGODB_DB_NAME não definidas")
			return
		}

		clientOptions := options.Client().ApplyURI(uri)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		clientInstance, clientInstanceErr = mongo.Connect(ctx, clientOptions)

		if clientInstanceErr == nil {
			clientInstanceErr = clientInstance.Ping(ctx, nil)
		}

	})

	if clientInstanceErr != nil {
		log.Println("Erro: erro na conexão com o banco de dados -> ", clientInstanceErr.Error())
		return nil, clientInstanceErr
	}

	db := clientInstance.Database(mongoDBName)
	return &MongoDB{Client: clientInstance, Database: db}, nil
}



