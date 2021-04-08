//+build wireinject

package log

import (
	"github.com/google/wire"
	"conflict-template/internal/service/config"
)

func initDep() *service {
	wire.Build(
		config.ProvideLog,
		wire.Struct(new(service), "*"),
	)

	//wire.Build(config.GetService().GetNekoMySQLConfig())
	return &service{}
}
