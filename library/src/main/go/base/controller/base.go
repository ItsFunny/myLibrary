package controllers

import (
	"errors"
	"github.com/akkagao/citizens/base"
	"github.com/akkagao/citizens/common"
	"github.com/akkagao/citizens/common/log"
	"github.com/akkagao/citizens/utils"
	"github.com/akkagao/citizens/webbase/impl"
	"myLibrary/library/src/main/go/base/constants"
	"myLibrary/library/src/main/go/base/services"
	utils2 "myLibrary/library/src/main/go/utils"
	"net/http"
	"runtime"
	"strings"

	"github.com/astaxie/beego"
)

// 暂时全局变量,多系统考虑通过请求的session 或 其他方式获取
type BaseControllerInit struct {
	ReqID string
	ReqIP string
	Log   *log.Log
}

// BaseController : 基础 controller, 提供基础方法
type BaseController struct {
	beego.Controller
	BaseControllerInit
}

// Init : 处理完路由后调用
func (receiver *BaseController) Prepare() {
	ct := receiver.Ctx
	reqId := utils.GetReqID(ct.Request.Context())
	addr := ""
	{
		if ct.Request.Header.Get("UserIP") != "" {
			addr = ct.Request.Header.Get("UserIP")
		} else if ct.Request.Header.Get("X-Real-IP") != "" {
			addr = ct.Request.Header.Get("X-Real-IP")
		} else {
			addr = ct.Request.RemoteAddr
		}
	}

	// addr := ct.Request.RemoteAddr
	reqIP := addr
	if i := strings.Index(addr, ":"); i >= 0 {
		reqIP = addr[:i]
	}

	receiver.ReqID = reqId
	receiver.ReqIP = reqIP
	receiver.Log = log.NewLog(log.InitLog{ReqID: reqId})
}

// BeforeStart :方法开始调用
func (receiver *BaseController) BeforeStart(methodName string) {
	receiver.Log.SetPrefix(methodName)
}

// AfterEnd :方法结束调用, 开始的时候用 defer 调用
func (receiver *BaseController) AfterEnd() {
	if err := recover(); err != nil {
		pc, _, lineNO, ok := runtime.Caller(1)

		if ok {
			receiver.Log.Error("结束方法时, (%s:%d)出现panic:%s", runtime.FuncForPC(pc).Name(), lineNO, err)
		} else {
			receiver.Log.Error("结束方法时,出现panic:%s", err)
		}

		receiver.returnError(utils.NewSysErr("系统错误"))
	}
}


func (receiver *BaseController) SignBeforeStart(methodName string) {
	utils2.DefaultDebugDecorateShowSignal(methodName)
	receiver.BeforeStart(methodName)
}
func (receiver *BaseController) SignAfterEnd() {
	receiver.AfterEnd()
	utils2.DefaultDebugDecorateShowSignal(receiver.Log.GetPrefix())
}

func (receiver *BaseController)ReturnFail(obj services.IBaseRepsonseService,msg string)error{
	receiver.Ctx.Output.Status = http.StatusBadRequest
	obj.SetResponseCode(constants.FAIL)
	obj.SetResponseMsg(msg)
	receiver.Data["json"] = obj
	receiver.ServeJSON()
	return nil
}


// 返回 200
func (receiver *BaseController) returnSuccess(resp services.IBaseRepsonseService) (err error) {
	receiver.Ctx.Output.Status = http.StatusOK
	resp.SetResponseMsg("Success")
	resp.SetResponseCode(constants.SUCCESS)
	receiver.Data["json"] = resp
	receiver.ServeJSON()

	return
}

// 返回 400
// FIXME
func (receiver *BaseController) returnError(e interface{}) (err error) {
	receiver.Ctx.Output.Status = http.StatusBadRequest

	switch e.(type) {
	case common.FabricError:
		receiver.Data["json"] = "区块链调用发生错误"
	case common.SystemError:
		receiver.Data["json"] = "系统错误"
	case common.BussError:
		receiver.Data["json"] = "业务错误"
	case error:
		receiver.Data["json"] = e.(error).Error()
	default:
		receiver.Data["json"] = e
	}

	receiver.ServeJSON()

	return
}

func (receiver *BaseController) returnSysError() (err error) {
	return errors.New("系统错误")
}

// 返回详细的参数错误
func (receiver *BaseController) returnDetailParamError(format string, a ...interface{}) (err error) {

	return receiver.returnError(utils.NewApiDetailParamErr(format, a...))
}

// 返回 参数错误
func (receiver *BaseController) returnParamError() (err error) {
	return receiver.returnError(utils.ParamErr)
}

// 返回 401
func (receiver *BaseController) returnUnauthorized() (err error) {
	receiver.Ctx.Output.Status = http.StatusUnauthorized
	resp := map[string]string{"resultCode": "401", "resultMsg": "未授权请求"}
	receiver.Data["json"] = resp
	receiver.ServeJSON()

	return
}

// BaseControllerInit 转换为 BaseServicesInit
func (receiver *BaseController) GetServiceInit() base.IBaseServiceInit {
	init := new(webImpl.WebBaseServiceInitImpl)
	init.SetLogger(receiver.Log)
	init.SetReqId(receiver.ReqID)
	init.SetFabricSetup(receiver.BlockChainSetup)

	return init
}
