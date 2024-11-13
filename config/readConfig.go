package config

import "github.com/cdfmlr/crud/config"

//运行环境模式
type EnvMode string

const (
	Dev  EnvMode = "dev"  // 开发环境
	Test         = "test" // 测试环境
	Prod         = "prod" // 生产环境
)

func ReadConfigInfo(evn EnvMode) config.Option {

	return config.FromFile("./config/application_" + string(evn) + ".yaml")
}
