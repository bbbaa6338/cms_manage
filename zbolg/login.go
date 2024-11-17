package zbolg

import (
	"cmsManage/utils/reqRequest"
	"errors"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/strutil"
	"strings"
)

func (z *ZBolg) Login() (err error) {
	if err = z.loginStep1(); err != nil {
		return
	}

	if err = z.loginStep2(); err != nil {
		return
	}

	return
}

// loginStep1
//
//	@Description: 登录获取第一步的 csrfToken
//	@receiver z
//	@return err
func (z *ZBolg) loginStep1() (err error) {
	resp, err := z.Session.Get(reqRequest.RequestOption{
		Url: z.LoginUrl,
	})
	if err != nil {
		return
	}

	status := resp.StatusCode
	if status != 200 {
		err = errors.New("loginStep1 error, status code is not 200")
		return
	}

	sourceHtml := resp.SourceHtml
	if !strings.Contains(sourceHtml, `<input type="hidden" name="csrfToken" value="`) {
		err = errors.New("loginStep1 error, csrfToken not found")
		return
	}

	csrfToken := strings.Split(strings.Split(sourceHtml, `<input type="hidden" name="csrfToken" value="`)[1], `"`)[0]
	z.csrfToken = csrfToken
	return
}

// loginStep2
//
//	@Description: 登录获取第二步的 csrfToken 并保存
//	@receiver z
//	@return err
func (z *ZBolg) loginStep2() (err error) {
	verityUrl := strutil.Before(z.LoginUrl, z.WebSite) + z.WebSite + "/zb_system/cmd.php?act=verify"

	md5Str := cryptor.Md5String(zBolgPassword)
	data := map[string]string{
		"csrfToken":   z.csrfToken,
		"edtUserName": "",
		"edtPassWord": "",
		"btnPost":     "登录",
		"username":    zBolgUserName,
		"password":    md5Str,
		"savedate":    "1",
	}

	resp, err := z.Session.Post(reqRequest.RequestOption{
		Url:  verityUrl,
		Data: data,
	})
	if err != nil {
		return err
	}

	status := resp.StatusCode
	if status != 200 {
		err = errors.New("loginStep2 error, status code is not 200")
		return
	}

	sourceHtml := resp.SourceHtml
	if !strings.Contains(sourceHtml, `name="csrfToken" content="`) {
		err = errors.New("loginStep2 error, csrfToken not found")
		return
	}

	z.csrfToken = strings.Split(strings.Split(sourceHtml, `name="csrfToken" content="`)[1], `"`)[0]
	return
}

func (z *ZBolg) setAppCentre() (err error) {
	appCentreUrl := strutil.Before(z.LoginUrl, z.WebSite) + z.WebSite + "/plugin/AppCentre/setting.php?act=save"
	resp, err := z.Session.Post(reqRequest.RequestOption{
		Url: appCentreUrl,
		Data: map[string]string{
			"token":                z.csrfToken,
			"app_enabledcheck":     "0",
			"app_checkbeta":        "0",
			"app_enabledevelop":    "0",
			"app_enablegzipapp":    "0",
			"app_enablepluginsort": "0",
			"app_forcehttps":       "",
		},
	})
	if err != nil {
		return
	}

	status := resp.StatusCode
	if status != 200 {
		err = errors.New(z.WebSite + "在设置 AppCentre 失败, 状态码不对")
		return
	}
	return
}
