package config

func ProvideLog() Log {
	return configService.Log
}

func ProvideDbMap() DbMap {
	return configService.DbMap
}
