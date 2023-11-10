package container

import (
	"fww-core/internal/adapter"
	"fww-core/internal/config"
	"fww-core/internal/container/infrastructure/database"
	grpcclient "fww-core/internal/container/infrastructure/grpc/client"
	grpcserver "fww-core/internal/container/infrastructure/grpc/server"
	"fww-core/internal/container/infrastructure/http"
	"fww-core/internal/container/infrastructure/http/router"
	logger "fww-core/internal/container/infrastructure/log"
	messagestream "fww-core/internal/container/infrastructure/message_stream"
	"fww-core/internal/container/infrastructure/redis"
	"fww-core/internal/controller"
	"fww-core/internal/repository"
	"fww-core/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func InitService(cfg *config.Config) *fiber.App {
	// init database
	db := database.GetConnection(&cfg.Database)
	// init redis
	clientRedis := redis.SetupClient(&cfg.Redis)
	// init redis cache
	redis.InitRedisClient(clientRedis)
	// Init Tracing
	// Init Logger
	log := logger.Initialize(cfg)
	// Init HTTP Server
	server := http.SetupHttpEngine()
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

	// Init Publisher
	// Init Subscriber

	// Init Adapter
	adapter := adapter.NewBPM(nil, nil)
	// Init Repository
	repo := repository.New(db)
	// Init UseCase
	usecase := usecase.New(repo, adapter)
	// Init Controller
	ctrl := controller.Controller{UseCase: usecase, Log: log}

	// Init Router
	app := router.Initialize(server, &ctrl)

	return app
}
