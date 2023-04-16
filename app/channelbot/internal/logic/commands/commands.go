package commands

import (
	datastructure "github.com/duke-git/lancet/v2/datastructure/set"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/cmdset"
)

func RegisterCommands() *cmdset.CommandContainer {
	cc := cmdset.NewCommandContainer()

	cycleCmd := &cmdset.BotCommand{
		Use:         "/状态",
		SubRestrain: datastructure.NewSet("地球", "金星", "火卫二", "希图斯", "扎里曼"),
		RunE:        NewCycleCmdLogic,
	}

	alertsCmd := &cmdset.BotCommand{
		Use:  "/警报",
		RunE: nil,
	}

	sortieCmd := &cmdset.BotCommand{
		Use:  "/突击",
		RunE: nil,
	}

	invasionCmd := &cmdset.BotCommand{
		Use:  "/入侵",
		RunE: nil,
	}

	nightwaveCmd := &cmdset.BotCommand{
		Use:  "/午夜电波",
		RunE: nil,
	}

	voidTraderCmd := &cmdset.BotCommand{
		Use:  "/奸商",
		RunE: nil,
	}

	wmCmd := &cmdset.BotCommand{
		Use:         "/wm",
		SubRestrain: datastructure.NewSet("*"),
		RunE:        nil,
	}

	kuvaCmd := &cmdset.BotCommand{
		Use:  "/赤毒",
		RunE: nil,
	}

	arbitrationCmd := &cmdset.BotCommand{
		Use:  "/仲裁",
		RunE: nil,
	}

	archonHuntCmd := &cmdset.BotCommand{
		Use:  "/执行官",
		RunE: nil,
	}

	cc.Register(cycleCmd, alertsCmd, sortieCmd, invasionCmd, nightwaveCmd, voidTraderCmd, wmCmd, kuvaCmd,
		arbitrationCmd, archonHuntCmd)
	return cc
}
