package zbolg

import (
	"cmsManage/utils/reqRequest"
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/datetime"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/duke-git/lancet/v2/strutil"
	"gorm.io/gorm"
	"strings"
)

var themes = map[string]string{
	//"aymeight": "zblog博客CMS主题aymeight",
	//// "default":      "默认主题",
	//"expolee":      "博览（繁花似锦）",
	//"Hipaper_vip":  "Hipaper报纸新闻[VIP]",
	//"hopelee":      "希望（心动之选）",
	//"Jz52_acolumn": "极致·单栏简约 Pro 隐藏式侧栏 SEO友好",
	//"Jz52_autumn":  "极致·立秋",
	//"Jz52_guopi":   "极致·果皮博客 清新资讯",
	//"Jz52_maxr":    "极致·大R 功能强大 圆角主题",
	//"Jz52_tsc":     "极致·三栏简约轻社区 朋友圈SNS风",
	//"jz_zpojie":    "极致·自适应万能主题 下载站主题 自带下载工具箱",
	//"koilee":       "好运锦鲤（时来运转）",
	//"maoc_grace":   "猫C优雅博客主题 高端大气 时尚美观",
	//"maoc_yaa":     "猫C清新博客主题",
	//"umModern":     "优美追梦版响应式主题",
	//"mxlee":    "梦想家（重塑经典）",
	"quietlee": "宁静致远（国潮主题）",
	"youthlee": "致青春（风华正茂）",
	//"zbox":         "博客主题zbox",
}

func (z *ZBolg) Theme() (err error) {

	themeKeys := make([]string, 0)
	for k := range themes {
		themeKeys = append(themeKeys, k)
	}
	z.theme, _ = slice.Random(themeKeys)
	//z.theme = "expolee"

	// 设置 TDK
	err = z.WebSetting(z.Title, z.SubTitle, z.themeData.CopyRight(z.WebSite))
	if err != nil {
		return
	}

	category, err := z.GetCategory()
	if err != nil {
		return
	}
	z.allCategory = category

	categoryIDs, categoryInfo, err := z.GetCategoryUrl(category)
	if err != nil {
		return
	}

	z.categoryInfo = categoryInfo
	z.categoryIDs = categoryIDs

	// 获取文章
	articleIds, err := z.GetArticleIds()
	if err != nil {
		return
	}
	z.articleIds = articleIds

	err = z.checkTheme()
	if err != nil {
		return
	}

	err = z.startTheme()
	if err != nil {
		return
	}

	// 构造主题的数据
	//themeData := fake.WebFake{}
	// 设置主题
	switch z.theme {
	case "Hipaper_vip":
		err = z.setHipaperVip()
	case "aymeight":
		err = z.setAymeight()
	case "expolee":
		err = z.setExpolee()
	case "hopelee":
		err = z.setHopolee()
	case "Jz52_acolumn":
		err = z.setJz52Acolumn()
	case "Jz52_autumn":
		err = z.setJz52Autumn()
	case "Jz52_guopi":
		err = z.setJz52Guopi()
	case "Jz52_maxr":
		err = z.setJz52Maxr()
	case "Jz52_tsc":
		err = z.setJz52Tsc()
	case "jz_zpojie":
		err = z.setJzZpojie()
	case "koilee":
		err = z.setKoilee()
	case "maoc_grace":
		err = z.setMaocGrace()
	case "maoc_yaa":
		err = z.setMaocYaa()
	case "umModern":
		err = z.setUmModern()
	case "zbox":
		err = z.setZbox()
	case "mxlee":
		err = z.setMxlee()
	case "quietlee":
		err = z.setQuietlee()
	case "youthlee":
		err = z.setYouthlee()
	default:

		err = errors.New(z.WebSite + "主题不存在: " + z.theme)

	}
	if err != nil {
		return
	}

	// 没有设置导航的，需要设置导航
	err = z.SerNavs(z.categoryInfo)
	if err != nil {
		return
	}
	err = z.SetSelfLink()

	return
}

// checkTheme
//
//	@Description: 检查主题是否存在
//	@receiver z
//	@return err
func (z *ZBolg) checkTheme() (err error) {
	reqUrl := strutil.Before(z.LoginUrl, z.WebSite) + z.WebSite + "/zb_system/admin/index.php?act=ThemeMng"
	response, err := z.Session.Get(reqRequest.RequestOption{
		Url: reqUrl,
	})
	if err != nil {
		return
	}

	if response.StatusCode != 200 {
		err = errors.New(z.WebSite + "获取主题页面失败: ")
		return
	}

	rt := response.SourceHtml
	if strings.Contains(rt, `data-themeid="`+z.theme+`"`) {
		return
	} else {
		err = errors.New(z.WebSite + "主题不存在: " + z.theme)
		return
	}
}

