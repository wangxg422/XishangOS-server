package config

type Database struct {
	Type  string `mapstructure:"type" json:"type" yaml:"type"`    // 数据库类型
	Mysql Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"` // mysql数据库配置
}
type Mysql struct {
	Address     string `mapstructure:"address" json:"address" yaml:"address"`                   // 服务器地址:端口
	Port        string `mapstructure:"port" json:"port" yaml:"port"`                            //:端口
	ConnConfig  string `mapstructure:"conn-config" json:"conn-config" yaml:"conn-config"`       // 高级配置
	Dbname      string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`                   // 数据库名
	Username    string `mapstructure:"username" json:"username" yaml:"username"`                // 数据库用户名
	Password    string `mapstructure:"password" json:"password" yaml:"password"`                // 数据库密码
	MaxIdleConn int    `mapstructure:"max-idle-conn" json:"max-idle-conn" yaml:"max-idle-conn"` // 空闲中的最大连接数
	MaxOpenConn int    `mapstructure:"max-open-conn" json:"max-open-conn" yaml:"max-open-conn"` // 打开到数据库的最大连接数
}
