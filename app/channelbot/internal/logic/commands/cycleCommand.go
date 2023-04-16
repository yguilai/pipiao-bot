package commands

import (
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/cmdset"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/svc"
)

type CycleCmdLogic struct {
}

func NewCycleCmdLogic(svcCtx *svc.ServiceContext) cmdset.CommandHandler {
	return &CycleCmdLogic{}
}

func (c *CycleCmdLogic) Handle(ctx *cmdset.CommandContext) (string, error) {
	if !ctx.Cmd.SubRestrain.Contain(ctx.Arg) {
		return "", nil
	}
	return "解析成功", nil
}

// 地球 https://docs.warframestat.us/#tag/Worldstate/operation/getEarthByPlatform
// 希图斯 https://docs.warframestat.us/#tag/Worldstate/operation/getCetusByPlatform
// 金星 https://docs.warframestat.us/#tag/Worldstate/operation/getVallisByPlatform
// 火卫二 https://docs.warframestat.us/#tag/Worldstate/operation/getCambionByPlatform
// 扎利曼
