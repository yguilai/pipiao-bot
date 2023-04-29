package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tencent-connect/botgo"
	botlog "github.com/tencent-connect/botgo/log"
	"github.com/tencent-connect/botgo/token"
	hdl "github.com/yguilai/pipiao-bot/app/channelbot/internal/biz/handler"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/conf"
	"github.com/yguilai/pipiao-bot/pkg/consts"
	"github.com/yguilai/pipiao-bot/pkg/logadapter"
)

type BotServer struct {
	sm botgo.SessionManager
	c  *conf.Data
	eh *hdl.EventHandler
	*log.Helper
}

func NewBotServer(logger log.Logger, eh *hdl.EventHandler, c *conf.Data) *BotServer {
	return &BotServer{
		sm:     botgo.NewSessionManager(),
		Helper: log.NewHelper(log.With(logger, consts.ModuleKey, "transport.server.botserver")),
		eh:     eh,
		c:      c,
	}
}

func (m *BotServer) Start(ctx context.Context) error {
	botlog.DefaultLogger = logadapter.NewBotgoLogger(m.Helper)
	botToken := token.BotToken(m.c.Bot.AppId, m.c.Bot.Token)
	api := botgo.NewOpenAPI(botToken)
	ws, err := api.WS(ctx, nil, "")
	if err != nil {
		return err
	}
	intent := m.eh.Register()
	return m.sm.Start(ws, botToken, &intent)
}

func (m BotServer) Stop(_ context.Context) error {
	m.Info("bot server stop...")
	return nil
}
