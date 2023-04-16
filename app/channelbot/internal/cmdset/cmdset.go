package cmdset

import (
	datastructure "github.com/duke-git/lancet/v2/datastructure/set"
	"github.com/tencent-connect/botgo/dto"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/svc"
	"sync"
)

type (
	CommandContainer struct {
		root *BotCommand
		lock sync.RWMutex
	}

	BotCommand struct {
		Use         string                 `json:"use"`
		Subs        map[string]*BotCommand `json:"subs"`
		SubRestrain datastructure.Set[string]
		RunE        func(svcCtx *svc.ServiceContext) CommandHandler
	}

	CommandContext struct {
		Cmds []string

		Arg string
		Cmd *BotCommand
		Msg *dto.Message
	}

	CommandHandler interface {
		Handle(ctx *CommandContext) (string, error)
	}
)

const wildcard = "*"

func NewCommandContainer() *CommandContainer {
	return &CommandContainer{
		root: &BotCommand{Subs: make(map[string]*BotCommand)},
	}
}

func (c *CommandContainer) Register(cmds ...*BotCommand) {
	c.lock.Lock()
	defer c.lock.Unlock()
	for _, cmd := range cmds {
		c.root.Subs[cmd.Use] = cmd
	}
}

func (c *CommandContainer) MatchHandler(cmds []string) *BotCommand {
	c.lock.RLock()
	defer c.lock.RUnlock()
	m := c.root.Subs
	for i := 0; i < len(cmds); i++ {
		if realCmd, ok := m[cmds[i]]; ok {
			if realCmd.Subs == nil {
				return realCmd
			} else if realCmd.SubRestrain != nil && realCmd.SubRestrain.Contain(wildcard) {
				return realCmd
			}
			m = realCmd.Subs
		} else {
			return nil
		}
	}
	return nil
}

func (c *CommandContainer) AllCommands() map[string]*BotCommand {
	c.lock.RLock()
	defer c.lock.RUnlock()
	cmdMap := make(map[string]*BotCommand)
	for k, v := range c.root.Subs {
		cmdMap[k] = v
	}
	return cmdMap
}
