// Package zbolg
// @Description: zbolg 的插件设置
package zbolg

import (
	"cmsManage/utils/reqRequest"
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/random"
	"github.com/duke-git/lancet/v2/strutil"
	"strings"
)

var (
	plugs = map[string]string{
		"RunAnyOne": "启用任意应用 1.0", // 必须第一个启动

		"cat_spider":   "蜘蛛日志/爬虫日志/分析插件",
		"ly_cache":     "老阳插件：百万数据库优化缓解服务器爆红",
		"ly_homego":    "搜索来路屏蔽拦截跳转引流",
		"ly_sitemap":   "sitemap多级目录索引型网站地图",
		"ly_themex":    "编译主题伪原创模板",
		"ly_autoimage": "LY_文章匹配补图，标签标题匹配图片",
		"STACentre":    "控制ZBLOG的伪静态化设置.",
	}
)

func (z *ZBolg) Plugin() (err error) {
	if err = z.checkPlugin(); err != nil {
		return
	}

	if err = z.startUpPlug(); err != nil {
		return
	}

	if err = z.setPlugin(); err != nil {
		return
	}

	return
}

func (z *ZBolg) checkPlugin() (err error) {
	plugUrl := strutil.Before(z.LoginUrl, z.WebSite) + z.WebSite + "/zb_system/admin/index.php?act=PluginMng"
	resp, err := z.Session.Get(reqRequest.RequestOption{
		Url: plugUrl,
	})
	if err != nil {
		return
	}

	status := resp.StatusCode
	if status != 200 {
		err = errors.New(z.WebSite + " checkPlugin error, status code is not 200")
		return
	}

	sourceHtml := resp.SourceHtml

	for k, v := range plugs {
		if strings.Contains(sourceHtml, k) {
			continue
		}
		err = errors.New(z.WebSite + " startUpPlugin error, " + v + " not found")
		return
	}
	return
}

func (z *ZBolg) startUpPlug() (err error) {
	// 优先启动 RunAnyOne，没有这个插件有的插件无法启动
	plugRunAnyOneUrl := strutil.Before(z.LoginUrl, z.WebSite) + z.WebSite + "/zb_system/cmd.php"

	var resp reqRequest.Response
	resp, err = z.Session.Get(reqRequest.RequestOption{
		Url: plugRunAnyOneUrl,
		Params: map[string]string{
			"act":       "PluginEnb",
			"name":      "RunAnyOne",
			"csrfToken": z.csrfToken,
		},
	})
	if err != nil {
		return
	}

	status := resp.StatusCode
	if status != 200 {
		err = errors.New(z.WebSite + " startUpPlugin error RunAnyOne  status code is not 200")
		return
	}

	// 启动其他插件
	for k, v := range plugs {
		if k == "RunAnyOne" {
			continue
		}
		plugUrl := strutil.Before(z.LoginUrl, z.WebSite) + z.WebSite + "/zb_system/cmd.php"

		var resp reqRequest.Response
		resp, err = z.Session.Get(reqRequest.RequestOption{
			Url: plugUrl,
			Params: map[string]string{
				"act":       "PluginEnb",
				"name":      k,
				"csrfToken": z.csrfToken,
			},
		})
		if err != nil {
			return
		}

		status := resp.StatusCode
		if status != 200 {
			err = errors.New(z.WebSite + " startUpPlugin error " + v + "  status code is not 200")
			return
		}
	}

	return
}

