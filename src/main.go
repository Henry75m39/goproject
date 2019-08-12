package main

import (
	"GRM/src/common/configs"
	"GRM/src/common/utils/db"
	"GRM/src/common/utils/log"
	"GRM/src/tms-srv/provider"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"go.uber.org/zap"
)

func main() {

	log.Init("GRM-TMS-MicroService")
	logger := log.Instance()

	//define a new micro service
	service := micro.NewService(
		micro.Name(configs.Namespace+configs.ServiceNameTMS),
		micro.Version("latest"),
	)

	//define service action
	service.Init(
		micro.Action(func(context *cli.Context) {
			//initialize LevelDB for application.
			db.InitLevelDB()
			logger.Info("Info", zap.Any("tms-srv", "tms-srv is starting now ..."))

			provider.TMSAPIProvider()
		}),

		micro.AfterStop(func() error {
			return nil
		}),

		micro.AfterStart(func() error {
			return nil
		}),
	)

	//start TMS micro service
	if err := service.Run(); err != nil {
		logger.Panic("TMS micro service startup failed")
	}
}
