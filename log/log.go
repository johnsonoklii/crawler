package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

// zapcore.Core 定义了日志的编码格式以及输出位置等核心功能
type Plugin = zapcore.Core

func NewLogger(plugin Plugin, opts ...zap.Option) *zap.Logger {
	return zap.New(plugin, append(DefaultOption(), opts...)...)
}

func NewPlugin(writer zapcore.WriteSyncer, enabler zapcore.LevelEnabler) Plugin {
	return zapcore.NewCore(DefaultEncoder(), writer, enabler)
}

func NewStdoutPlugin(enabler zapcore.LevelEnabler) Plugin {
	return NewPlugin(zapcore.Lock(zapcore.AddSync(os.Stdout)), enabler)
}

func NewStderrPlugin(enabler zapcore.LevelEnabler) Plugin {
	return NewPlugin(zapcore.Lock(zapcore.AddSync(os.Stderr)), enabler)
}

// Lumberjack logger虽然持有File但没有暴露sync方法，所以没办法利用zap的sync特性
// 所以额外返回一个closer，需要保证在进程退出前close以保证写入的内容可以全部刷到到磁盘
func NewFilePlugin(filePath string, enabler zapcore.LevelEnabler) (Plugin, io.Closer) {
	var writer = DefaultLumberjackLogger()
	writer.Filename = filePath

	return NewPlugin(zapcore.AddSync(writer), enabler), writer
}
