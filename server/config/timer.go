package config

type Timer struct {
	Start  bool     `mapstructure:"start" json:"start" yaml:"start"` // 是否启用
	Spec   string   `mapstructure:"spec" json:"spec" yaml:"spec"`    // CRON表达式
	Detail []Detail `mapstructure:"detail" json:"detail" yaml:"detail"`
}

type Detail struct {
	TaskName string `mapstructure:"taskName" json:"taskName" yaml:"taskName"` // 任务名称
	KeepDays int    `mapstructure:"keepDays" json:"keepDays" yaml:"keepDays"` // 日志保留天数
}
