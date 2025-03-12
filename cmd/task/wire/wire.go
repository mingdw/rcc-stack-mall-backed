//go:build wireinject
// +build wireinject

package wire

import (
	"rcc-stake-mall-backed/internal/repository"
	"rcc-stake-mall-backed/internal/server"
	"rcc-stake-mall-backed/internal/task"
	"rcc-stake-mall-backed/pkg/app"
	"rcc-stake-mall-backed/pkg/log"
	"rcc-stake-mall-backed/pkg/sid"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
)

var taskSet = wire.NewSet(
	task.NewTask,
	task.NewUserTask,
)
var serverSet = wire.NewSet(
	server.NewTaskServer,
)

// build App
func newApp(
	task *server.TaskServer,
) *app.App {
	return app.NewApp(
		app.WithServer(task),
		app.WithName("demo-task"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		taskSet,
		serverSet,
		newApp,
		sid.NewSid,
	))
}
