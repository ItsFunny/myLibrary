# JOB 任务分发
- consume->execute->handle
- 角色介绍: 
    -   executor: 
        -   任务的第一级执行者
        -   基础结构体为: BaseJobExecutor
    -   handler: 
        -   任务的二级执行者
        -   基础base结构体为: BaseJobHandler
# 使用方法
- 具体看最后吧,这里没写好
* 比如说有validatePacket的业务
    * 先创建对应的executor:
    ```
    type PacketValidateJobExecutor struct {
    	*job.BaseJobExecutor
    	JobValidate job.IJobHandler
    }
    内部只需要继承BaseJobExecutor,再添加job.IJobHandler即可
    这里添加IJobHandler的原因在于,handler是executor的具体执行者,根据自身业务自定义接口继承handler接口
  
    
    ```
    * 创建的executor实现job.IConcreteExecutor接口
    ```
    func (e *PacketValidateJobExecutor) DoExecute(job job.IAppEvent) (interface{}, bool) {
    	return e.JobValidate.Handle(job), false
    }
    
    ```
    *  创建业务的handler接口
    ```
    type IPacketValidateJob interface {
    	job.IJobConcreteHandler
    }
    内部可根据需求再添加额外的方法
    ```
    * 根据不同的策略实现具体的方法
    
    若handler也分层级,则内部也需要创建一个struct继承*BaseJobHandler
    ```
    func (this *TcpPacketValidation) DoHandle(job job.IAppEvent) interface{} {
    	panic("implement me")
    }
    
    func (this *UdpPacketValidation) DoHandle(job job.IAppEvent) interface{} {
    	panic("implement me")
    }
    ```
    * 最后,工厂方式创建:
        -   创建业务对应的executor工厂方法:
        ```
        func NewPacketValidateJobExecutor(handler job.IJobHandler) job.IJobExecutor {
        	executor := &PacketValidateJobExecutor{}
        	executor.BaseJobExecutor = new(job.BaseJobExecutor)
        	executor.Type = JOB_VALIDATE
        	executor.JobValidate = handler
        	return executor
        }
        ```
        - 创建业务对应的handler工厂方法:
        ```
        func NewValidateHandler() job.IJobHandler {
        	tcpHander := &BasePacketValidateJobHandler{}
        	tcpHander.Type = "tcp"
        	tcpHander.ConcreteHandler = new(TcpPacketValidation)
        
        	udpHandler := &BasePacketValidateJobHandler{}
        	udpHandler.Type = "udp"
        	udpHandler.ConcreteHandler = new(UdpPacketValidation)
        	tcpHander.NextHandler = udpHandler
        
        	return tcpHander
        }
        ```
        - 最后再统一固定格式:
        ```
            mediator := job.NewJobExecuteMediator(jobs.NewPacketValidateJobExecutor(jobs.NewValidateHandler()))
            res := mediator.Execute(job.NewAppEvent().SetData(packet).SetJobType(jobs.JOB_VALIDATE))
            NewJobExecuteMediator方法,在base中已提供
        ```
---

- 具体demo: 
- 先自定义业务接口: ISmsHandler
-   实现自定义的实现类: SmsWorker -> SmsExecutor-> SmsHandler,复写DoConsume->DoExecute->DoHandle方法
-   编写各自工厂函数 由上至下:NewSmsWorker->NewSMSExecutor->NewSMSHandler
-   test测试: 
```
func TestNewSmsWorker(t *testing.T) {
	dispatcher.InitDispatcher(10)
	RegisteSmsWorker()
	go dispatcher.Run()
	event := job.NewAppEvent()
	event.SetData("123")
	event.SetJobType(SMS_EXECUTOR_JOB_TYPE, SMS_HANDLER_JOB_TYPE)
	dispatcher.AddJob(event)
	time.Sleep(time.Second * 1)
}
```
```

const (
	SMS_EXECUTOR_JOB_TYPE         = job.JobType("SMS")
	SMS_HANDLER_JOB_TYPE = job.JobType("SMS_HANDLER")
)

func RegisteSmsWorker() {
	// 3个sms worker 处理发送短信的任务
	// dispatcher.Register(NewSmsWorker())
	// dispatcher.Register(NewSmsWorker())
	dispatcher.Register(NewSmsWorker())
}

type ISmsHandler interface {
	job.IJobConcreteHandler
}

type SmsWorker struct {
	*worker.BaseWoker
	SmsExecutor *SmsExecutor
}

type SmsExecutor struct {
	*job.BaseJobExecutor
	SmsHandler ISmsHandler
}

type SmsHandler struct {
	*job.BaseJobHandler
}

func NewSmsWorker() *SmsWorker {
	w := new(SmsWorker)
	w.BaseWoker = worker.NewBaseWorkerWithOutConfig(w, nil)
	w.SmsExecutor = NewSMSExecutor()
	w.InitJobQueue(1)
	return w
}

func NewSMSExecutor() *SmsExecutor {
	sms := new(SmsExecutor)
	sms.BaseJobExecutor = job.NewBaseJobExecutor(SMS_EXECUTOR_JOB_TYPE, sms, nil)
	sms.SmsHandler = NewSMSHandler()
	return sms
}

func NewSMSHandler() *SmsHandler {
	s := new(SmsHandler)
	s.BaseJobHandler = job.NewBaseJobHandler(SMS_HANDLER_JOB_TYPE, s, nil)
	return s
}

func (this *SmsWorker) DoConsume(d interface{}) (interface{}, error) {
	event := d.(job.IAppEvent)
	return this.SmsExecutor.Execute(event)
}

func (this *SmsExecutor) DoExecute(job job.IAppEvent) (interface{}, bool, error) {
	handle, e := this.SmsHandler.DoHandle(job)
	return handle, false, e
}

func (this *SmsHandler) DoHandle(job job.IAppEvent) (interface{}, error) {
	// 发送短信
	fmt.Println("收到了job:", job.GetData())
	return nil, nil
}


```