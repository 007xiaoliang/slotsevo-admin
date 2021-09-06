package model

type ActivityTypeResponse struct {
	ServerTime       string      `json:"server_time"`
	ActivityTypeList []string `json:"activity_type_list"`
}