// startTheme
//
//	@Description: 启动主题
//	@receiver z
//	@return err
func (z *ZBolg) startTheme() (err error) {
	reqUrl := strutil.Before(z.LoginUrl, z.WebSite) + z.WebSite + "/zb_system/cmd.php"

	style := "style"
	if z.theme == "Hipaper_vip" {
		style = "default"
	} else if z.theme == "umModern" {
		style = "um"
	}

	response, err := z.Session.Post(reqRequest.RequestOption{
		Url: reqUrl,
		Data: map[string]string{
			"csrfToken": z.csrfToken,
			"theme":     z.theme,
			"style":     style,
		},
		Params: map[string]string{
			"act": "ThemeSet",
		},
	})
	if err != nil {
		return
	}

	if response.StatusCode != 200 {
		err = errors.New(z.WebSite + "主题设置失败")
		return
	}

	return
}

func (z *ZBolg) setHipaperVip() (err error) {
	// 设置主题的数据 Hipaper_vip,footmsg

	var configs = map[string]string{
		"topnavset":   `s:1:"0";`,
		"icset":       `s:1:"0";`,
		"imshow":      `s:1:"0";`,
		"imshow_mb":   `s:1:"0";`,
		"ocmt":        `s:1:"0";`,
		"seotool":     `s:1:"0";`,
		"backtop":     `s:1:"1";`,
		"description": fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
		"keywords":    fmt.Sprintf(`s:%d:"%s";`, len(z.Keywords), z.Keywords),
		"footmsg":     `s:1:"1";`,
		"setsave":     `s:1:"0";`,
	}

	for s, s2 := range configs {
		// 更新数据
		err = z.DB.Model(&ZbpConfig{}).Where("conf_Name = ? AND conf_Key = ?", "Hipaper_vip", s).Update("conf_Value", s2).Error
		if err != nil {
			return
		}
	}

	return

}

func (z *ZBolg) setAymeight() (err error) {
	// 设置主题的数据 aymeight

	welcome := "欢迎访问 " + z.Title + ""
	hmcats := `i:0;s:5:"false";`
	a := 1
	for i, d := range z.allCategory {
		i += 1
		Id := fmt.Sprintf("%d", d.CateID)
		//i:1;s:1:"1";
		hmcats += fmt.Sprintf(`i:%d;s:%d:"%s";`, i, len(Id), Id)
		a += 1
	}
	// a:11:{i:0;s:5:"false";i:1;s:1:"1";i:2;s:1:"2";i:3;s:1:"3";i:4;s:1:"4";i:5;s:1:"5";i:6;s:1:"6";i:7;s:1:"7";i:8;s:1:"8";i:9;s:1:"9";i:10;s:2:"10";}
	hmcats = fmt.Sprintf(`a:%d:{%s}`, a, hmcats)
	var configs = map[string]string{
		"welcome":     fmt.Sprintf(`s:%d:"%s";`, len(welcome), welcome),
		"icp":         fmt.Sprintf(`s:%d:"%s";`, len(z.themeData.ICP()), z.themeData.ICP()),
		"beian":       `s:0:"";`,
		"seoON":       `s:1:"1";`,
		"hmcats":      hmcats,
		"weibo":       fmt.Sprintf(`s:%d:"%s";`, len(z.themeData.Weibo()), z.themeData.Weibo()),
		"qq":          fmt.Sprintf(`s:%d:"%s";`, len(z.themeData.QQ()), z.themeData.QQ()),
		"title":       fmt.Sprintf(`s:%d:"%s";`, len(z.Title), z.Title),
		"keywords":    fmt.Sprintf(`s:%d:"%s";`, len(z.Keywords), z.Keywords),
		"description": fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
	}

	for s, s2 := range configs {
		// 更新数据
		err = z.DB.Model(&ZbpConfig{}).Where("conf_Name = ? AND conf_Key = ?", "aymeight", s).Update("conf_Value", s2).Error
		if err != nil {
			return
		}
	}

	return
}

