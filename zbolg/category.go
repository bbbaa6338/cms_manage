package zbolg

import (
	"cmsManage/utils/reqRequest"
	toPinyin "cmsManage/utils/to_pinyin"
	"errors"
	"github.com/duke-git/lancet/v2/slice"
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
	categoryName = strutil.Trim(categoryName)
	Intro = strutil.Trim(Intro)

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

// GetCategory
//
//	@Description: 获取所有的分类
//	@receiver z
//	@return allCategory
//	@return err
func (z *ZBolg) GetCategory() (allCategory []Category, err error) {
	// 获取分类
	var category []Category
	err = z.DB.Find(&category).Error
	if err != nil {
		return
	}
	allCategory = category
	return
}

// GetCategoryUrl
//
//	@Description: 获取分类的 URL 类型,返回分类的信息，以便后期进行拼接
//	@receiver z
//	@param categories
//	@return ids
//	@return ca[]Category
//	@return err
func (z *ZBolg) GetCategoryUrl(categories []Category) (ids []int, ca []Category, err error) {

	count := len(categories)
	r := 0 // 拼接链接的数量
	if count >= 5 {
		r = 5
	} else {
		r = count
	}

	s := 0
	casesID := make([]int, r)
	cases := make([]Category, r)
	for {
		if s == r {
			break
		}
		ca, _ := slice.Random(categories)
		iD := ca.CateID
		if slice.Contain(casesID, iD) {
			continue
		} else {
			//casesID = append(casesID, iD)
			//cases = append(cases, ca)
			casesID[s] = iD
			cases[s] = ca
			s++
		}
	}

	ids = casesID
	ca = cases
	return
}

// SerNavs
//
//	@Description: 设置导航栏
//	@receiver z
//	@param categories
//	@return err
func (z *ZBolg) SerNavs(categories []Category) (err error) {

	// 已有的导航栏
	var menu ZPBModule
	err = z.DB.Where("mod_FileName = ?", "navbar").First(&menu).Error
	if err != nil {
		return
	}

	// 获取分类的 URL 类型
	var config ZbpConfig
	err = z.DB.Where("conf_Key = ?", "ZC_CATEGORY_REGEX").First(&config).Error
	if err != nil {
		return
	}

	if config.ConfValue == "" {
		err = errors.New("ZC_CATEGORY_REGEX is empty")
		return
	}

	configPre := strutil.Before(config.ConfValue, "/") // 获取 / 前面的部分
	configPre = strutil.After(configPre, `}`)

	configAfter := strutil.AfterLast(config.ConfValue, "}") // 获取 / 后面的部分
	configAfter = strutil.Before(configAfter, `"`)

	li := `<li id="nvabar-item-index"><a href="{#ZC_BLOG_HOST#}">首页</a></li>`
	for _, category := range categories {
		li += `<li id="nvabar-item-index"><a href="{#ZC_BLOG_HOST#}` + configPre + "/" + category.CateAlias + configAfter + `" title="` + category.CateName + `">` + category.CateName + `</a></li>`
	}

	menu.ModContent = li
	err = z.DB.Save(&menu).Error
	if err != nil {
		return
	}
	return
}

// SetSelfLink
//
//	@Description: 设置自身链接为友链
//	@receiver z
//	@return err
func (z *ZBolg) SetSelfLink() (err error) {
	// 已有的友链
	var link ZPBModule
	err = z.DB.Where("mod_FileName = ?", "link").First(&link).Error
	if err != nil {
		return
	}

	// <li><a href="https://github.com/zblogcn" target="_blank" title="Z-Blog on Github">Z-Blog on Github</a></li>
	st := strutil.Before(z.LoginUrl, z.WebSite) + z.WebSite
	link.ModContent = `<li><a href="` + st + `" target="_blank" title="` + z.Title + `">` + z.Title + `</a></li>`
	err = z.DB.Save(&link).Error

	return
}
