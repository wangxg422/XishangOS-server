package config

type Redis struct {
	Address  string `mapstructure:"address" json:"address" yaml:"address"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Db       int    `mapstructure:"initial" json:"initial" yaml:"initial"`
}
