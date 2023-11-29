package usecase_test

import (
	"fww-core/internal/mocks"
	"fww-core/internal/usecase"
	"time"

	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"

	redis_utils "fww-core/internal/container/infrastructure/redis"
)

var (
	uc             usecase.UseCase
	repositoryMock *mocks.Repository
	adapterMock    *mocks.Adapter
	redisMock      redismock.ClientMock
	clientMock     *redis.Client
	timeTimeNow    = time.Now().Round(time.Minute)
	timeNow        = time.Now().Format("2006-01-02 15:04:05")
	dateTime       = time.Now().Format("2006-01-02")
	t              = time.Now()
	dateOnly       = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.UTC().Location())
)

func setup() {
	repositoryMock = &mocks.Repository{}
	adapterMock = &mocks.Adapter{}
	clientMock, redisMock = redismock.NewClientMock()
	redis_utils.InitRedisClient(clientMock)
	uc = usecase.New(repositoryMock, adapterMock, clientMock)
}
