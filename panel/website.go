// Package panel
// @Description: panel 网站的管理，数据库的管理
package panel

import (
	"cmsManage/utils/reqRequest"
	"cmsManage/utils/setPWD"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/random"
	"github.com/duke-git/lancet/v2/strutil"
	"strconv"
	"strings"
)

func (p *Panel) CreateWebsite(website string) (err error) {

	website = strutil.RemoveWhiteSpace(website, true)
	if website == "" {
		return errors.New(p.Website + "站点名称不能为空")
	}

	p.Website = website

	// 搜索网站是否存在
	_, exist, err := p.searchWebsite()
	if err != nil {
		return
	}

	if exist {
		err = errors.New(p.Website + "网站已存在")
		return
	}

	// 搜索网站的数据库是否存在
	_, exist, err = p.searchDatabase()
	if err != nil {
		return
	}

	if exist {
		err = errors.New(p.Website + "数据库已存在")
		return
	}

	//创建网站
	err = p.createWebsite()
	if err != nil {
		return
	}

	// 创建数据库
	err = p.createDatabase()
	if err != nil { // 失败需要删除网站。。。
		err = p.delWebsite()
		if err != nil {
			return
		}
		err = errors.New(p.Website + "创建数据库失败之后成功删除网站")
		return
	}

	return
}

// DelWebsite
//
//	@Description: 删除网站
//	@receiver p
//	@param website
//	@return err
func (p *Panel) DelWebsite(website string) (err error) {

	website = strutil.RemoveWhiteSpace(website, true)
	if website == "" {
		return errors.New(p.Website + "站点名称不能为空")
	}
	p.Website = website

	err = p.delWebsite()
	if err != nil {
		return
	}

	err = p.delDatabase()
	if err != nil {
		return
	}
	return
}

// createWebsite
//
//	@Description: 创建网站
//	@receiver p
//	@return err
func (p *Panel) createWebsite() (err error) {

	reqUrl := loginHost + "/api/v1/websites"
	for i := 0; i < 3; i++ {

		port := 9000 + random.RandInt(22222, 55555)

		jsonData := map[string]any{
			"primaryDomain":  p.Website,
			"type":           "runtime",
			"alias":          p.Website,
			"remark":         p.Website,
			"appType":        "installed",
			"webSiteGroupId": 2,
			"otherDomains":   "",
			"proxy":          "",
			"runtimeID":      1,
			"appinstall": map[string]any{
				"appId":       0,
				"name":        "",
				"appDetailId": 122,
				"params": map[string]any{
					"PANEL_APP_PORT_HTTP": port,
				},
				"version":       "",
				"appkey":        "",
				"advanced":      true,
				"cpuQuota":      0,
				"memoryLimit":   0,
				"memoryUnit":    "MB",
				"containerName": p.Website,
				"allowPort":     false,
			},
			"IPV6":          false,
			"enableFtp":     false,
			"ftpUser":       "",
			"ftpPassword":   "",
			"proxyType":     "tcp",
			"port":          9000,
			"proxyProtocol": "http://",
			"proxyAddress":  "",
			"runtimeType":   "php",
		}

		var response reqRequest.Response

		response, err = p.Session.Post(reqRequest.RequestOption{
			Url:  reqUrl,
			Json: jsonData,
		})
		if err != nil {
			fmt.Println("创建网站失败: ", err.Error())
			continue
		}

		sourceHtml := response.SourceHtml
		if strings.Contains(sourceHtml, "端口已被应用") {
			fmt.Println("端口已被应用，将进行重试...")
			continue
		} else if strings.Contains(sourceHtml, "代号已存在") {
			err = errors.New(p.Website + "代号已存在")
			return
		} else {
			if !strings.Contains(sourceHtml, `"code":200`) {
				continue
			}
			return
		}
	}
	return nil
}

func (p *Panel) createDatabase() (err error) {
	reqUrl := loginHost + "/api/v1/databases"

	dbName, dbPwd := setPWD.GetSiteDbname(p.Website)
	dbPwd = base64.StdEncoding.EncodeToString([]byte(dbPwd))
	jsonData := map[string]any{
		"name":          dbName,
		"from":          "local",
		"type":          "mysql",
		"database":      "mysql",
		"format":        "utf8mb4",
		"username":      dbName,
		"password":      dbPwd,
		"permission":    "%",
		"permissionIPs": "",
		"description":   p.Website + " 网站的数据库",
	}

	response, err := p.Session.Post(reqRequest.RequestOption{
		Url:  reqUrl,
		Json: jsonData,
	})
	if err != nil {
		return
	}

	sourceHtml := response.SourceHtml
	if strings.Contains(sourceHtml, `"code":200`) {
		return
	}
	err = errors.New(p.Website + "创建数据库失败")
	return
}

