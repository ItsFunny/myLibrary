package baseImpl

import (
	"myLibrary/library/src/main/go/base/services"
	"myLibrary/library/src/main/go/common/log"
	"runtime"
	"strings"
)

// BaseService : 所有 service 都包含, 提供一些基础通用的方法
type BaseServiceImpl struct {
	MethodName     string
	BaseInitConifg services.IBaseServiceInit
}

// BeforeStart :方法开始调用
func (receiver *BaseServiceImpl) BeforeStart(method string) {
	receiver.MethodName = method
	methodName := receiver.BaseInitConifg.GetLogger().GetPrefix() + " -> " + method
	receiver.BaseInitConifg.GetLogger().SetPrefix(methodName)
	receiver.BaseInitConifg.GetLogger().Info("开始调用:" + methodName)
}

func (receiver *BaseServiceImpl) UnMarshalErr(err error) error {
	receiver.BaseInitConifg.GetLogger().Error("["+receiver.MethodName+"] json unmarshal occur err:%v", err)
	return err
}

func (receiver *BaseServiceImpl) MarshalErr(err error) error {
	receiver.BaseInitConifg.GetLogger().Error("["+receiver.MethodName+"] json unmarshal occur err:%v", err)
	return err
}

// AfterEnd :方法结束调用, 开始的时候用 defer 调用
func (receiver *BaseServiceImpl) AfterEnd() {
	if err := recover(); err != nil {
		pc, _, lineNO, ok := runtime.Caller(1)

		if ok {
			receiver.BaseInitConifg.GetLogger().Error("结束方法时, (%s:%d)出现panic:%s", runtime.FuncForPC(pc).Name(), lineNO, err)
		} else {
			receiver.BaseInitConifg.GetLogger().Error("结束方法时,出现panic:%s", err)
		}
	}

	pre := receiver.BaseInitConifg.GetLogger().GetPrefix()
	receiver.BaseInitConifg.GetLogger().SetPrefix(strings.TrimRight(pre, " -> "+receiver.MethodName))
	receiver.BaseInitConifg.GetLogger().Info("结束对: { " + receiver.MethodName + " } 方法的调用")
}
// 同时打印结果
func (receiver *BaseServiceImpl)AfterEndWithResp(resp interface{}){
	receiver.BaseInitConifg.GetLogger().Info("[%s]的结果为:{%v}",receiver.MethodName,resp)
	receiver.AfterEnd()
}

// 设置基础信息
func (receiver *BaseServiceImpl) SetInitInfo(init services.IBaseServiceInit) {
	initInfo := new(BaseServiceInitImpl)
	initInfo.ReqID = init.GetReqId()
	initInfo.Log = init.GetLogger()
	receiver.BaseInitConifg = initInfo
}

// 获取基础信息
func (receiver *BaseServiceImpl) GetInitInfo() services.IBaseServiceInit {
	return receiver.BaseInitConifg
}
func (this *BaseServiceImpl) GetLogger() *log.Log {
	return this.GetInitInfo().GetLogger()
}
