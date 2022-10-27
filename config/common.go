package config

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

type Common struct {
	Zap Zap `yaml:"zap"`
}

type Zap struct {
	Level         string `yaml:"level"`           // 级别
	Prefix        string `yaml:"prefix"`          // 日志前缀
	Format        string `yaml:"format"`          // 输出
	Director      string `yaml:"director"`        // 日志文件夹
	EncodeLevel   string `yaml:"encode-level"`    // 编码级
	StacktraceKey string `yaml:"stacktrace-key"`  // 栈名
	MaxAge        int    `yaml:"max-age"`         // 日志留存时间
	ShowLine      bool   `yaml:"show-line"`       // 显示行
	LogInConsole  bool   `yaml:"log-in-console" ` // 输出控制台
}

func (z *Zap) ZapEncodeLevel() zapcore.LevelEncoder {
	switch {
	case z.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		return zapcore.LowercaseLevelEncoder
	case z.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case z.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case z.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

func (z *Zap) TransportLevel() zapcore.Level {
	z.Level = strings.ToLower(z.Level)
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
