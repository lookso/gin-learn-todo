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
		MaxSize:    1024,                     // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 3,                        // 日志文件最多保存多少个备份
		MaxAge:     7,                        // 文件最多保存多少天
		Compress:   true,                     // 是否压缩
	}
	defer hook.Close()
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	w := zapcore.AddSync(&hook)

	// 调用库函数生成配置
	// 写文件
	productionEncoderConfig := zap.NewProductionEncoderConfig()
	//时间格式
	productionEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	productionEncoderConfig.CallerKey = "caller"

	//productionEncoderConfig.EncodeCaller = zapcore.FullCallerEncoder // 全路径编码器
	productionEncoderConfig.EncodeName = zapcore.FullNameEncoder

	fileEncoder := zapcore.NewJSONEncoder(productionEncoderConfig)
	fileWriter := zapcore.NewCore(
		fileEncoder,
		w,
		atomicLevel,
	)

	// 控制台输出
	//consoleDebugging := zapcore.Lock(os.Stdout)
	//developmentEncoderConfig := zap.NewDevelopmentEncoderConfig()
	//developmentEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//consoleEncoder := zapcore.NewConsoleEncoder(developmentEncoderConfig)
	//consoleWriter := zapcore.NewCore(
	//	consoleEncoder,
	//	consoleDebugging,
	//	atomicLevel,
	//)
	//var allCore []zapcore.Core
	//allCore = append(allCore, fileWriter, consoleWriter)
	//core := zapcore.NewTee(allCore...).With([]zap.Field{
	//	//zap.String("app", "appName"),
	//})
	//logger = zap.New(
	//	core,
	//	zap.AddCaller(),
	//	zap.AddCallerSkip(1),
	//)

	logger = zap.New(
		fileWriter,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
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