// searchWebsite
//
//	@Description: 搜索网站
//	@receiver p
//	@return iD 如果网站存在，返回网站的ID
//	@return exist 是否存在 true 存在 false 不存在
//	@return err 错误信息
func (p *Panel) searchWebsite() (iD int, exist bool, err error) {
	exist = false

	reqUrl := loginHost + "/api/v1/websites/search"

	jsonData := map[string]any{
		"name":           p.Website,
		"page":           1,
		"pageSize":       10,
		"orderBy":        "created_at",
		"order":          "null",
		"websiteGroupId": 0,
	}

	response, err := p.Session.Post(reqRequest.RequestOption{
		Url:  reqUrl,
		Json: jsonData,
	})
	if err != nil {
		return
	}

	status := response.StatusCode
	if status != 200 {
		err = errors.New(p.Website + "搜索网站失败")
		return
	}

	sourceHtml := response.SourceHtml
	if strings.Contains(sourceHtml, `"`+p.Website+`"`) { // 加上引号，防止 aa.com 匹配到 a.com
		exist = true
		res := strutil.Before(sourceHtml, `"`+p.Website+`"`)
		res = strutil.After(res, `"id":`)
		res = strutil.Before(res, `,`)
		iD, err = strconv.Atoi(res)
		if err != nil {
			return 0, false, err
		}
	}
	return
}

// searchDatabase
//
//	@Description: 搜索数据库
//	@receiver p
//	@return iD 如果数据库存在，返回数据库的ID
//	@return exist 是否存在 true 存在 false 不存在
//	@return err 错误信息
func (p *Panel) searchDatabase() (iD int, exist bool, err error) {
	exist = false

	reqUrl := loginHost + "/api/v1/databases/search"

	dbName, _ := setPWD.GetSiteDbname(p.Website)

	jsonData := map[string]any{
		"page":     1,
		"pageSize": 10,
		"info":     dbName,
		"database": "mysql",
		"orderBy":  "created_at",
		"order":    "null",
	}

	response, err := p.Session.Post(reqRequest.RequestOption{
		Url:  reqUrl,
		Json: jsonData,
	})
	if err != nil {
		return
	}

	status := response.StatusCode
	if status != 200 {
		err = errors.New(p.Website + "搜索数据库失败")
		return
	}

	sourceHtml := response.SourceHtml
	if strings.Contains(sourceHtml, `"`+dbName+`"`) { // 加上引号，防止 aa.com 匹配到 a.com
		exist = true
		res := strutil.Before(sourceHtml, `"`+dbName+`"`)
		res = strutil.After(res, `"id":`)
		res = strutil.Before(res, `,`)
		iD, err = strconv.Atoi(res)
		if err != nil {
			return 0, false, err
		}
	}
	return
}

func (p *Panel) delWebsite() (err error) {
	iD, exist, err := p.searchWebsite()
	if err != nil {
		return
	}

	if !exist { // website 不存在
		return
	}

	reqUrl := loginHost + "/api/v1/websites/del"

	jsonData := map[string]any{
		"id":           iD,
		"deleteApp":    true,
		"deleteBackup": true,
		"forceDelete":  true,
	}

	response, err := p.Session.Post(reqRequest.RequestOption{
		Url:  reqUrl,
		Json: jsonData,
	})
	if err != nil {
		return
	}

	status := response.StatusCode
	if status != 200 {
		err = errors.New(p.Website + "删除网站失败")
		return
	}
	return
}

func (p *Panel) delDatabase() (err error) {
	iD, exist, err := p.searchDatabase()
	if err != nil {
		return
	}

	if !exist { // 数据库不存在
		return
	}

	reqUrl := loginHost + "/api/v1/databases/del"

	jsonData := map[string]any{
		"id":           iD,
		"type":         "mysql",
		"database":     "mysql",
		"deleteBackup": true,
		"forceDelete":  true,
	}

	response, err := p.Session.Post(reqRequest.RequestOption{
		Url:  reqUrl,
		Json: jsonData,
	})

	if err != nil {
		return
	}

	status := response.StatusCode
	if status != 200 {
		err = errors.New(p.Website + "删除数据库失败")
		return
	}
	return
}
