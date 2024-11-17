package panel

import (
	"cmsManage/utils/reqRequest"
	"errors"
	"fmt"
	"strings"
)

// ZBolg
//
//	@Description: zBolg 的伪静态设置
//	@receiver p
//	@return err
func (p *Panel) ZBolg() (err error) {
	iD, exist, err := p.searchWebsite()
	if err != nil {
		return
	}

	if !exist {
		err = fmt.Errorf("站点 %s 不存在", p.Website)
		return
	}

	reqUrl := loginHost + "/api/v1/websites/rewrite/update"

	jsonData := map[string]interface{}{
		"websiteID": iD,
		"content":   "if (-f $request_filename/index.html){\n\trewrite (.*) $1/index.html break;\n}\nif (-f $request_filename/index.php){\n\trewrite (.*) $1/index.php;\n}\nif (!-f $request_filename){\n\trewrite (.*) /index.php;\n}",
		"name":      "zblog",
	}

	response, err := p.Session.Post(reqRequest.RequestOption{
		Url:  reqUrl,
		Json: jsonData,
	})
	if err != nil {
		err = p.delWebsite()
		if err != nil {
			err = errors.New(p.Website + "伪静态设置失败, 删除网站失败" + err.Error())
		} else {
			err = errors.New(p.Website + "伪静态设置失败")
		}
		return
	}

	sourceHtml := response.SourceHtml
	if !strings.Contains(sourceHtml, `"code":200`) { // 200 为成功
		err = p.delWebsite()
		if err != nil {
			err = errors.New(p.Website + "伪静态设置失败, 删除网站失败" + err.Error())
		} else {
			err = errors.New(p.Website + "伪静态设置失败")
		}
		return
	}
	return
}
