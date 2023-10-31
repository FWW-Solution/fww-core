package container

import (
	"fww-core/internal/config"
	"fww-core/internal/container/infrastructure/database"
	grpcclient "fww-core/internal/container/infrastructure/grpc/client"
	grpcserver "fww-core/internal/container/infrastructure/grpc/server"
	logger "fww-core/internal/container/infrastructure/log"
	messagestream "fww-core/internal/container/infrastructure/message_stream"
	"fww-core/internal/container/infrastructure/redis"
)

func InitService(cfg *config.Config) {
	// init database
	_ = database.GetConnection(&cfg.Database)
	// init redis
	clientRedis := redis.SetupClient(&cfg.Redis)
	// init redis cache
	redis.InitRedisClient(clientRedis)
	// Init Tracing
	// Init Logger
	log := logger.Initialize(cfg)
	// Init GRPC Server
	grpcserver.Init(&cfg.GrpcServer)
	// Init GRPC Client
	_, err := grpcclient.Init(&cfg.GrpcClient)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	amqpMessageStream := messagestream.NewAmpq(&cfg.MessageStream)

	// set message stream subscriber
	_, err = amqpMessageStream.NewSubscriber()
	if err != nil {
		log.Error(err)
		panic(err)
	}

	// set message stream publisher
	_, err = amqpMessageStream.NewPublisher()
	if err != nil {
		log.Error(err)
		panic(err)
	}
}
