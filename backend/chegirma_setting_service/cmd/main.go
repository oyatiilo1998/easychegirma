package main

import (
	"context"
	"fmt"
	"net"

	"chegirma_setting_service/config"
	"chegirma_setting_service/genproto/setting_service"
	"chegirma_setting_service/pkg/logger"
	"chegirma_setting_service/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	log logger.Logger
	cfg config.Config
)

func main() {
	cfg = config.Load()
	log = logger.New(cfg.LogLevel, "chegirma_setting_service")
	defer logger.Cleanup(log)

	log.Info("main: MONGOCONFIG",
		logger.String("Host", cfg.MongoHost),
		logger.Int("Port", cfg.MongoPort),
		logger.String("Database", cfg.MongoDatabase),
	)

	credential := options.Credential{
		Username: cfg.MongoUser,
		Password: cfg.MongoPassword,
	}

	mongoString := fmt.Sprintf("mongodb://%s:%d", cfg.MongoHost, cfg.MongoPort)

	mongoConn, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoString).SetAuth(credential))

	if err != nil {
		log.Error("error to connect to mongo database", logger.Error(err))
	}
	defer mongoConn.Disconnect(context.Background())
	if err := mongoConn.Ping(context.TODO(), nil); err != nil {
		log.Error("Cannot connect to database error ->", logger.Error(err))
		panic(err)
	}
	connDB := mongoConn.Database(cfg.MongoDatabase)
	log.Info("Connected to MongoDB", logger.Any("database: ", connDB.Name()))

	categoryService := service.NewCategoryService(connDB, log)
	productService := service.NewProductService(connDB, log)
	userService := service.NewUserService(connDB, log)

	listen, err := net.Listen("tcp", cfg.RPCPort)

	if err != nil {
		log.Error("Error while listening: %v", logger.Error(err))
		panic(err)
	}
	server := grpc.NewServer()
	setting_service.RegisterCategoryServiceServer(server, categoryService)
	setting_service.RegisterProductServiceServer(server, productService)
	setting_service.RegisterUserServiceServer(server, userService)
	reflection.Register(server)

	log.Info("main: server is running", logger.String("port", cfg.RPCPort))

	if err := server.Serve(listen); err != nil {
		log.Error("error while listening: %v", logger.Error(err))
		panic(err)
	}

}
