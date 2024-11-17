package zbolg

import (
	"cmsManage/utils/reqRequest"
	toPinyin "cmsManage/utils/to_pinyin"
	"errors"
	"github.com/duke-git/lancet/v2/strutil"
)

// AddCategory
//
//	@Description: 添加分类
//	@receiver z
//	@param categoryName 分类名
//	@param Intro 简介
//	@param addNavbar	是否添加到导航栏, 1:添加, 0:不添加
//	@return err
func (z *ZBolg) AddCategory(categoryName, Intro, addNavbar string) (err error) {

	if categoryName == "" || addNavbar == "" {
		err = errors.New("AddCategory params is empty")
		return
	}

	alias := toPinyin.ToPinyin(categoryName)

	var data map[string]string
	data = map[string]string{
		"ID":          "0",
		"Type":        "0",
		"Name":        categoryName,
		"Alias":       alias,
		"Order":       "0",
		"ParentID":    "0",
		"Template":    "index",
		"edtTemplate": "index",
		"LogTemplate": "single",
		"Intro":       Intro,
		"AddNavbar":   addNavbar,
	}

	var params map[string]string
	params = map[string]string{
		"act":       "CategoryPst",
		"csrfToken": z.csrfToken,
	}

	reqUrl := strutil.Before(z.LoginUrl, z.WebSite) + z.WebSite + "/zb_system/cmd.php"
	resp, err := z.Session.Post(reqRequest.RequestOption{
		Url:    reqUrl,
		Params: params,
		Data:   data,
	})
	if err != nil {
		return
	}

	status := resp.StatusCode
	if status != 200 {
		err = errors.New(z.WebSite + " AddCategory error status code is not 200")
		return
	}
	return
}
