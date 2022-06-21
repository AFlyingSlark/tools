package model

// 定义所有常用类型枚举

// 性别
type GenderCode int

const (
	Unknown GenderCode = iota // 未知
	Man                       // 男
	Woman                     // 女
)

func (g GenderCode) String() string {
	switch g {
	case Unknown:
		return "unknown"
	case Man:
		return "man"
	case Woman:
		return "woman"
	}

	return ""
}

// 语言
type LanguageCode int

const (
	_  LanguageCode = iota // 未知
	EN                     // 英语
	TR                     // 土耳其语
	HI                     // 印地语
	ZH                     // 中文
)

var (
	languageCodeStrings = map[LanguageCode]string{
		EN: "en",
		TR: "tr",
		HI: "hi",
		ZH: "zh",
	}
)

func (l LanguageCode) String() string {
	return languageCodeStrings[l]
}

// 成功失败状态
type StatusCode int

const (
	None StatusCode = iota // 未知
	OK                     // 成功
	Fail                   // 失败
)

func (s StatusCode) String() string {
	switch s {
	case OK:
		return "ok"
	case Fail:
		return "fail"

	default:
		return "none"
	}
}

// 开启关闭状态
type OpenStatusCode int

const (
	_     OpenStatusCode = iota
	Open                 // 开启
	Close                // 关闭
)

// 启用停用状态
type StartStatusCode int

const (
	_     StartStatusCode = iota
	Start                 // 启用
	Stop                  // 停用
)

// 订单状态
type OrderStatusCode int

const (
	OrderWait    OrderStatusCode = iota // 待处理
	OrderDoing                          // 处理中
	OrderSuccess                        // 处理成功
	OrderFail                           // 处理失败
	OrderReject                         // 处理驳回(拒绝)
)
