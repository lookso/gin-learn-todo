/*
@Time : 2019-11-24 16:14 
@Author : Tenlu
@File : rotate
@Software: GoLand
*/
package zaplog

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"time"
)

// writer 返回按天分隔的writer
func writer(filename string, age int64) io.Writer {
	// 实际生成的文件名 demo.log.YYmmdd
	hook, err := rotatelogs.New(
		filename+".%Y%m%d", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if err != nil {
		// 日志初始化不成功有可能是文件权限问题，所以要
		panic(err)
	}
	return hook
}