func (z *ZBolg) setExpolee() (err error) {
	// 设置主题的数据 expolee

	ftwenzi := z.themeData.CopyRight(z.WebSite)
	footabout := z.themeData.CopyRight(z.WebSite)
	indextyle, _ := slice.Random([]string{"1", "2", "3"})
	postime, _ := slice.Random([]string{"1", "2", "3"})
	footpage, _ := slice.Random([]string{"1", "2", "3", "4"})
	ii, _ := slice.Random([]int{6, 9, 12})
	blognumSlice := z.RandomArticleIds(z.articleIds, ii)
	blognum := strings.Join(blognumSlice, ",")

	iii, _ := slice.Random([]int{8, 9, 10})
	cltjwzidSlice := z.RandomArticleIds(z.articleIds, iii)
	cltjwzid := strings.Join(cltjwzidSlice, ",")

	catstyle, _ := slice.Random([]string{"1", "2"})
	nocatitle, _ := slice.Random([]string{"1", "0"})

	var configs = map[string]string{
		"flashlights":  `s:1:"1";`,
		"webtitle":     fmt.Sprintf(`s:%d:"%s";`, len(z.Title), z.Title),
		"websubtitle":  fmt.Sprintf(`s:%d:"%s";`, len(z.SubTitle), z.SubTitle),
		"Keywords":     fmt.Sprintf(`s:%d:"%s";`, len(z.Keywords), z.Keywords),
		"Description":  fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
		"qqkfoff":      `s:1:"0";`,
		"toploginoff":  `s:1:"0";`,
		"qqloginoff":   `s:1:"0";`,
		"topnavoff":    `s:1:"0";`,
		"tabfloff":     `s:1:"0";`,
		"icpbeian":     `s:0:"";`,
		"gabbeian":     `s:0:"";`,
		"hotcoff":      `s:1:"0";`,
		"cmslistoff":   `s:1:"0";`,
		"goodsoff":     `s:1:"0";`,
		"textcoff":     `s:1:"0";`,
		"importantoff": `s:1:"0";`,
		"zdywzseo":     `s:1:"1";`,
		"fjzhon":       `s:1:"1";`,
		"nightoff":     `s:1:"1";`,
		"msideoff":     `s:1:"1";`,
		"runtimeoff":   `s:1:"1";`,
		"indextyle":    fmt.Sprintf(`s:%d:"%s";`, len(indextyle), indextyle),
		"ftwenzi":      fmt.Sprintf(`s:%d:"%s";`, len(ftwenzi), ftwenzi),
		"blognum":      fmt.Sprintf(`s:%d:"%s";`, len(blognum), blognum),
		"catstyle":     fmt.Sprintf(`s:%d:"%s";`, len(catstyle), catstyle),
		"nocatitle":    fmt.Sprintf(`s:%d:"%s";`, len(nocatitle), nocatitle),
		"footpage":     fmt.Sprintf(`s:%d:"%s";`, len(footpage), footpage),
		"postime":      fmt.Sprintf(`s:%d:"%s";`, len(postime), postime),
		"cltjwzid":     fmt.Sprintf(`s:%d:"%s";`, len(cltjwzid), cltjwzid),
		"footabout":    fmt.Sprintf(`s:%d:"%s";`, len(footabout), footabout),
	}

	for s, s2 := range configs {
		// 获取数据，没有数据则插入，有数据则更新
		var config ZbpConfig

		err = z.DB.Where("conf_Name = ? AND conf_Key = ?", "expolee", s).First(&config).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				config.ConfName = "expolee"
				config.ConfKey = s
				config.ConfValue = s2
				err = z.DB.Create(&config).Error
				if err != nil {
					return
				}
			} else {
				return
			}
		} else {
			err = z.DB.Model(&ZbpConfig{}).Where("conf_Name = ? AND conf_Key = ?", "expolee", s).Update("conf_Value", s2).Error
			if err != nil {
				return
			}
		}
	}

	return
}

func (z *ZBolg) setHopolee() (err error) {
	// 设置主题的数据 expolee

	postime, _ := slice.Random([]string{"1", "2", "3"})

	catstyle, _ := slice.Random([]string{"1", "2"})
	nocatitle, _ := slice.Random([]string{"1", "0"})

	webtime := datetime.GetNowDate()

	var configs = map[string]string{
		"flashlights": `s:1:"1";`,
		"webtitle":    fmt.Sprintf(`s:%d:"%s";`, len(z.Title), z.Title),
		"websubtitle": fmt.Sprintf(`s:%d:"%s";`, len(z.SubTitle), z.SubTitle),
		"Keywords":    fmt.Sprintf(`s:%d:"%s";`, len(z.Keywords), z.Keywords),
		"Description": fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
		"qqkfoff":     `s:1:"0";`,
		"toploginoff": `s:1:"0";`,
		"onqqlogin":   `s:1:"0";`,
		"linkoff":     `s:1:"1";`,
		"ftrwenzi":    fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
		"authoroff":   `s:1:"0";`,
		"synavtaboff": `s:1:"0";`,
		"sypicidoff":  `s:1:"0";`,
		"sypicadoff":  `s:1:"0";`,
		"sytextidoff": `s:1:"0";`,
		"sytextadoff": `s:1:"0";`,
		"dbnavbq":     `s:0:"";`,
		"icpbeian":    `s:0:"";`,
		"webtime":     fmt.Sprintf(`s:%d:"%s";`, len(webtime), webtime),
		"toptextoff":  `s:1:"0";`,
		"zdywzseo":    `s:1:"1";`,
		"introSource": `s:1:"1";`,
		"listree":     `s:1:"1";`,
		"fjzhon":      `s:1:"1";`,
		"nighton":     `s:1:"1";`,
		"runtimeoff":  `s:1:"1";`,
		"cysortoff":   `s:1:"1";`,
		"msideoff":    `s:1:"1";`,
		"wtsjon":      `s:1:"1";`,
		"catstyle":    fmt.Sprintf(`s:%d:"%s";`, len(catstyle), catstyle),
		"nocatitle":   fmt.Sprintf(`s:%d:"%s";`, len(nocatitle), nocatitle),
		"postime":     fmt.Sprintf(`s:%d:"%s";`, len(postime), postime),
	}

	for s, s2 := range configs {
		// 获取数据，没有数据则插入，有数据则更新
		var config ZbpConfig

		err = z.DB.Where("conf_Name = ? AND conf_Key = ?", "hopelee", s).First(&config).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				config.ConfName = "hopelee"
				config.ConfKey = s
				config.ConfValue = s2
				err = z.DB.Create(&config).Error
				if err != nil {
					return
				}
			} else {
				return
			}
		} else {
			err = z.DB.Model(&ZbpConfig{}).Where("conf_Name = ? AND conf_Key = ?", "hopelee", s).Update("conf_Value", s2).Error
			if err != nil {
				return
			}
		}
	}

	return
}

