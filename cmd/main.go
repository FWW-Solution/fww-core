package main

import (
	"context"
	"fww-core/internal/config"
	"fww-core/internal/container"
	"fww-core/internal/container/infrastructure/http"
	"log"
)

func main() {
	// init config
	cfg := config.InitConfig()

	// init service
	app, routers := container.InitService(cfg)

	for _, router := range routers {
		ctx := context.Background()
		go func() {
			err := router.Run(ctx)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

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
	// 			zap.Int("cou nt", count))
	// 	}
	// 	count++
	// 	time.Sleep(time.Second * 2)
	// }

}
