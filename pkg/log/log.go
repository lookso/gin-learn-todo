package log

import (
	"gin-learn-todo/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

func Init() {
	hook := lumberjack.Logger{
		Filename:   setting.Conf.Logger.Path, // 日志文件路径
		MaxSize:    1024,                        // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 3,                           // 日志文件最多保存多少个备份
		MaxAge:     7,                           // 文件最多保存多少天
		Compress:   true,                        // 是否压缩
	}
	defer hook.Close()
	// 调用库函数生成配置
	encoderConfig := zap.NewProductionEncoderConfig() //生成配置
	//时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.CallerKey = "linenum"

	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder // 全路径编码器
	encoderConfig.EncodeName = zapcore.FullNameEncoder

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	w := zapcore.AddSync(&hook)
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), //编码器配置
		w,                                        //打印到控制台和文件
		atomicLevel,                              //日志等级
	)
	logger = zap.New(core)
	return
}

func Logger() *zap.Logger {
	return logger
}

func Sugar() *zap.SugaredLogger {
	if sugar == nil {
		sugar = logger.Sugar()
	}
	return sugar
}

func Sync() error {
	return logger.Sync()
}