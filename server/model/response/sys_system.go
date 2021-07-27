package response

import "slotsevo-admin/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