func (z *ZBolg) setJz52Acolumn() (err error) {
	var configs = map[string]string{
		"checkbox":       `s:1:"1";`,
		"xgwz":           `s:1:"1";`,
		"themes":         `s:1:"1";`,
		"lin":            `s:1:"1";`,
		"zbbq":           `s:1:"1";`,
		"og	":            `s:1:"1";`,
		"sodes	":         `s:1:"1";`,
		"nseo	":          `s:1:"1";`,
		"imgseo	":        `s:1:"1";`,
		"imgbox	":        `s:1:"1";`,
		"cats	":          `s:1:"1";`,
		"tagdes	":        `s:1:"1";`,
		"seo	":           `s:1:"1";`,
		"seotitle":       fmt.Sprintf(`s:%d:"%s";`, len(z.Title), z.Title),
		"seokeywords":    fmt.Sprintf(`s:%d:"%s";`, len(z.Keywords), z.Keywords),
		"seodescription": fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
	}

	for s, s2 := range configs {
		// 获取数据，没有数据则插入，有数据则更新
		var config ZbpConfig

		err = z.DB.Where("conf_Name = ? AND conf_Key = ?", "Jz52_acolumn", s).First(&config).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				config.ConfName = "Jz52_acolumn"
				config.ConfKey = s
				config.ConfValue = s2
				err = z.DB.Create(&config).Error
				if err != nil {
					return
				}
			} else {
				return
			}
		} else {
			err = z.DB.Model(&ZbpConfig{}).Where("conf_Name = ? AND conf_Key = ?", "Jz52_acolumn", s).Update("conf_Value", s2).Error
			if err != nil {
				return
			}
		}
	}

	return
}

func (z *ZBolg) setJz52Autumn() (err error) {
	fanyeanniu, _ := slice.Random([]string{"1", "2", "0"})
	lbtime, _ := slice.Random([]string{"1", "2", "0"})
	wztime, _ := slice.Random([]string{"1", "2", "0"})

	jzwebTime := datetime.GetNowDate()

	var configs = map[string]string{
		"imglogo":        `s:1:"1";`,
		"szmtx":          `s:1:"1";`,
		"thumb":          `s:1:"1";`,
		"imgcheck":       `s:1:"1";`,
		"themes":         `s:1:"1";`,
		"zbbq":           `s:1:"1";`,
		"seo":            `s:1:"1";`,
		"cltime":         `s:1:"1";`,
		"zanshang":       `s:1:"0";`,
		"zstxt":          `s:0:"";`,
		"footnav":        `s:0:"";`,
		"footico":        `s:0:"";`,
		"wzggtxt":        fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
		"fanyeanniu":     fmt.Sprintf(`s:%d:"%s";`, len(fanyeanniu), fanyeanniu),
		"lbtime":         fmt.Sprintf(`s:%d:"%s";`, len(lbtime), lbtime),
		"wztime":         fmt.Sprintf(`s:%d:"%s";`, len(wztime), wztime),
		"jzwebTime":      fmt.Sprintf(`s:%d:"%s";`, len(jzwebTime), jzwebTime),
		"nseo	":          `s:1:"1";`,
		"imgseo	":        `s:1:"1";`,
		"sjyz	":          `s:1:"1";`,
		"seotitle":       fmt.Sprintf(`s:%d:"%s";`, len(z.Title), z.Title),
		"seokeywords":    fmt.Sprintf(`s:%d:"%s";`, len(z.Keywords), z.Keywords),
		"seodescription": fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
	}

	for s, s2 := range configs {
		// 获取数据，没有数据则插入，有数据则更新
		var config ZbpConfig

		err = z.DB.Where("conf_Name = ? AND conf_Key = ?", "Jz52_autumn", s).First(&config).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				config.ConfName = "Jz52_autumn"
				config.ConfKey = s
				config.ConfValue = s2
				err = z.DB.Create(&config).Error
				if err != nil {
					return
				}
			} else {
				return
			}
		} else {
			err = z.DB.Model(&ZbpConfig{}).Where("conf_Name = ? AND conf_Key = ?", "Jz52_autumn", s).Update("conf_Value", s2).Error
			if err != nil {
				return
			}
		}
	}

	return
}

