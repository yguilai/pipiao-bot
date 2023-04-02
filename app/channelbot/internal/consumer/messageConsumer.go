package consumer

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/nsqio/go-nsq"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/yguilai/pipiao-bot/app/core/cmdset"
)

type MessageConsumer struct {
	c   *cmdset.CommandContainer
	api openapi.OpenAPI
}

func NewMessageConsumer(container *cmdset.CommandContainer, api openapi.OpenAPI) *MessageConsumer {
	return &MessageConsumer{
		c:   container,
		api: api,
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
			return replyMsg(c.api, &qqMsg, "暂不支持该指令")
		}
		return err
	}
	realCommand := c.c.MatchHandler(commands)
	arg := ""
	if realCommand.SubRestrain != nil {
		arg = commands[1]
	}
	if realCommand.RunE == nil {
		return replyMsg(c.api, &qqMsg, "该指令还在开发中")
	}
	replyContent, err := realCommand.RunE(&cmdset.CommandContext{
		Cmds: commands,
		Arg:  arg,
		Msg:  &qqMsg,
		Cmd:  realCommand,
	})
	if err != nil {
		return err
	}
	return replyMsg(c.api, &qqMsg, replyContent)
}

func replyMsg(api openapi.OpenAPI, msg *dto.Message, content string) error {
	ctx := context.Background()
	replyMessage := &dto.MessageToCreate{
		Content: content,
		MsgID:   msg.ID,
		MessageReference: &dto.MessageReference{
			MessageID: msg.ID,
		},
	}
	if msg.DirectMessage {
		directMessage, err := api.CreateDirectMessage(ctx, &dto.DirectMessageToCreate{
			SourceGuildID: msg.SrcGuildID,
			RecipientID:   msg.Author.ID,
		})
		if err != nil {
			return err
		}
		_, err = api.PostDirectMessage(ctx, directMessage, replyMessage)
		return err
	}
	_, err := api.PostMessage(context.Background(), msg.ChannelID, replyMessage)
	return err
}
