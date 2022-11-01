package autoload

type MongoConfig struct {
	Enable   bool   `ini:"enable" yaml:"enable"`
	Host     string `ini:"host" yaml:"host"`
	Username string `ini:"username" yaml:"username"`
	Password string `ini:"password" yaml:"password"`
	Port     uint16 `ini:"port" yaml:"port"`
	Database string `ini:"database" yaml:"database"`
	LogLevel int    `ini:"log_level" yaml:"log_level"`
	PrintSql bool   `ini:"print_sql" yaml:"print_sql"`
}

var Mongo = MongoConfig{
	Enable:   false,
	Host:     "127.0.0.1",
	Username: "root",
	Password: "root1234",
	Port:     3306,
	Database: "test",
	LogLevel: 4,
	PrintSql: false,
}
