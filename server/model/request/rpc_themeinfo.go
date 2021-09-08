package request

type ThemeInfo struct {
	RpcPort string `json:"rpcPort" binding:"required"`
	RpcHost string `json:"rpcHost" binding:"required"`
}

type ThemeWeight struct {
	ThemeInfo
	ThemeID string `json:"themeID" binding:"required"`
}