func (z *ZBolg) setJz52Guopi() (err error) {
	fanyeanniu, _ := slice.Random([]string{"1", "2", "0"})
	var configs = map[string]string{
		"nseo":           `s:1:"1";`,
		"imgseo":         `s:1:"1";`,
		"themes":         `s:1:"1";`,
		"zbbq":           `s:1:"1";`,
		"seo":            `s:1:"1";`,
		"copyright":      `s:1:"1";`,
		"zanshang":       `s:1:"0";`,
		"seomb":          `s:0:"";`,
		"footnav":        `s:0:"";`,
		"sige":           `s:1:"0";`,
		"footico":        `s:0:"";`,
		"foottxt":        fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
		"fanyeanniu":     fmt.Sprintf(`s:%d:"%s";`, len(fanyeanniu), fanyeanniu),
		"seotitle":       fmt.Sprintf(`s:%d:"%s";`, len(z.Title), z.Title),
		"seokeywords":    fmt.Sprintf(`s:%d:"%s";`, len(z.Keywords), z.Keywords),
		"seodescription": fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
	}
	err = z.saveSql(configs, "Jz52_guopi")
	return
}

func (z *ZBolg) setJz52Maxr() (err error) {
	fanyeanniu, _ := slice.Random([]string{"1", "2", "0"})
	wztime, _ := slice.Random([]string{"1", "2", "0"})

	var configs = map[string]string{
		"us":             `s:1:"0";`,
		"seo":            `s:1:"1";`,
		"nseo":           `s:1:"1";`,
		"themes":         `s:1:"1";`,
		"zbbq":           `s:1:"1";`,
		"relatedlist":    `s:1:"1";`,
		"zanshangm":      `s:1:"0";`,
		"indextop":       `s:1:"0";`,
		"zanshang":       `s:1:"0";`,
		"zszfb":          `s:0:"";`,
		"zswx":           `s:0:"";`,
		"classdiy":       `s:0:"";`,
		"side_tui_list":  `s:0:"";`,
		"wztime":         fmt.Sprintf(`s:%d:"%s";`, len(wztime), wztime),
		"fanyeanniu":     fmt.Sprintf(`s:%d:"%s";`, len(fanyeanniu), fanyeanniu),
		"seotitle":       fmt.Sprintf(`s:%d:"%s";`, len(z.Title), z.Title),
		"seokeywords":    fmt.Sprintf(`s:%d:"%s";`, len(z.Keywords), z.Keywords),
		"seodescription": fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
	}
	err = z.saveSql(configs, "Jz52_maxr")
	return
}

func (z *ZBolg) setJz52Tsc() (err error) {
	fanyeanniu, _ := slice.Random([]string{"1", "2", "0"})
	wztime, _ := slice.Random([]string{"1", "2", "0"})
	lbtime, _ := slice.Random([]string{"1", "2", "0"})

	var configs = map[string]string{
		"seo":            `s:1:"1";`,
		"nseo":           `s:1:"1";`,
		"imgseo":         `s:1:"1";`,
		"tagdes":         `s:1:"1";`,
		"zbbq":           `s:1:"1";`,
		"hdpyc":          `s:1:"1";`,
		"cltime":         `s:1:"1";`,
		"copyright":      `s:1:"1";`,
		"xgwz":           `s:1:"1";`,
		"zanshang":       `s:1:"0";`,
		"zszfb":          `s:0:"";`,
		"zswx":           `s:0:"";`,
		"lbtime":         fmt.Sprintf(`s:%d:"%s";`, len(lbtime), lbtime),
		"wztime":         fmt.Sprintf(`s:%d:"%s";`, len(wztime), wztime),
		"fanyeanniu":     fmt.Sprintf(`s:%d:"%s";`, len(fanyeanniu), fanyeanniu),
		"seotitle":       fmt.Sprintf(`s:%d:"%s";`, len(z.Title), z.Title),
		"seokeywords":    fmt.Sprintf(`s:%d:"%s";`, len(z.Keywords), z.Keywords),
		"seodescription": fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
	}
	err = z.saveSql(configs, "Jz52_tsc")
	return
}

func (z *ZBolg) setJzZpojie() (err error) {

	navradio, _ := slice.Random([]string{"1", "2"})
	pageon, _ := slice.Random([]string{"1", "2", "3"})

	side_tui_list_slice := make([]string, 5)
	for i, id := range z.articleIds {
		if i > 4 {
			break
		}
		side_tui_list_slice[i] = fmt.Sprintf("%d", id)
	}
	side_tui_list := strings.Join(side_tui_list_slice, ",")

	var configs = map[string]string{
		"seo":            `s:1:"1";`,
		"nseo":           `s:1:"1";`,
		"zbbq":           `s:1:"1";`,
		"newred":         `s:1:"1";`,
		"yhtim":          `s:1:"1";`,
		"hdpyc":          `s:1:"1";`,
		"commentlog":     `s:1:"1";`,
		"pageon":         fmt.Sprintf(`s:%d:"%s";`, len(pageon), pageon),
		"foottit2":       `s:0:"";`,
		"foottit2code":   `s:0:"";`,
		"navradio":       fmt.Sprintf(`s:%d:"%s";`, len(navradio), navradio),
		"side_tui_list":  fmt.Sprintf(`s:%d:"%s";`, len(side_tui_list), side_tui_list),
		"seotitle":       fmt.Sprintf(`s:%d:"%s";`, len(z.Title), z.Title),
		"seokeywords":    fmt.Sprintf(`s:%d:"%s";`, len(z.Keywords), z.Keywords),
		"seodescription": fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
		"foottit1txt":    fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
	}
	err = z.saveSql(configs, "jz_zpojie")
	return
}

