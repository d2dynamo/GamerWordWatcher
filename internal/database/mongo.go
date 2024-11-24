package database

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	MAX_POOL_SIZE      = 600
	MIN_POOL_SIZE      = 20
	MAX_CONN_IDLE_TIME = 3 * time.Minute
	APP_NAME           = "GamerWordWatcher"
)

var ctxMongo = context.Background()
var mongoClient *mongo.Client = nil

var mu sync.Mutex

func GetMongoDBClient() (*mongo.Client, *context.Context, error) {
	if mongoClient != nil {
		err := mongoClient.Ping(ctxMongo, readpref.Primary())
		if err == nil {
			return mongoClient, &ctxMongo, nil
		}
	}

	mu.Lock()
	defer mu.Unlock()

	if mongoClient != nil {
		err := mongoClient.Ping(ctxMongo, readpref.Primary())
		if err == nil {
			log.Println("Using existing mongo client")
			return mongoClient, &ctxMongo, nil
		}
	}

	log.Println("Creating new mongo client")

	maxPoolSize := uint64(MAX_POOL_SIZE)
	minPoolSize := uint64(MIN_POOL_SIZE)
	maxIdleTime := MAX_CONN_IDLE_TIME
	appName := APP_NAME

	uri := os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT")

	rootCA, err := os.ReadFile(os.Getenv("MONGO_CA_PATH"))
	if err != nil {
		fmt.Println("Error while loading root CA:", err)
		return nil, nil, err
	}

	rootCAs := x509.NewCertPool()
	if ok := rootCAs.AppendCertsFromPEM(rootCA); !ok {
		fmt.Println("Error while adding root CA to cert pool")
		return nil, nil, err
	}

	clientCert, err := tls.LoadX509KeyPair(os.Getenv("MONGO_CLIENT_CERT_PATH"), os.Getenv("MONGO_CLIENT_KEY_PATH"))
	if err != nil {
		fmt.Println("Error while loading client cert:", err)
		return nil, nil, err
	}

	tlsConfig := &tls.Config{
		RootCAs:            rootCAs,
		Certificates:       []tls.Certificate{clientCert},
		InsecureSkipVerify: true,
	}

	clientOptions := options.ClientOptions{
		Auth: &options.Credential{
			AuthMechanism: "MONGODB-X509",
			AuthSource:    "$external",
		},
		MaxPoolSize:     &maxPoolSize,
		MinPoolSize:     &minPoolSize,
		MaxConnIdleTime: &maxIdleTime,
		AppName:         &appName,
		TLSConfig:       tlsConfig,
	}

	clientOptions.ApplyURI(uri)

	m, err := mongo.NewClient(&clientOptions)
	if err != nil {
		return nil, nil, err
	}

	err = m.Connect(ctxMongo)
	if err != nil {
		return nil, nil, err
	}

	err = m.Ping(ctxMongo, readpref.Primary())
	if err != nil {
		return nil, nil, err
	}

	mongoClient = m

	// alreadyCreatingNewMongoClient = false

	return m, &ctxMongo, nil
}