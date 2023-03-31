package msghandle

import (
	"errors"
	"github.com/yguilai/pipiao-bot/app/channelbot/types"
	"strings"
)

var ErrCmdFormat = errors.New("消息格式不正确, 无法转成操作命令")

func AtMessage2Command(msg string) (*types.BotCommand, error) {
	msgSlice := strings.Split(msg, " ")
	if len(msgSlice) < 2 {
		return nil, ErrCmdFormat
	}
	at := msgSlice[0]
	if !strings.HasPrefix(at, "<") || !strings.HasSuffix(at, ">") {
		return nil, ErrCmdFormat
	}
	cmd := &types.BotCommand{
		At:      at,
		MainCmd: msgSlice[1],
	}
	if len(msgSlice) >= 3 {
		cmd.SubCmd = msgSlice[3]
	}
	return cmd, nil
}
