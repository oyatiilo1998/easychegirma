package main

import (
	"fmt"
	"time"

	"chegirma_api_gateway/api"
	"chegirma_api_gateway/config"
	"chegirma_api_gateway/pkg/logger"
	"chegirma_api_gateway/services"

	"github.com/go-redis/redis"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "chegirma_api_gateway")

	grpcClients, err := services.NewGrpcClients(cfg)
	if err != nil {
		log.Error("main.ErrorConnectionGRPCClients", logger.Error(err))
		panic(err)
	}

	redisConn := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
	})
	_, err = redisConn.Ping().Result()
	if err != nil {
		log.Error("main.ErrorConnectionRedis", logger.Error(err))
		panic(err)
	}
	defer redisConn.Close()

	redisManager := services.NewRedisClient(redisConn, grpcClients, time.Duration(cfg.RedisTimeoutInSeconds*int(time.Second)))

	// TODO: make sure kafka is working on server
	// TODO: Without this you cannot create action history
	// kafka := event.NewKafka(context.Background(), cfg, log)

	server := api.New(&api.RouterOptions{
		Log:         log,
		Cfg:         cfg,
		Services:    grpcClients,
		RedisClient: redisManager,
		// Kafka:       kafka,
	})
	server.Run(cfg.HttpPort)
}
