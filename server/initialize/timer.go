package initialize

import (
	"fmt"
	"slotsevo-admin/config"
	"slotsevo-admin/global"
	"slotsevo-admin/utils"
)

func Timer() {
	if global.SlotsConfig.Timer.Start {
		for _, detail := range global.SlotsConfig.Timer.Detail {
			go func(detail config.Detail) {
				_, _ = global.SlotsTimer.AddTaskByFunc(detail.TaskName, global.SlotsConfig.Timer.Spec, func() {
					err := utils.ClearLog(detail.KeepDays, global.SlotsConfig.Zap.Director)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(detail)
		}
	}
}
