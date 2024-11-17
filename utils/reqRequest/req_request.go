package reqRequest

import (
	"fmt"
	"github.com/imroc/req/v3"
	"math"
	"net/url"
	"strings"
	"time"
)

type Request struct {
	Session *req.Client // 全局的客户端实例
}

type RequestOption struct {
	Url     string
	Params  map[string]string
	Headers map[string]string
	Data    map[string]string
	Json    map[string]any

	Proxy   string
	TimeOut time.Duration // 超时时间，默认 8s

	RetryCount   int    // 重试次数，默认 0，不重试，如果大于0，那么可以与 RetryMessage 配合使用
	RetryMessage string // FIXME:  重试消息，如果返回的 html 中包含该消息，则进行重试

	RedirectPolicy bool // true 表示禁止重定向，默认为 false

}

type Response struct {
	SourceHtml string
	StatusCode int
	Url        *url.URL // 最后请求的 url, 直接使用 String() 方法转为字符串

	ReqResponse *req.Response // 原始的 req.Response
}

// NewRequest 创建一个 Request 实例并初始化客户端
func NewRequest() *Request {
	return &Request{
		Session: req.NewClient(),
	}
}

func (r *Request) initSession(reqOpt *RequestOption) {

	if r.Session == nil {
		r.Session = req.NewClient()
	}

	// 设置 Headers
	if reqOpt.Headers != nil {
		r.Session.SetCommonHeaders(reqOpt.Headers)
	} else {
		r.Session.SetCommonHeaders(map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
		})
	}

	// 设置 Content-Type
	if reqOpt.Data != nil {
		r.Session.SetCommonHeaders(map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		})
	} else if reqOpt.Json != nil {
		r.Session.SetCommonHeaders(map[string]string{
			"Content-Type": "application/json",
		})
	}

	// 设置代理
	if reqOpt.Proxy != "" {
		r.Session.SetProxyURL(reqOpt.Proxy)
	}

	// 设置超时时间
	if reqOpt.TimeOut > 0 {
		r.Session.SetTimeout(reqOpt.TimeOut)
	} else {
		r.Session.SetTimeout(8 * time.Second)
	}

	// 设置重定向策略
	if reqOpt.RedirectPolicy {
		r.Session.SetRedirectPolicy(req.NoRedirectPolicy())
	}
}

func (r *Request) prepareRequest(reqOpt RequestOption) *req.Request {
	r.initSession(&reqOpt)
	request := r.Session.R()

	if reqOpt.Params != nil {
		request.SetQueryParams(reqOpt.Params)
	}

	if reqOpt.RetryCount > 0 {
		request.SetRetryCount(reqOpt.RetryCount)
		request.SetRetryInterval(func(resp *req.Response, attempt int) time.Duration {
			sleep := 0.01 * math.Exp2(float64(attempt))
			return time.Duration(math.Min(2, sleep)) * time.Second
		})

		request.AddRetryCondition(func(resp *req.Response, err error) bool {
			if err != nil {
				fmt.Println("网络错误，将进行重试..." + err.Error())
				return true
			}
			if resp != nil && resp.StatusCode == 200 {
				body := resp.String()
				if strings.Contains(body, reqOpt.RetryMessage) {
					fmt.Println(fmt.Sprintf("响应中包含 %s 字符串，将进行重试...", reqOpt.RetryMessage))
					return true
				}
			}
			return false
		})
	}

	return request
}

func (r *Request) Get(reqOpt RequestOption) (response Response, err error) {
	request := r.prepareRequest(reqOpt)

	resp, err := request.Get(reqOpt.Url)
	if err != nil {
		return
	}

	response = Response{
		SourceHtml:  resp.String(),
		StatusCode:  resp.StatusCode,
		Url:         resp.Request.URL,
		ReqResponse: resp,
	}
	return
}

func (r *Request) Post(reqOpt RequestOption) (response Response, err error) {
	request := r.prepareRequest(reqOpt)

	if reqOpt.Json != nil {
		request.SetBody(reqOpt.Json)
	} else if reqOpt.Data != nil {
		request.SetFormData(reqOpt.Data)
	}

	resp, err := request.Post(reqOpt.Url)
	if err != nil {
		return
	}

	response = Response{
		SourceHtml:  resp.String(),
		StatusCode:  resp.StatusCode,
		Url:         resp.Request.URL,
		ReqResponse: resp,
	}
	return
}
