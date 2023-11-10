package usecase_test

import (
	"fww-core/internal/mocks"
	"fww-core/internal/usecase"
	"time"
)

var (
	uc             usecase.UseCase
	repositoryMock *mocks.Repository
	adapterMock    *mocks.Adapter
	timeNow        = time.Now().Format("2006-01-02 15:04:05")
	dateTime       = time.Now().Format("2006-01-02")
	t              = time.Now()
	dateOnly       = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.UTC().Location())
)

func setup() {
	repositoryMock = &mocks.Repository{}
	adapterMock = &mocks.Adapter{}
	uc = usecase.New(repositoryMock, adapterMock)
}