func (z *ZBolg) setKoilee() (err error) {

	catstyle, _ := slice.Random([]string{"1", "2"})
	nocatitle, _ := slice.Random([]string{"1", "0"})
	banstyle, _ := slice.Random([]string{"1", "2"})
	postime, _ := slice.Random([]string{"1", "2", "3"})

	var configs = map[string]string{
		"imgtitle":  `s:0:"";`,
		"imgwzms":   `s:0:"";`,
		"imgtitle2": `s:0:"";`,
		"imgwzms2":  `s:0:"";`,
		"imgtitle3": `s:0:"";`,
		"imgwzms3":  `s:0:"";`,
		"imgtitle4": `s:0:"";`,
		"imgwzms4":  `s:0:"";`,

		"dbnavbq":     `s:0:"";`,
		"icpbeian":    `s:0:"";`,
		"gabbeian":    `s:0:"";`,
		"qqkfoff":     `s:0:"";`,
		"wxkfoff":     `s:0:"";`,
		"sidebarggnr": `s:0:"";`,
		"toploginoff": `s:1:"0";`,
		"linkoff":     `s:1:"0";`,
		"jxddoff":     `s:1:"0";`,
		"sycms01off":  `s:1:"0";`,
		"hxcmsl2off":  `s:1:"0";`,
		"textoff":     `s:1:"0";`,
		"hxcmsloff":   `s:1:"0";`,
		"zdywzseo":    `s:1:"1";`,
		"wzcopyright": `s:1:"1";`,
		"titleoff":    `s:1:"1";`,
		"slideoff":    `s:1:"1";`,
		"fjzhon":      `s:1:"1";`,
		"runtimeoff":  `s:1:"1";`,
		"catstyle":    fmt.Sprintf(`s:%d:"%s";`, len(catstyle), catstyle),
		"nocatitle":   fmt.Sprintf(`s:%d:"%s";`, len(nocatitle), nocatitle),
		"banstyle":    fmt.Sprintf(`s:%d:"%s";`, len(banstyle), banstyle),
		"postime":     fmt.Sprintf(`s:%d:"%s";`, len(postime), postime),
		"webtitle":    fmt.Sprintf(`s:%d:"%s";`, len(z.Title), z.Title),
		"websubtitle": fmt.Sprintf(`s:%d:"%s";`, len(z.SubTitle), z.SubTitle),
		"Keywords":    fmt.Sprintf(`s:%d:"%s";`, len(z.Keywords), z.Keywords),
		"Description": fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
		"ftwenzi":     fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
	}
	err = z.saveSql(configs, "koilee")
	return
}

func (z *ZBolg) setMaocGrace() (err error) {

	fanyeanniu, _ := slice.Random([]string{"1", "2", "0"})

	var configs = map[string]string{
		"Thumb":      `s:1:"1";`,
		"hdp":        `s:1:"1";`,
		"themes":     `s:1:"1";`,
		"zbbq":       `s:1:"1";`,
		"listnavi":   `s:0:"";`,
		"seo":        `s:1:"1";`,
		"nseo":       `s:1:"1";`,
		"tims":       `s:1:"1";`,
		"fanyeanniu": fmt.Sprintf(`s:%d:"%s";`, len(fanyeanniu), fanyeanniu),

		"seotitle":       fmt.Sprintf(`s:%d:"%s";`, len(z.Title), z.Title),
		"seokeywords":    fmt.Sprintf(`s:%d:"%s";`, len(z.Keywords), z.Keywords),
		"seodescription": fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
	}
	err = z.saveSql(configs, "maoc_grace")
	return
}

func (z *ZBolg) setMaocYaa() (err error) {

	fanyeanniu, _ := slice.Random([]string{"1", "2", "0"})

	var configs = map[string]string{
		"hdp":        `s:1:"1";`,
		"themes":     `s:1:"1";`,
		"ttrm":       `s:1:"0";`,
		"zbbq":       `s:1:"1";`,
		"seo":        `s:1:"1";`,
		"nseo":       `s:1:"1";`,
		"tims":       `s:1:"1";`,
		"fanyeanniu": fmt.Sprintf(`s:%d:"%s";`, len(fanyeanniu), fanyeanniu),

		"seotitle":       fmt.Sprintf(`s:%d:"%s";`, len(z.Title), z.Title),
		"seokeywords":    fmt.Sprintf(`s:%d:"%s";`, len(z.Keywords), z.Keywords),
		"seodescription": fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
	}
	err = z.saveSql(configs, "maoc_yaa")
	return
}

