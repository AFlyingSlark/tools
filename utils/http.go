package utils

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// HTTPOperationArgument http操作参数接口
type HTTPOperationArgument interface {
	Logger() *zap.Logger
	Client() *http.Client
	Ctx() context.Context
	Name() string
	SetName(name string)
}

// httpOperationArgument  HTTP操作必要参数
type httpOperationArgument struct {
	name   string          // 名称
	ctx    context.Context // 上下文，主要用作超时控制
	client *http.Client    // http连接，主要是为了复用连接
	logger *zap.Logger     // 日志器
}

/*NewHttpOperationArgument 新建一个HTTP操作参数 不能为空
参数:
*	ctx           	context.Context         上下文
*	name          	string                  名称
*	client        	*http.Client            http客户端
*	logger        	*slog.Logger            日志器
返回值:
*	HTTPOperationArgument	HTTPOperationArgument
*/
func NewHttpOperationArgument(ctx context.Context, name string, client *http.Client, logger *zap.Logger) HTTPOperationArgument {
	return &httpOperationArgument{
		name:   name,
		ctx:    ctx,
		client: client,
		logger: logger,
	}
}

// SetName 设置名称，适用于资源不变，但是不同名称或者步骤
func (h *httpOperationArgument) SetName(name string) {
	h.name = name
}

// Name 名称
func (h httpOperationArgument) Name() string {
	return h.name
}

// Logger 日志器
func (h httpOperationArgument) Logger() *zap.Logger {
	return h.logger
}

// Client http客户端
func (h httpOperationArgument) Client() *http.Client {
	return h.client
}

// Ctx 上下文
func (h httpOperationArgument) Ctx() context.Context {
	return h.ctx
}

/*HttpGet 执行GET请求
参数:
* 	argument        HTTPOperationArgument     http资源
*	url           	string                    请求地址
*	headers       	map[string]string         首部
返回值:
*	data   	        []byte                    应答内容
*	statusCode	    int            			  应答code码
*	err 	        error                     应答错误
*/
func HttpGet(argument HTTPOperationArgument, url string, headers map[string]string) (data []byte, statusCode int, err error) {
	var resp *http.Response

	data, resp, err = HttpGeneric(argument, url, headers, http.MethodGet, nil)

	statusCode = resp.StatusCode

	return
}

/*HttpPOST 执行POST请求
参数:
*	argument        HTTPOperationArgument     请求信息
*	url           	string                    请求地址
*	headers       	map[string]string         首部
* 	body            io.Reader                 请求body reader
返回值:
*	data   	        []byte                    应答内容
*	statusCode	    int            			  应答code码
*	err 	        error                     应答错误
*/

func HttpPOST(argument HTTPOperationArgument, url string, headers map[string]string, body io.Reader) (data []byte, statusCode int, err error) {
	var resp *http.Response

	data, resp, err = HttpGeneric(argument, url, headers, http.MethodPost, body)
	if err != nil {
		return
	}

	statusCode = resp.StatusCode
	// cookies []*http.Cookie
	_ = resp.Cookies() // 如果需要cookies,请赋值

	return
}

/*HttpPOSTGeneric 执行请求,并且返回完整的*http.Response,只是其中的Body已经不能再读,而返回值data返回了Body的内容
参数:
*	argument        HTTPOperationArgument     请求信息
*	url           	string                    请求地址
*	headers       	map[string]string         首部
*	method  		string                	  HTTP方法(http.MethodPost)
* 	body            io.Reader                 请求body reader
返回值:
*	data   	        []byte                    应答内容
*	resp	        *http.Response            应答
*	err 	        error                     应答错误
*/
func HttpGeneric(argument HTTPOperationArgument, url string, headers map[string]string, method string, body io.Reader) (data []byte, resp *http.Response, err error) {
	var req *http.Request

	ctx := argument.Ctx()
	if ctx == nil {
		ctx = context.Background()
	}
	req, err = http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		err = errors.Wrap(err, "构建请求")
		return
	}

	// 设置头信息
	for k, v := range headers {
		if k == "" {
			continue
		}
		req.Header.Set(k, v)
		if k == "Host" {
			req.Host = v
		}
	}

	// 2.执行请求
	resp, err = argument.Client().Do(req)
	if err != nil {
		err = errors.Wrap(err, "执行请求")
		return
	}

	defer IgnoreError(argument.Logger(), "关闭body应答", resp.Body.Close)

	// 3.读取body数据
	if data, err = ioutil.ReadAll(resp.Body); err != nil {
		err = errors.Wrap(err, "读取resp应答body")
	}

	return
}

/*IgnoreError 忽略方法的错误 ,如果fun执行返回错误且logger!=nil,那么日志输出
参数:
*	logger	*zap.Logger   日志器
*	msg   	string        相关信息
*	fun   	func() error  方法
返回值:
*/
func IgnoreError(logger *zap.Logger, msg string, fun func() error) {
	err := fun()
	if err != nil && logger != nil {
		logger.Error(msg, zap.Error(err))
	}
}
