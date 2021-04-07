//+build wireinject

package db

import (
	"github.com/google/wire"
	"template/internal/service/config"
)

func InitDep() *service {
	wire.Build(
		config.ProvideDbMap,
		wire.Struct(new(service), "*"),
	)

	//wire.Build(config.GetService().GetNekoMySQLConfig())
	return &service{}
}
