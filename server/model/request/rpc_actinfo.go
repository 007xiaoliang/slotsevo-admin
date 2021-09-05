package request

type ActInfo struct {
	RpcPort    string `json:"rpcPort" binding:"required"`
	RpcHost    string `json:"rpcHost" binding:"required"`
}

