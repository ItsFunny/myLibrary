/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-27 14:08
# @File : debug.go
# @Description :
# @Attention :
*/
package debugutil

import (
	"fmt"
	logplugin "github.com/hyperledger/fabric-droplib/base/log"
	libutils "github.com/hyperledger/fabric-droplib/base/utils"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/constants"
	"github.com/hyperledger/fabric/trace"
	"github.com/hyperledger/fabric/utils"
	"go.uber.org/zap"
	"os"
	"strconv"
	"strings"
)

var (
	// logg          *zap.SugaredLogger
	// consoleLogger *flogging.FabricLogger
	islog bool
)

var (
	defaultBlackList = []string{"debugutil/debug", "utils/log"}
)

func RegisterBlackList(bl ...string) {
	for _, v := range bl {
		libutils.RegisterBlackList(v)
	}
}

func UpdateLogValue(logAble bool) {
	islog = logAble

}
func init() {
	prepareDebugAble()
	RegisterBlackList("debugutil/debug")
	RegisterBlackList("utils/log","log/g_log")
	// config.InitViper(nil, "core")
	// viper.ReadInConfig()
	// peerName := viper.GetString("peer.id")
	// logg = flogging.MustFileLogger("debugUtil", "/Users/joker/go/src/github.com/hyperledger/fabric/logs/"+peerName+".log")
	// consoleLogger = flogging.MustGetLogger("consoleLogger")
}
func prepareDebugAble() {
	islog = true
	able := os.Getenv(constants.DEBUG_WARNING_LOGGABLE)
	if len(able) == 0 {
		islog = false
		return
	}
	b, e := strconv.ParseBool(able)
	if nil != e {
		b = false
	}
	islog = b
}

func InfoLogging(msg ...string) {
	if islog {
		sb := strings.Builder{}
		for _, s := range msg {
			sb.WriteString(decorate(s))
		}
		consoleInfo(nil, sb.String())
	}
}
func OneLineWarning(msg ...string) {
	sb := strings.Builder{}
	for _, s := range msg {
		sb.WriteString(decorate(s))
	}
	str := "当前协程:[%d]  %s"
	logplugin.Info(fmt.Sprintf(str, trace.CurGoroutineID(), msg))
	// consoleLogger.Warn(fmt.Sprintf(str, trace.CurGoroutineID(), msg))
}
func ErrorLogging(msg ...string) {
	sb := strings.Builder{}
	for _, s := range msg {
		sb.WriteString(decorate(s))
	}
	m := "\n" +
		"<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<\n" +
		" 当前协程:[%d]  %s  \n" +
		">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> "
	// consoleLogger.Error(fmt.Sprintf(m, trace.CurGoroutineID(), sb.String()))
	logplugin.Error(fmt.Sprintf(m, trace.CurGoroutineID(), sb.String()))
}
func ImportantLogging(msg ...string) {
	sb := strings.Builder{}
	s, ok := utils.FindCaller(1, defaultBlackList)
	if ok {
		sb.WriteString(s + "   ")
	}
	for _, s := range msg {
		sb.WriteString(decorate(s))
	}
	m := "\n" +
		"<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<\n" +
		" 当前协程:[%d]  %s  \n" +
		">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> "
	logplugin.Info(fmt.Sprintf(m, trace.CurGoroutineID(), sb.String()))
	// consoleLogger.Warn(fmt.Sprintf(m, trace.CurGoroutineID(), sb.String()))
	// warn(logg, sb.String())
}
func MsgLogging(msg ...string) {
	if islog {
		sb := strings.Builder{}
		for _, s := range msg {
			sb.WriteString(decorate(s))
		}
		// warn(logg, sb.String())
		consoleInfo(nil, sb.String())
	}
}
func MultiWarning(msg ...string) {
	sb := strings.Builder{}
	for _, s := range msg {
		sb.WriteString(decorate(s))
	}
	// warn(logg, sb.String())
	consoleInfo(nil, sb.String())
}

func FileWarning(msg ...string) {
	MultiWarning(msg...)
	// sb := strings.Builder{}
	// for _, s := range msg {
	// 	sb.WriteString(decorate(s))
	// }
	// warn(logg, sb.String())
}
func decorate(str string) string {
	return str + "\n"
}

func Warning(logger *flogging.FabricLogger, msg string) {
	// warn(logg, msg)
	consoleInfo(nil, msg)
}

func Warningf(logger *flogging.FabricLogger, template string, data ...interface{}) {
	sprintf := fmt.Sprintf(template, data...)
	if nil == logger {
		consoleInfo(nil, sprintf)
	} else {
		consoleInfo(logger, sprintf)
	}
}

func consoleInfo(loggg *flogging.FabricLogger, msg string) {
	if islog {
		m := "\n" +
			"<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<\n" +
			" 当前协程:[%d]  %s  \n" +
			">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> "
		// loggg.Info(fmt.Sprintf(m, trace.CurGoroutineID(), msg))
		logplugin.Info(fmt.Sprintf(m, trace.CurGoroutineID(), msg))
	}
}
func warn(loggg *zap.SugaredLogger, msg string) {
	// base64.StdEncoding.EncodeToString(sig)
	m := "\n" +
		"<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<\n" +
		" 当前协程:[%d]  %s  \n" +
		">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> "
	loggg.Warn(fmt.Sprintf(m, trace.CurGoroutineID(), msg))
}
