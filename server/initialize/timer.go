package initialize

import (
	"fmt"
	"slotsevo-admin/config"
	"slotsevo-admin/global"
	"slotsevo-admin/utils"
)

func Timer() {
	if global.GVA_CONFIG.Timer.Start {
		for _, detail := range global.GVA_CONFIG.Timer.Detail {
			fmt.Println(detail)
			go func(detail config.Detail) {
				global.GVA_Timer.AddTaskByFunc("ClearDB", global.GVA_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.GVA_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(detail)
		}
	}
}