func (z *ZBolg) setPlugin() (err error) {
	for k, v := range plugs {
		if k == "ly_cache" { // 缓存在建站最后才设置
			continue
		}

		var data map[string]string
		var params map[string]string
		var reqUrl string

		reqUrl = strutil.Before(z.LoginUrl, z.WebSite) + z.WebSite + "/zb_users/plugin/" + k + "/main.php"
		switch k {
		case "STACentre": // 伪静态化设置
			categoryCount := random.RandInt(3, 6)
			category := random.RandLower(categoryCount)
			data = map[string]string{
				"csrfToken":         z.csrfToken,
				"reset":             "",
				"ZC_STATIC_MODE":    "REWRITE",
				"ZC_ARTICLE_REGEX":  "{%host%}{%category%}/article/{%alias%}.html",
				"ZC_PAGE_REGEX":     "{%host%}{%id%}.html",
				"ZC_INDEX_REGEX":    "{%host%}page_{%page%}.html",
				"ZC_CATEGORY_REGEX": "{%host%}" + category + "/{%alias%}/{%page%}/",
				"ZC_TAGS_REGEX":     "{%host%}tags/{%alias%}/{%page%}/",
				"ZC_DATE_REGEX":     "{%host%}date/{%date%}_{%page%}.html",
				"ZC_AUTHOR_REGEX":   "{%host%}author/{%alias%}/{%page%}/",
			}
		case "cat_spider":
			dbExists := z.checkTableExists(z.DB, "zbp_cat_spider")
			if !dbExists {
				err = errors.New(z.WebSite + " setPlugin error " + v + "  table zbp_cat_spider not found")
				return
			}

			params = map[string]string{
				"act":       "save",
				"csrfToken": z.csrfToken,
			}

			data = map[string]string{
				"baidu":     "1",
				"baidum":    "1",
				"google":    "1",
				"biying":    "1",
				"haosou":    "1",
				"sogou":     "1",
				"youdao":    "1",
				"shenma":    "1",
				"toutiao":   "1",
				"other":     "1",
				"csrfToken": z.csrfToken,
			}

		case "ly_homego":
			params = map[string]string{
				"act": "save",
			}

			data = map[string]string{
				"IsLock":    "1",
				"reua[]":    "0",
				"name[]":    "",
				"post[]":    "",
				"cate[]":    "",
				"user[]":    "",
				"tags[]":    "",
				"date[]":    "",
				"type[]":    "ad",
				"data[]":    "",
				"csrfToken": z.csrfToken,
			}
		case "ly_themex":
			params = map[string]string{
				"act": "save",
			}

			slInt := random.RandInt(10, 20)
			slStr := fmt.Sprintf("%d", slInt)
			data = map[string]string{
				"is":        "1",
				"st":        "",
				"sp":        "",
				"sl":        slStr,
				"se":        "1",
				"sh":        "1",
				"sd":        "1",
				"sm":        "1",
				"csrfToken": z.csrfToken,
			}
		case "ly_sitemap":
			params = map[string]string{
				"act": "save",
			}

			data = map[string]string{
				"num":        "2000",
				"dirs":       "",
				"lastmod":    "1",
				"changefreq": "weekly",
				"priority":   "0.8",
				"xml":        "1",
				"txt":        "1",
				"htm":        "1",
				"bot":        "1",
				"tags":       "1",
				"cate":       "1",
				"user":       "1",
				"links":      "网站地图",
				"link":       "2",
				"s404":       "1",
				"cache":      "168",
				"csrfToken":  z.csrfToken,
			}
		case "ly_autoimage":
			params = map[string]string{
				"act": "save",
			}

			data = map[string]string{
				"is":        "1",
				"host":      "",
				"cate":      "",
				"num":       "1-3",
				"p":         "0",
				"txt":       "",
				"align":     "2",
				"iv":        "2",
				"links":     "",
				"isthum":    "",
				"istag":     "0", // 更新为 "0"
				"isatt":     "1",
				"ixatt":     "5", // 更新为 "5"
				"ismode":    "0",
				"proxy":     "0",
				"tags":      "",
				"catex":     "",
				"isimg":     "1",
				"iv1":       "1",   // 更新为 "1"
				"title":     "1",   // 更新为 "1"
				"titles":    "1-3", // 更新为 "1-3"
				"ttf":       "",
				"ttfile":    "",
				"size":      "26",
				"color":     "",
				"bg":        "",
				"top":       "0",
				"box":       "5",
				"bgh":       "0",
				"opacity":   "66",
				"width":     "0",
				"height":    "0",
				"csrfToken": z.csrfToken,
			}
		case "RunAnyOne":
			continue
		default:
			err = errors.New(z.WebSite + " setPlugin error " + v + "  not found")
			return
		}

		if len(data) == 0 {
			continue
		}

		var resp reqRequest.Response
		resp, err = z.Session.Post(reqRequest.RequestOption{
			Url:    reqUrl,
			Data:   data,
			Params: params,
		})
		if err != nil {
			return
		}

		status := resp.StatusCode
		if status != 200 {
			err = errors.New(z.WebSite + " setPlugin error " + v + "  status code is not 200")
			return
		}
	}

	return
}

// LyCache
//
//	@Description: 老阳插件：百万数据库优化缓解服务器爆红
//	@receiver z
//	@return err
func (z *ZBolg) LyCache(site string) (err error) {
	response, err := z.Session.Get(reqRequest.RequestOption{
		Url: "http://" + site + "/zb_users/plugin/ly_cache/main.php",
	})

	if err != nil {
		return
	}

	if response.StatusCode != 200 {
		err = errors.New(site + " LyCache error status code is not 200")
		return
	}

	// 统计 tags 的数量
	var tagsCount int64
	err = z.DB.Model(&ZPBTag{}).Select("tag_ID").Count(&tagsCount).Error
	if err != nil {
		return
	}

	return err
}
