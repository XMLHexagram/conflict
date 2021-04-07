package config

type Config struct {
	Server `toml:"Server"`
	DbMap  `toml:"DbMap"`
	Log    `toml:"Log"`
}

type Http struct {
	Port    string `toml:"Port"`
	Timeout int    `toml:"Timeout"`
}

type Server struct {
	Http Http `toml:"Http"`
}

type Db struct {
	Driver   string `toml:"Driver"`
	Address  string `toml:"Address"`
	User     string `toml:"User"`
	Password string `toml:"Password"`
	DBName   string `toml:"DBName"`
}

type DbMap map[string]Db

type Log struct {
	Level       string
	ToStd       bool
	LogRotate   bool
	Development bool
	Sampling    bool
	Rotate      Rotate
}

type Rotate struct {
	Filename   []string
	MaxSize    int
	MaxAge     int
	MaxBackups int
}
