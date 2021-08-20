package config

type Rpc struct {
	Host             string `mapstructure:"host" json:"host" yaml:"host"`                                     // 远程服务器ip
	Port             string `mapstructure:"port" json:"port" yaml:"port"`                                     // 端口值
	Timeout          int    `mapstructure:"timeout" json:"timeout" yaml:"timeout"`                            // 连接超时时间
	RetryCount       int    `mapstructure:"retryCount" json:"retryCount" yaml:"retryCount"`                   // 重试次数
	RetryWaitTime    int    `mapstructure:"retryWaitTime" json:"retryWaitTime" yaml:"retryWaitTime"`          // 重试等待时间
	RetryMaxWaitTime int    `mapstructure:"retryMaxWaitTime" json:"retryMaxWaitTime" yaml:"retryMAxWaitTime"` // 重试最大等待时间
}
