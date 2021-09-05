package request

type ThemeInfo struct {
	RpcPort    string `json:"rpcPort" binding:"required"`
	RpcHost    string `json:"rpcHost" binding:"required"`
}

