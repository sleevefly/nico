package conf

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type LogFormatter struct {
	logrus.JSONFormatter
}

func Init(logFilePath string) error {
	// 创建日志文件路径

	//err := os.MkdirAll(logFilePath, os.ModePerm)

	// 创建按小时切割的日志文件
	outputFile := logFilePath + "nico.log"

	if _, err := os.Stat(outputFile); err != nil {
		if os.IsNotExist(err) {
			if _, err := os.Create(outputFile); err != nil {
				return err
			}
		}
	}
	// 检查输出文件是否具有可写权限
	if err := os.Chmod(outputFile, 0644); err != nil {
		// 修改输出文件的权限失败
		return err
	}
	logWriter, err := rotatelogs.New(
		logFilePath+"nico.log",
		rotatelogs.WithLinkName(logFilePath+"log.log"),
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithMaxAge(time.Minute*2),
		rotatelogs.WithRotationSize(1024*1024*1024),
	)

	if err != nil {
		fmt.Println("Failed to create log file:", err)
		return err
	}

	logrus.SetOutput(logWriter)
	//级别月底，日志越小
	logrus.SetLevel(logrus.DebugLevel)

	defer logWriter.Close()
	return nil
}

func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	prefix := fmt.Sprintf("[%s] ", entry.Level.String())
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	entry.Message = prefix + entry.Message
	entry.Time = time.Now()
	return []byte(fmt.Sprintf("[%s] %s\n", timestamp, entry.Message)), nil
}

func Init_logrus(logFilePath string) {
	log := &lumberjack.Logger{
		Filename:   logFilePath + "file.log", // 日志文件的位置
		MaxSize:    1024 * 1024,              // 文件最大尺寸（以MB为单位）
		MaxBackups: 3,                        // 保留的最大旧文件数量
		MaxAge:     28,                       // 保留旧文件的最大天数
		Compress:   true,                     // 是否压缩/归档旧文件
		LocalTime:  true,                     // 使用本地时间创建时间戳

	}

	logrus.SetLevel(logrus.WarnLevel)
	logrus.SetOutput(log)
}
