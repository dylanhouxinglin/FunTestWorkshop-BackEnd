package service

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type redisConf struct {
	Addr     string
	PoolSize int
	UserName string
	Passwd   string
}

func getRedisConf() (conf *redisConf, err error) {
	viper.SetConfigName("redis")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf/")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	conf = &redisConf{
		Addr:     viper.GetString("Addr"),
		PoolSize: viper.GetInt("PoolSize"),
		UserName: viper.GetString("UserName"),
		Passwd:   viper.GetString("Passwd"),
	}
	return
}

func ConnToRedis() (rdb *redis.Client, err error) {
	conf, err := getRedisConf()
	if err != nil {
		return
	}
	rdb = redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Passwd,
		PoolSize: conf.PoolSize,
		Username: conf.UserName,
	})
	return
}
