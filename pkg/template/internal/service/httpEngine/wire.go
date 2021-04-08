//+build wireinject

package httpEngine

import (
	"github.com/google/wire"
	"conflict-template/internal/service/config"
)

func InitDep() *service {
	wire.Build(
		config.ProvideHttp,
		wire.Struct(new(service), "*"),
	)

	//wire.Build(config.GetService().GetNekoMySQLConfig())
	return &service{}
}
