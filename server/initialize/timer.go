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
			fmt.Println(detail)
			go func(detail config.Detail) {
				_, _ = global.SlotsTimer.AddTaskByFunc("ClearDB", global.SlotsConfig.Timer.Spec, func() {
					err := utils.ClearTable(global.SlotsDb, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(detail)
		}
	}
}
