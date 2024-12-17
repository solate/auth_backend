package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DataSource string
	ShowSQL    bool
	Logx       Logx  `json:"Logx"`
	Redis      Redis `json:"Redis"`
}

type Logx struct {
	Mode     string
	Encoding string
	Level    string
	Stat     bool
}

type Redis struct {
	Host string
	Type string
	Pass string
	Port int
}

func (c Config) LoadLogConf() {
	logx.MustSetup(logx.LogConf{
		Mode:     c.Logx.Mode,
		Encoding: c.Logx.Encoding,
		Level:    c.Logx.Level,
		Stat:     c.Logx.Stat,
	})
}
