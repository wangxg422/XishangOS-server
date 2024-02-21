package config

type Jwt struct {
	TokenName    string `mapstructure:"token-name" json:"token-name" yaml:"token-name"`
	SigningKey   string `mapstructure:"signing-key" json:"signing-key" yaml:"signing-key"`
	ExpiresTime  string `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"`
	BufferTime   string `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`
	Issuer       string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
	RedisKey     string `mapstructure:"redis-key" json:"redis-key" yaml:"redis-key"`
	BlackListKey string `mapstructure:"black-list-key" json:"black-list-key" yaml:"black-list-key"`
}
