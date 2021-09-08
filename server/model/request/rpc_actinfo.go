package request

type ActInfo struct {
	RpcPort string `json:"rpcPort" binding:"required"`
	RpcHost string `json:"rpcHost" binding:"required"`
}
type ActWeight struct {
	ActInfo
	ActType string `json:"actType" binding:"required"`
}
