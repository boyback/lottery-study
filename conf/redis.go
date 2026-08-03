package conf

type RedisConfig struct {
	Host      string
	Port      string
	Username  string
	Password  string
	IsRunning bool
}

var RedisCacheList []RedisConfig = []RedisConfig{
	{
		Host:      "127.0.0.1",
		Port:      "6379",
		Username:  "root",
		Password:  "123456",
		IsRunning: true,
	},
}
var RedisCache RedisConfig = RedisCacheList[0]
