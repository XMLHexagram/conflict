//+build wireinject

package db

import (
	"github.com/google/wire"
	"conflict-template/internal/service/config"
)

func InitDep() *service {
	wire.Build(
		provideDbMap,
		config.ProvideDbMap,
		wire.Struct(new(service), "*"),
	)

	//wire.Build(config.GetService().GetNekoMySQLConfig())
	return &service{}
}
