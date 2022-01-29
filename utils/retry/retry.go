package retry

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
)

/* QuickRetry  重试
参数:
*	fun            	func() error      重试的函数
*	retryExtraCheck	func(error) bool  重试额外处理检查
*	retryExtra     	func()            额外处理
*	timeout        	time.Duration     重试最长时间
*	retryInterval  	time.Duration     重试间隔
返回值:
*	error	error
*/
func QuickRetry(fun func() error, retryExtraCheck func(error) bool, retryExtra func(), timeout, retryInterval time.Duration) error {
	times := 0
	start := time.Now()

	var err error

	for {
		times++

		if time.Since(start).Seconds() >= timeout.Seconds() {
			return fmt.Errorf("重试超过%s,依然失败,重试次数%d", timeout.String(), times)
		}

		if err = fun(); err == nil {
			return nil
		}

		if retryExtraCheck(err) {
			retryExtra()
		}

		time.Sleep(retryInterval)
	}
}

/* Retry 重试
请求参数：
* timeout   time.Duration 超时时间
* interval  time.Duration 重试等待时间
* logger    *slog.Logger   日志器
* function  func() error  执行的操作
返回值：
* error 异常
*/
func Retry(timeout, interval time.Duration, logger *zap.Logger, function func() error) error {
	return RetryGeneric(NewTimeoutJudge(timeout), interval, logger, function)
}

/* RetryAllTheTime 重试失败，没有时间上限
参数:
*	interval	time.Duration 重试间隔
*	logger  	*slog.Logger  日志器
*	function	func() error  执行的操作
返回值:
*	error	error
*/
func RetryAllTheTime(interval time.Duration, logger *zap.Logger, function func() error) error {
	return RetryGeneric(nil, interval, logger, function)
}

/* RetryGeneric 重试失败,自定义的判断
参数:
*	judge   	RetryJudge    自定义判断,返回true表示结束
*	interval	time.Duration 重试间隔
*	logger  	*slog.Logger  日志器
*	function	func() error  执行的操作
返回值:
*	error	error
*/
func RetryGeneric(judge RetryJudge, interval time.Duration, logger *zap.Logger, function func() error) error {
	var errMsg []string

	judge = NewNeverFinishJudge(judge)

	i := 0
	for ; !judge.Finished(); time.Sleep(interval) { // 失败,重试间隔时间
		if i != 0 {
			logger.Debug("重试")
		}
		// 需要被执行的方法
		err := function()
		if err == nil {
			// 没有错误直接返回
			return nil
		}

		logger.Error("错误", zap.String("错误信息", err.Error()))

		errMsg = append(errMsg, err.Error())
		i++
	}

	return errors.New(strings.Join(errMsg, ","))
}

/* ------------------------------ 下面是对RetryJudge的定义  ------------------------------*/

// RetryJudge Retry是否已经结束
type RetryJudge interface {
	// Finished 返回true，表示已经结束
	Finished() bool
}

// TimeoutJudge 基于总时长的RetryJudge
type TimeoutJudge struct {
	deadLine time.Time
}

/* NewTimeoutJudge 新建一个基于总时长的RetryJudge
参数:
*	timeout	time.Duration   总时长
返回值:
*	RetryJudge
*/
func NewTimeoutJudge(timeout time.Duration) RetryJudge {
	return &TimeoutJudge{deadLine: time.Now().Add(timeout)}
}

/* Finished 是否已经结束
参数:
返回值:
*	bool	bool
*/
func (t TimeoutJudge) Finished() bool {
	return time.Now().After(t.deadLine)
}

// TimesJudge 基于重试次数的RetryJudge
type TimesJudge struct {
	leftTimes uint
}

/* NewTimesJudge 新建一个TimesJudge
参数:
*	times	uint
返回值:
*	RetryJudge
*/
func NewTimesJudge(times uint) RetryJudge {
	return &TimesJudge{leftTimes: times}
}

/* Finished 是否已经结束
参数:
返回值:
*	bool	bool
*/
func (l *TimesJudge) Finished() bool {
	finished := l.leftTimes == 0

	if !finished {
		l.leftTimes--
	}

	return finished
}

/*  ------------------------------ 这里开始是一些复合的RetryJudge  ------------------------------*/

// NeverFinishJudge 永不结束的RetryJudge,如果内嵌的Judge存在，那么以内嵌为主，否则永不结束
type NeverFinishJudge struct {
	RetryJudge // 内嵌的Judge
}

/* NewNeverFinishJudge 新建一个NeverFinishJudge
参数:
*	retryJudge	RetryJudge
返回值:
*	RetryJudge	RetryJudge
*/
func NewNeverFinishJudge(retryJudge RetryJudge) RetryJudge {
	return &NeverFinishJudge{RetryJudge: retryJudge}
}

/* Finished 是否已经结束
参数:
返回值:
*	bool	bool
*/
func (w NeverFinishJudge) Finished() bool {
	if w.RetryJudge != nil {
		return w.RetryJudge.Finished()
	}

	return false
}

// LogicOperator 逻辑操作符
type LogicOperator int

const (
	// LogicAnd 逻辑与,全部都是true,才返回true
	LogicAnd LogicOperator = iota + 1
	// LogicOr 逻辑或,任意一个是true,就返回true
	LogicOr
)

// LogicCompositeJudge 通过逻辑操作将多个RetryJudge组合
type LogicCompositeJudge struct {
	operator LogicOperator // 逻辑操作符
	judges   []RetryJudge  // 判断
}

/* NewLogicCompositeJudge 新建一个LogicCompositeJudge
参数:
*	operator	LogicOperator 逻辑操作符
*	judges  	[]RetryJudge  判断，任意一个元素都不能为nil
返回值:
*	RetryJudge
*/
func NewLogicCompositeJudge(operator LogicOperator, judges []RetryJudge) RetryJudge {
	return &LogicCompositeJudge{operator: operator, judges: judges}
}

/* Finished 是否已经结束
参数:
返回值:
*	bool	bool
*/
func (l LogicCompositeJudge) Finished() bool {
	switch l.operator {
	case LogicAnd:
		finished := true
		for i := range l.judges {
			finished = l.judges[i].Finished() && finished
		}

		return finished
	case LogicOr:
		finished := false
		for i := range l.judges {
			finished = finished || l.judges[i].Finished()
		}

		return finished
	default:
		return false
	}
}
