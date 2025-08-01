package xe

// 响应码
const (
	SuccessCode = 200
	FailureCode = 500
)

// Middleware 中间件构造函数
type Middleware func(TaskFunc) TaskFunc

func (e *Executor) chain(next TaskFunc) TaskFunc {
	for i := range e.middlewares {
		next = e.middlewares[len(e.middlewares)-1-i](next)
	}
	return next
}

// 通用响应
type Return[T any] struct {
	Code    int    `json:"code"`              // 200 表示正常、其他失败
	Msg     string `json:"msg omitempty"`     // 错误提示消息
	Content T      `json:"content omitempty"` // 响应内容
}

/*****************  上行参数  *********************/

// RegistryParam 注册参数
type RegistryParam struct {
	RegistryGroup string `json:"registry_group"`
	RegistryKey   string `json:"registry_key"`
	RegistryValue string `json:"registry_value"`
}

// 执行器执行完任务后，回调任务结果时使用
type CallbackParamList []HandleCallbackParam

type HandleCallbackParam struct {
	JobID      string `json:"job_id"`      // 任务ID
	ExecutorID string `json:"executor_id"` // 执行
	HandleCode int    `json:"handle_code"` //200表示正常,500表示失败
	HandleMsg  string `json:"handle_msg"`
}

/*****************  下行参数  *********************/

// 阻塞处理策略
const (
	serialExecution = "SERIAL_EXECUTION" //单机串行
	discardLater    = "DISCARD_LATER"    //丢弃后续调度
	coverEarly      = "COVER_EARLY"      //覆盖之前调度
)

// TriggerParam 触发任务请求参数
type TriggerParam struct {
	JobID                 string `json:"job_id"`                  // 任务ID
	ExecutorID            string `json:"executor_id"`             // 执行
	ExecutorHandler       string `json:"executor_handler"`        // 任务标识
	ExecutorParams        string `json:"executor_Params"`         // 任务参数
	ExecutorBlockStrategy string `json:"executor_block_strategy"` // 任务阻塞策略
	ExecutorTimeout       int64  `json:"executor_timeout"`        // 任务超时时间，单位秒，大于零时生效
	BroadcastIndex        int64  `json:"broadcast_index"`         // 分片参数：当前分片
	BroadcastTotal        int64  `json:"broadcast_total"`         // 分片参数：总分片
}

// 终止任务请求参数
type KillParam struct {
	JobID      string `json:"job_id"`      // 任务ID
	ExecutorID string `json:"executor_id"` // 执行
}

// 忙碌检测请求参数
type IdleBeatParam struct {
	JobID      string `json:"job_id"`      // 任务ID
	ExecutorID string `json:"executor_id"` // 执行
}

func newCallback(t TriggerParam, code int, msg string) HandleCallbackParam {
	return HandleCallbackParam{
		JobID:      t.JobID,
		ExecutorID: t.ExecutorID,
		HandleCode: code,
		HandleMsg:  msg,
	}
}

func Success(msg string) Return[string] {
	return Return[string]{
		Code: SuccessCode,
		Msg:  msg,
	}
}

func Failure(msg string) Return[string] {
	return Return[string]{
		Code: FailureCode,
		Msg:  msg,
	}
}