func (z *ZBolg) setUmModern() (err error) {

	umNavPage, _ := slice.Random([]string{"1", "2", "0"})
	umListTimeFor, _ := slice.Random([]string{"1", "2", "0", "4", "5"})
	umPostTimeFor, _ := slice.Random([]string{"1", "2", "0", "4", "5"})
	umListTime, _ := slice.Random([]string{"1", "0"})
	umPostTime, _ := slice.Random([]string{"1", "0"})

	var configs = map[string]string{
		"ftNav":         `s:1:"0";`,
		"umCrumbs":      `s:1:"1";`,
		"umAni":         `s:1:"1";`,
		"umCateTit":     `s:1:"1";`,
		"umThemeBy":     `s:1:"1";`,
		"zbBy":          `s:1:"1";`,
		"indexSeo":      `s:1:"1";`,
		"postBaidu":     `s:1:"1";`,
		"umDes":         `s:1:"1";`,
		"umCatalog":     `s:1:"1";`,
		"umQRcode":      `s:1:"1";`,
		"umReadmoreA":   `s:1:"1";`,
		"umReadmoreB":   `s:1:"1";`,
		"loginOff":      `s:1:"0";`,
		"umListTime":    fmt.Sprintf(`s:%d:"%s";`, len(umListTime), umListTime),
		"umPostTime":    fmt.Sprintf(`s:%d:"%s";`, len(umPostTime), umPostTime),
		"umListTimeFor": fmt.Sprintf(`s:%d:"%s";`, len(umListTimeFor), umListTimeFor),
		"umPostTimeFor": fmt.Sprintf(`s:%d:"%s";`, len(umPostTimeFor), umPostTimeFor),
		"umNavPage":     fmt.Sprintf(`s:%d:"%s";`, len(umNavPage), umNavPage),

		"indexTitle": fmt.Sprintf(`s:%d:"%s";`, len(z.Title), z.Title),
		"indexKwds":  fmt.Sprintf(`s:%d:"%s";`, len(z.Keywords), z.Keywords),
		"indexDesc":  fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
	}
	err = z.saveSql(configs, "umModern")
	return
}

func (z *ZBolg) setZbox() (err error) {

	friendlinks, _ := slice.Random([]string{"1", "2"})
	var configs = map[string]string{
		"weibo":             `s:0:"";`,
		"pagesArray":        `s:0:"";`,
		"leonhereCopyright": `s:1:"0";`,
		"seoON":             `s:1:"1";`,
		"friendlinks":       fmt.Sprintf(`s:%d:"%s";`, len(friendlinks), friendlinks),
		"title":             fmt.Sprintf(`s:%d:"%s";`, len(z.Title), z.Title),
		"keywords":          fmt.Sprintf(`s:%d:"%s";`, len(z.Keywords), z.Keywords),
		"description":       fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
	}
	err = z.saveSql(configs, "zbox")
	return
}

func (z *ZBolg) setMxlee() (err error) {
	nocatitle, _ := slice.Random([]string{"1", "0"})
	postime, _ := slice.Random([]string{"1", "2", "3"})
	var configs = map[string]string{
		"d_about":     `s:0:"";`,
		"icpbeian":    `s:0:"";`,
		"tongji":      `s:0:"";`,
		"gabbeian":    `s:0:"";`,
		"qqkfoff":     `s:1:"0";`,
		"userregoff":  `s:1:"0";`,
		"topcmsoff":   `s:1:"0";`,
		"imgtitle":    `s:0:"";`,
		"imgwzms":     `s:0:"";`,
		"imgtitle2":   `s:0:"";`,
		"imgwzms2":    `s:0:"";`,
		"imgtitle3":   `s:0:"";`,
		"imgwzms3":    `s:0:"";`,
		"imgtitle4":   `s:0:"";`,
		"imgwzms4":    `s:0:"";`,
		"flashlights": `s:1:"1";`,
		"zdywzseo":    `s:1:"1";`,
		"listree":     `s:1:"1";`,
		"imgbox":      `s:1:"1";`,
		"introSource": `s:1:"1";`,
		"nocatitle":   fmt.Sprintf(`s:%d:"%s";`, len(nocatitle), nocatitle),
		"postime":     fmt.Sprintf(`s:%d:"%s";`, len(postime), postime),
		"Keywords":    fmt.Sprintf(`s:%d:"%s";`, len(z.Keywords), z.Keywords),
		"Description": fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
		"ftwenzi":     fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
	}
	err = z.saveSql(configs, "mxlee")
	return
}

