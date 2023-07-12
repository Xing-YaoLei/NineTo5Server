package config

type Server struct {
	Zap        Zap        `mapstructure:"zap" json:"zap" yaml:"zap"`                      // zap 日志配置
	Mysql      Mysql      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`                // mysql 配置
	RotateLogs RotateLogs `mapstructure:"rotateLogs" json:"rotateLogs" yaml:"rotateLogs"` // 日志切割配置
	System     System     `mapstructure:"system" json:"system" yaml:"system"`             // 系统配置
	Captcha    Captcha    `mapstructure:"captcha" json:"captcha" yaml:"captcha"`          // 验证码配置
}
