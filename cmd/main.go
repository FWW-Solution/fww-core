package main

import (
	"fww-core/internal/config"
	"fww-core/internal/container"
	"fww-core/internal/container/infrastructure/http"
)

func main() {
	// init config
	cfg := config.InitConfig()

	// init service
	app := container.InitService(cfg)

	// run service
	http.StartHttpServer(app, cfg.HttpServer.Port)

	// // init config
	// cfg := config.InitConfig()

	// log := logger.Initialize(cfg)
	// count := 0
	// for {

	// 	if rand.Float32() > 0.8 {
	// 		log.Error("oops...something is wrong",
	// 			zap.Int("count", count),
	// 			zap.Error(errors.New("error details")))
	// 	} else {
	// 		log.Info("everything is fine",
	// 			zap.Int("count", count))
	// 	}
	// 	count++
	// 	time.Sleep(time.Second * 2)
	// }

}
