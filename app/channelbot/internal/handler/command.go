package handler

import (
	"github.com/nsqio/go-nsq"
	consumer2 "github.com/yguilai/pipiao-bot/app/channelbot/internal/consumer"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/svc"
	"github.com/zeromicro/go-zero/core/proc"
)

func RegisterNsqConsumer(svcCtx *svc.ServiceContext) error {
	c := svcCtx.Config
	nsqConfig := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(c.Nsq.Topic, "channel_bot", nsqConfig)
	if err != nil {
		return err
	}
	consumer.AddConcurrentHandlers(consumer2.NewMessageConsumer(svcCtx.CmdContainer, svcCtx.Api), 5)
	proc.AddShutdownListener(func() {
		consumer.Stop()
	})
	return consumer.ConnectToNSQLookupd(c.Nsq.Lookup)
}
