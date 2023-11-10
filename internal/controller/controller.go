package controller

import (
	"fww-core/internal/usecase"

	"go.uber.org/zap"
)

type Controller struct {
	UseCase usecase.UseCase
	Log     *zap.SugaredLogger
}
