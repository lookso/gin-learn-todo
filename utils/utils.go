package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetMinVer(v string) (uint64, error) {
	first := strings.IndexByte(v, '.')    // 3
	last := strings.LastIndexByte(v, '.') // 6
	if first == last {
		return strconv.ParseUint(v[first+1:], 10, 64)
	}
	return strconv.ParseUint(v[first+1:last], 10, 64)
}

func DebugPrint(format string, values ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	// os
	// var (
	//	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	//	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	//	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
	//)
	//Stdin、Stdout和Stderr是指向标准输入、标准输出、标准错误输出的文件描述符。
	//
	fmt.Fprintf(os.Stdout, "[GIN-debug] "+format, values...)
}

