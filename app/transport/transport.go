package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/log"
	"github.com/tencent-connect/botgo/token"
	"github.com/yguilai/pipiao-bot/app/transport/internal/config"
	"github.com/yguilai/pipiao-bot/app/transport/internal/handler"
	"github.com/yguilai/pipiao-bot/app/transport/internal/svc"
	"github.com/yguilai/pipiao-bot/pkg/logadapter"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/%s/config.yaml", "the config file")
var mode = flag.String("env", "dev", "the environment")

func main() {
	var c config.Config
	cfgPath := fmt.Sprintf(*configFile, *mode)
	conf.MustLoad(cfgPath, &c)

	log.DefaultLogger = logadapter.NewBotgoLogger()
	botToken := token.BotToken(c.Bot.AppId, c.Bot.Token)
	api := botgo.NewOpenAPI(botToken)
	ws, err := api.WS(context.Background(), nil, "")
	handleError(err)

	svcCtx, err := svc.NewServiceContext(c)
	handleError(err)

	intent := handler.Register(svcCtx)
	manager := botgo.NewSessionManager()
	err = manager.Start(ws, botToken, &intent)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		logx.Error(err)
		panic(err)
	}
}
