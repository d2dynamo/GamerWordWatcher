package database

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	MAX_POOL_SIZE      = 200
	MIN_POOL_SIZE      = 10
	MAX_CONN_IDLE_TIME = 3 * time.Minute
	APP_NAME           = "GamerWordWatcher"
)

var ctxMongo = context.TODO()
var mongoClient *mongo.Client = nil

var mu sync.Mutex

func ConnectCollection(collName string) (*mongo.Collection, *context.Context, error) {
	db, ctx, err := GetClient()
	if err != nil {
		return nil, nil, err
	}

	dbName := os.Getenv("MONGO_DB_NAME")

	coll := db.Database(dbName).Collection(collName)

	return coll, ctx, nil
}

func GetClient() (*mongo.Client, *context.Context, error) {
	if mongoClient != nil {
		err := mongoClient.Ping(ctxMongo, readpref.Primary())
		if err == nil {
			return mongoClient, &ctxMongo, nil
		}
	}

	mu.Lock()
	defer mu.Unlock()

	maxPoolSize := uint64(MAX_POOL_SIZE)
	minPoolSize := uint64(MIN_POOL_SIZE)
	maxIdleTime := MAX_CONN_IDLE_TIME
	appName := APP_NAME

	caFile := os.Getenv("MONGO_CA_PATH")
	certFile := os.Getenv("MONGO_CERT_PATH")
	keyFile := os.Getenv("MONGO_KEY_PATH")

	caCert, err := os.ReadFile(caFile)
	if err != nil {
		panic(err)
	}
	caCertPool := x509.NewCertPool()

	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		panic("Error: CA file must be in PEM format")
	}

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		panic(err)
	}

	tlsConfig := &tls.Config{
		RootCAs:            caCertPool,
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}

	uri := os.Getenv("MONGO_URI")

	clientOpts := options.Client().ApplyURI(uri).SetAppName(appName).SetMaxPoolSize(maxPoolSize).SetMinPoolSize(minPoolSize).SetMaxConnIdleTime(maxIdleTime).SetTLSConfig(tlsConfig)

	m, err := mongo.Connect(ctxMongo, clientOpts)
	if err != nil {
		return nil, nil, err
	}

	err = m.Ping(ctxMongo, readpref.Primary())
	if err != nil {
		return nil, nil, err
	}

	mongoClient = m

	return m, &ctxMongo, nil
}
