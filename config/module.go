package config

type Module struct {
	Public      ApiVersion `mapstructure:"public" json:"public" yaml:"public"`
	System      ApiVersion `mapstructure:"system" json:"system" yaml:"system"`
	Common      ApiVersion `mapstructure:"common" json:"common" yaml:"common"`
	Application ApiVersion `mapstructure:"application" json:"application" yaml:"application"`
}

type ApiVersion struct {
	ApiVersionEnabled bool   `mapstructure:"api-version-enabled" json:"api-version-enabled" yaml:"api-version-enabled"`
	ApiVersion        string `mapstructure:"api-version" json:"api-version" yaml:"api-version"`
}
