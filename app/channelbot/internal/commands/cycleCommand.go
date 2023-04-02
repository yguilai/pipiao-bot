package commands

import (
	"github.com/yguilai/pipiao-bot/app/core/cmdset"
)

func CycleCommandHandler(ctx *cmdset.CommandContext) (string, error) {
	if !ctx.Cmd.SubRestrain.Contain(ctx.Arg) {
		return "", nil
	}
	return "解析成功", nil
}

// 地球 https://docs.warframestat.us/#tag/Worldstate/operation/getEarthByPlatform
// 希图斯 https://docs.warframestat.us/#tag/Worldstate/operation/getCetusByPlatform
// 金星 https://docs.warframestat.us/#tag/Worldstate/operation/getVallisByPlatform
// 火卫二 https://docs.warframestat.us/#tag/Worldstate/operation/getCambionByPlatform
