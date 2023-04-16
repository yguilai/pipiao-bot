package consumer

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/nsqio/go-nsq"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/cmdset"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
)

type MessageConsumer struct {
	c      *cmdset.CommandContainer
	api    openapi.OpenAPI
	svcCtx *svc.ServiceContext

	cmdCache map[string]cmdset.CommandHandler
	cmdLock  sync.Mutex
}

func NewMessageConsumer(svcCtx *svc.ServiceContext) *MessageConsumer {
	return &MessageConsumer{
		c:      svcCtx.CmdContainer,
		api:    svcCtx.Api,
		svcCtx: svcCtx,
	}
}

func (c *MessageConsumer) HandleMessage(nsqMsg *nsq.Message) error {
	var qqMsg dto.Message
	err := json.Unmarshal(nsqMsg.Body, &qqMsg)
	if err != nil {
		return err
	}
	commands, err := cmdset.Message2Commands(qqMsg.Content)

	if err != nil {
		if errors.Is(err, cmdset.ErrCmdFormat) {
			return c.replyMsg(&qqMsg, "暂不支持该指令")
		}
		return err
	}
	realCommand := c.c.MatchHandler(commands)
	arg := ""
	if realCommand.SubRestrain != nil {
		arg = commands[1]
	}
	if realCommand.RunE == nil {
		return c.replyMsg(&qqMsg, "该指令还在开发中")
	}

	handler := c.loadCommandHandler(realCommand)
	replyContent, err := handler.Handle(&cmdset.CommandContext{
		Cmds: commands,
		Arg:  arg,
		Msg:  &qqMsg,
		Cmd:  realCommand,
	})
	if err != nil {
		logx.Error(err)
		return c.replyMsg(&qqMsg, "指令处理出错")
	}
	return c.replyMsg(&qqMsg, replyContent)
}

func (c *MessageConsumer) loadCommandHandler(cmd *cmdset.BotCommand) cmdset.CommandHandler {
	c.cmdLock.Lock()
	defer c.cmdLock.Unlock()
	if handler, ok := c.cmdCache[cmd.Use]; ok {
		return handler
	}
	handler := cmd.RunE(c.svcCtx)
	c.cmdCache[cmd.Use] = handler
	return handler
}

func (c *MessageConsumer) replyMsg(msg *dto.Message, content string) error {
	ctx := context.Background()
	replyMessage := &dto.MessageToCreate{
		Content: content,
		MsgID:   msg.ID,
		MessageReference: &dto.MessageReference{
			MessageID: msg.ID,
		},
	}
	if msg.DirectMessage {
		directMessage, err := c.api.CreateDirectMessage(ctx, &dto.DirectMessageToCreate{
			SourceGuildID: msg.SrcGuildID,
			RecipientID:   msg.Author.ID,
		})
		if err != nil {
			return err
		}
		_, err = c.api.PostDirectMessage(ctx, directMessage, replyMessage)
		return err
	}
	_, err := c.api.PostMessage(context.Background(), msg.ChannelID, replyMessage)
	return err
}
