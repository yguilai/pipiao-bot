package config

import (
	"github.com/yguilai/pipiao-bot/app/core/conf"
)

type Config struct {
	Bot conf.BotConf
	Nsq conf.NsqConf
	//Redis redis.RedisConf
}
