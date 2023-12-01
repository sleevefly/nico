package conf

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

func InitLog() {
	log := &lumberjack.Logger{
		Filename:   "/path/file.log", // 日志文件的位置
		MaxSize:    10,               // 文件最大尺寸（以MB为单位）
		MaxBackups: 3,                // 保留的最大旧文件数量
		MaxAge:     28,               // 保留旧文件的最大天数
		Compress:   true,             // 是否压缩/归档旧文件
		LocalTime:  true,             // 使用本地时间创建时间戳
	}
	logger := logrus.New()
	logger.Out = log
}
