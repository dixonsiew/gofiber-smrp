package database

import (
    "context"
    "fmt"
    "smrp/config"
    "smrp/utils"
    "time"

    "github.com/jackc/pgx/v5/pgxpool"
    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
)

var dbVar *pgxpool.Pool
var ctx = context.Background()
var clientVar *mongo.Client
var ctxm = context.TODO()

func SetDb(db *pgxpool.Pool) {
    dbVar = db
}

func GetDb() *pgxpool.Pool {
    if dbVar == nil {
        ConnectDB()
    }

    return dbVar
}

func GetCtx() context.Context {
    return ctx
}

func ConnectDB() {
    name := config.Config("postgres.db")
    cfg, err := pgxpool.ParseConfig(fmt.Sprintf("postgres://postgres:postgres@localhost:5432/%s", name))
    if err != nil {
        utils.LogError(err)
        return
    }

    cfg.MaxConns = 10
    cfg.MinConns = 2
    cfg.MaxConnLifetime = time.Hour
    cfg.MaxConnIdleTime = 30 * time.Minute

    pool, err := pgxpool.NewWithConfig(ctx, cfg)
    if err != nil {
        utils.LogError(err)
    } else {
        SetDb(pool)
    }
}

func CloseDB() {
    dbVar.Close()
}

func ConnectMongo() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(clientOptions)
    if err != nil {
        utils.LogError(err)
    } else {
        clientVar = client
    }
}

func GetMongoClient() *mongo.Client {
    return clientVar
}

func GetMongoCtx() context.Context {
    return ctxm
}

func CloseMongo() {
    clientVar.Disconnect(ctxm)
}
