package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Logx       Logx `json:"Logx"`
	DataSource string
	ShowSQL    bool
}

type Logx struct {
	Mode     string
	Encoding string
	Level    string
	Stat     bool
}

func (c Config) LoadLogConf() {
	logx.MustSetup(logx.LogConf{
		Mode:     c.Logx.Mode,
		Encoding: c.Logx.Encoding,
		Level:    c.Logx.Level,
		Stat:     c.Logx.Stat,
	})
}
