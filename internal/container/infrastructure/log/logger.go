package logger

import (
	"fww-core/internal/config"
	"log"
	"os"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Initialize(cfg *config.Config) *zap.SugaredLogger {
	atom := zap.NewAtomicLevel()

	var core zapcore.Core
	zapOptions := []zap.Option{
		zap.AddStacktrace(zap.FatalLevel),
		zap.AddCaller(),
	}

	if cfg.Logger.LoggerCollector == "elastic" {
		// init ecszap logger
		encoderConfig := ecszap.NewDefaultEncoderConfig()
		if cfg.Logger.IsVerbose {
			atom.SetLevel(zap.DebugLevel)
		} else {
			atom.SetLevel(zap.ErrorLevel)
		}
		core = ecszap.NewCore(encoderConfig, os.Stdout, atom)

	} else {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
			atom,
		)
	}

	logger := zap.New(
		core,
		zapOptions...,
	)
	defer logSync(logger)
	logger = logger.With(zap.String("app", cfg.ServiceName))

	sugar := logger.Sugar()
	return sugar
}

func logSync(logger *zap.Logger) {
	err := logger.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
