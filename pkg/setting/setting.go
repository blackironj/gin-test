package setting

import (
	"log"

	"github.com/spf13/viper"
)

type Server struct {
	RunMode string
	Port    int
}

var ServerSetting Server

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
}

var DatabaseSetting Database

type Redis struct {
	Host     string
	Password string
}

var RedisSetting Redis

func Setup() {
	viper.SetConfigFile("./config.json")

	var err error
	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	server := viper.Sub("server")
	err = server.Unmarshal(&ServerSetting)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	database := viper.Sub("database")
	err = database.Unmarshal(&DatabaseSetting)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	redis := viper.Sub("redis")
	err = redis.Unmarshal(&RedisSetting)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}
