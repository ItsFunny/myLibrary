/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package flogging

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io"
	"log"
	"os"

	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/grpclog"
)

const (
	defaultFormat = "%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}"
	defaultLevel  = zapcore.InfoLevel
)

var Global *Logging

func init() {
	logging, err := New(Config{})
	if err != nil {
		panic(err)
	}

	Global = logging
	grpcLogger := Global.ZapLogger("grpc")
	grpclog.SetLogger(NewGRPCLogger(grpcLogger))
}

// Init initializes logging with the provided config.
func Init(config Config) {
	err := Global.Apply(config)
	if err != nil {
		panic(err)
	}
}

// Reset sets logging to the defaults defined in this package.
//
// Used in tests and in the package init
func Reset() {
	Global.Apply(Config{})
}

// LoggerLevel gets the current logging level for the logger with the
// provided name.
func LoggerLevel(loggerName string) string {
	return Global.Level(loggerName).String()
}

// MustGetLogger creates a logger with the specified name. If an invalid name
// is provided, the operation will panic.
func MustGetLogger(loggerName string) *FabricLogger {
	return Global.Logger(loggerName)
}

// FileExists checks whether the given file exists.
// If the file exists, this method also returns the size of the file.
func fileExists(filePath string) (bool, int64, error) {
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false, 0, nil
	}
	if err != nil {
		return false, 0, errors.Wrapf(err, "error checking if file [%s] exists", filePath)
	}
	return true, fileInfo.Size(), nil
}
func MustFileLogger(loggerName string, fileName string) *zap.SugaredLogger {
	var file *os.File
	if exist, _, e := fileExists(fileName); nil != e {
		log.Fatal(e)
	} else if !exist {
		if filePtr, e := os.Create(fileName); nil != e {
			log.Fatal(e)
		} else {
			file = filePtr
		}
	} else {
		file, _ = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0666)
	}
	sync := zapcore.AddSync(file)
	core := zapcore.NewCore(getEncoder(), sync, zapcore.DebugLevel)
	lg := zap.New(core, zap.AddCaller())
	return lg.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 修改时间编码器

	// 在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// NewConsoleEncoder 打印更符合人们观察的方式
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// ActivateSpec is used to activate a logging specification.
func ActivateSpec(spec string) {
	err := Global.ActivateSpec(spec)
	if err != nil {
		panic(err)
	}
}

// DefaultLevel returns the default log level.
func DefaultLevel() string {
	return defaultLevel.String()
}

// SetWriter calls SetWriter returning the previous value
// of the writer.
func SetWriter(w io.Writer) io.Writer {
	return Global.SetWriter(w)
}

// SetObserver calls SetObserver returning the previous value
// of the observer.
func SetObserver(observer Observer) Observer {
	return Global.SetObserver(observer)
}
