package cmdset

import (
	"errors"
	"strings"
)

var ErrCmdFormat = errors.New("消息格式不正确, 无法转成操作指令")

func Message2Commands(msg string) ([]string, error) {
	cmds := strings.Split(msg, " ")

	seemAt := cmds[0]
	if strings.HasPrefix(seemAt, "<@") && strings.HasSuffix(seemAt, ">") {
		if len(cmds) < 2 {
			return nil, ErrCmdFormat
		}
		cmds = cmds[1:]
	}

	return cmds, nil
}
