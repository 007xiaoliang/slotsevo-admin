package initialize

import (
	"fmt"
	"slotsevo-admin/config"
	"slotsevo-admin/global"
	"slotsevo-admin/utils"
)

func Timer() {
	if global.GvaConfig.Timer.Start {
		for _, detail := range global.GvaConfig.Timer.Detail {
			fmt.Println(detail)
			go func(detail config.Detail) {
				_, _ = global.GvaTimer.AddTaskByFunc("ClearDB", global.GvaConfig.Timer.Spec, func() {
					err := utils.ClearTable(global.GvaDb, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(detail)
		}
	}
}