func (z *ZBolg) setQuietlee() (err error) {
	nocatitle, _ := slice.Random([]string{"1", "0"})
	logotop, _ := slice.Random([]string{"1", "2"})
	indextyle, _ := slice.Random([]string{"1", "2"})
	webtime := datetime.GetNowDate()
	postime, _ := slice.Random([]string{"1", "2", "3"})
	var configs = map[string]string{
		"logotop":      fmt.Sprintf(`s:%d:"%s";`, len(logotop), logotop),
		"indextyle":    fmt.Sprintf(`s:%d:"%s";`, len(indextyle), indextyle),
		"toploginoff":  `s:1:"0";`,
		"sytextidoff":  `s:1:"0";`,
		"qqloginoff":   `s:1:"0";`,
		"shareoff":     `s:1:"0";`,
		"wzzsoff":      `s:1:"0";`,
		"syimgidoff":   `s:1:"0";`,
		"banneroff":    `s:1:"0";`,
		"frlbwzoff":    `s:1:"0";`,
		"fjzhon":       `s:1:"1";`,
		"nprogressoff": `s:1:"1";`,
		"msideoff":     `s:1:"1";`,
		"cnamaon":      `s:1:"0";`,
		"notopbgoff":   `s:1:"0";`,
		"linkoff":      `s:1:"0";`,
		"authoroff":    `s:1:"0";`,
		"zdywzseo":     `s:1:"1";`,
		"readtextoff":  `s:1:"1";`,
		"introSource":  `s:1:"1";`,
		"tougaoff":     `s:1:"1";`,
		"sidebarggnr":  `s:0:"";`,
		"dbnavbq":      `s:0:"";`,
		"icpbeian":     `s:0:"";`,
		"gabbeian":     `s:0:"";`,
		"synavtag":     `s:0:"";`,
		"synavgd":      `s:0:"";`,
		"catlistinfo":  `s:55:"{"user":"1","cate":"1","date":"1","view":"1","cmt":"1"}";`,
		"webtime":      fmt.Sprintf(`s:%d:"%s";`, len(webtime), webtime),
		"postime":      fmt.Sprintf(`s:%d:"%s";`, len(postime), postime),
		"nocatitle":    fmt.Sprintf(`s:%d:"%s";`, len(nocatitle), nocatitle),
		"webtitle":     fmt.Sprintf(`s:%d:"%s";`, len(z.Title), z.Title),
		"websubtitle":  fmt.Sprintf(`s:%d:"%s";`, len(z.SubTitle), z.SubTitle),
		"Keywords":     fmt.Sprintf(`s:%d:"%s";`, len(z.Keywords), z.Keywords),
		"Description":  fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
		"ftwenzi":      fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
	}
	err = z.saveSql(configs, "quietlee")
	return
}

func (z *ZBolg) setYouthlee() (err error) {
	nocatitle, _ := slice.Random([]string{"1", "0"})
	readstyle, _ := slice.Random([]string{"1", "0"})
	postime, _ := slice.Random([]string{"1", "2", "3"})
	readapi, _ := slice.Random([]string{"1", "2", "3"})
	var configs = map[string]string{

		"linkoff":     `s:1:"0";`,
		"cnamaon":     `s:1:"0";`,
		"sytjon":      `s:1:"0";`,
		"synavtaboff": `s:1:"0";`,
		"dbnavbq":     `s:0:"";`,
		"icpbeian":    `s:0:"";`,
		"gabbeian":    `s:0:"";`,
		"shareoff":    `s:1:"0";`,
		"wzzsoff":     `s:1:"0";`,
		"banfridoff":  `s:1:"0";`,
		"tagsidon":    `s:1:"0";`,
		"banneroff":   `s:1:"0";`,
		"introSource": `s:1:"1";`,
		"sideauthon":  `s:1:"1";`,
		"wztopinfo":   `s:65:"{"user":"1","date":"1","view":"1","cmt":"1","edit":"1","del":"1"}";`,
		"catlistinfo": `s:55:"{"cate":"1","user":"1","date":"1","view":"1","cmt":"1"}";`,
		"postime":     fmt.Sprintf(`s:%d:"%s";`, len(postime), postime),
		"readapi":     fmt.Sprintf(`s:%d:"%s";`, len(readapi), readapi),
		"readstyle":   fmt.Sprintf(`s:%d:"%s";`, len(readstyle), readstyle),
		"nocatitle":   fmt.Sprintf(`s:%d:"%s";`, len(nocatitle), nocatitle),
		"webtitle":    fmt.Sprintf(`s:%d:"%s";`, len(z.Title), z.Title),
		"websubtitle": fmt.Sprintf(`s:%d:"%s";`, len(z.SubTitle), z.SubTitle),
		"Keywords":    fmt.Sprintf(`s:%d:"%s";`, len(z.Keywords), z.Keywords),
		"Description": fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
		"ftwenzi":     fmt.Sprintf(`s:%d:"%s";`, len(z.Description), z.Description),
	}
	err = z.saveSql(configs, "youthlee")
	return
}

// saveSql
//
//	@Description: 保存数据到数据库
//	@receiver z
//	@param configs map[string]string // 配置
//	@param confName string // 配置名称
//	@return err
func (z *ZBolg) saveSql(configs map[string]string, confName string) (err error) {
	for s, s2 := range configs {
		// 获取数据，没有数据则插入，有数据则更新
		var config ZbpConfig

		err = z.DB.Where("conf_Name = ? AND conf_Key = ?", confName, s).First(&config).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				config.ConfName = confName
				config.ConfKey = s
				config.ConfValue = s2
				err = z.DB.Create(&config).Error
				if err != nil {
					return
				}
			} else {
				err = errors.New(z.WebSite + "查询数据库出错 ：" + err.Error())
				return
			}
		} else {
			err = z.DB.Model(&ZbpConfig{}).Where("conf_Name = ? AND conf_Key = ?", confName, s).Update("conf_Value", s2).Error
			if err != nil {
				return
			}
		}
	}
	return
}
