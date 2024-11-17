package panel

import (
	"cmsManage/utils/reqRequest"
	"errors"
	"fmt"
)

type Panel struct {
	Session *reqRequest.Request

	Website string // 站点地址

}

func NewPanel() *Panel {

	return &Panel{
		Session: reqRequest.NewRequest(),
	}
}

func (p *Panel) Login() error {

	jsonData := map[string]any{
		"name":          userName,
		"password":      password,
		"ignoreCaptcha": true,
		"captcha":       "",
		"captchaID":     "CNOMr9mDfqNhJjoHGbQX",
		"authMethod":    "session",
		"language":      "zh",
	}

	reqUrl := loginHost + "/api/v1/auth/login"
	response, err := p.Session.Post(reqRequest.RequestOption{
		Url:  reqUrl,
		Json: jsonData,
		Headers: map[string]string{
			"Accept":          "application/json, text/plain, */*",
			"Accept-Language": "zh",
			"Connection":      "keep-alive",
			"Content-Type":    "application/json",
			"EntranceCode":    "anNq",
			"Origin":          loginHost,
			"Referer":         loginHost + "/jsj",
			"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
		},
	})
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errors.New("登录失败, 状态码: " + fmt.Sprintf("%d", response.StatusCode))
	}
	return nil
}
