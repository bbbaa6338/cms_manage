package zbolg

import (
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/strutil"
	"strings"
)

//Title       *string        `json:"title" form:"title" gorm:"index;column:title;comment:标题;size:64;" binding:"required"`                                          //标题
//Keywords    *string        `json:"keywords" form:"keywords" gorm:"column:keywords;comment:关键词;size:64;"`                                                         //关键词
//Description *string        `json:"description" form:"description" gorm:"column:description;comment:描述;size:512;"`                                                //描述

//Categories  *string        `json:"categories" form:"categories" gorm:"column:categories;comment:分类;size:128;" binding:"required"`                                //分类
//UserID      *string        `json:"userID" form:"userID" gorm:"column:user_id;comment:用户ID;size:512;" binding:"required"`                                         //用户ID
//UrlConfig   datatypes.JSON `json:"urlConfig" form:"urlConfig" gorm:"column:url_config;comment:URL配置;size:512;type:text;" binding:"required"swaggertype:"object"` //URL配置

func (z *ZBolg) QueryInformation() (tagUrlConfig, userID, categories string, err error) {
	// 查询基础的数据
	var cf ZbpConfig
	err = z.DB.Where("conf_Key = ?", "ZC_TAGS_REGEX").First(&cf).Error
	if err != nil {
		err = errors.New(z.WebSite + "ZC_TAGS_REGEX  查询数据库出错 ：" + err.Error())
		return
	}

	// 获取 tagUrlConfig
	tagUrlConfig = cf.ConfValue
	if tagUrlConfig == "" {
		err = errors.New(z.WebSite + "ZC_TAGS_REGEX 为空")
		return
	}

	// 获取结果
	tagUrlConfig = strutil.After(tagUrlConfig, "{")
	tagUrlConfig = strutil.BeforeLast(tagUrlConfig, "}")
	tagUrlConfig = "{" + tagUrlConfig + "}"

	// 获取所有的分类ID
	var cfs []Category
	err = z.DB.Find(&cfs).Error
	if err != nil {
		err = errors.New(z.WebSite + " Category 查询数据库出错 ：" + err.Error())
		return
	}

	categoriesSlice := make([]string, 0)
	for _, v := range cfs {
		categoriesSlice = append(categoriesSlice, fmt.Sprintf("%d", v.CateID))
	}
	categories = strings.Join(categoriesSlice, ",")

	// 获取所有的用户ID
	var us []ZPBMember
	err = z.DB.Find(&us).Error
	if err != nil {
		err = errors.New(z.WebSite + " ZPBMember 查询数据库出错 ：" + err.Error())
		return
	}

	userIDSlice := make([]string, 0)
	for _, v := range us {
		userIDSlice = append(userIDSlice, fmt.Sprintf("%d", v.MemID))
	}
	userID = strings.Join(userIDSlice, ",")
	return
}

func QueryInformation(site, dbHost string, dbPort int) (tagUrlConfig, userID, categories string, err error) {

	z, err := NewZBolg(
		ZBolg{
			WebSite:    site,
			InstallUrl: "http://" + site + "/zb_install/jsj.php",
			LoginUrl:   "http://" + site + "/zb_system/jsj.php",
			Host:       dbHost,
			Port:       dbPort,
			//Username:   "admin",
			//Password:   "admin",
		})
	if err != nil {
		return
	}

	// 查询基础的数据
	var cf ZbpConfig
	err = z.DB.Where("conf_Key = ?", "ZC_TAGS_REGEX").First(&cf).Error
	if err != nil {
		err = errors.New(z.WebSite + "ZC_TAGS_REGEX  查询数据库出错 ：" + err.Error())
		return
	}

	// 获取 tagUrlConfig
	tagUrlConfig = cf.ConfValue
	if tagUrlConfig == "" {
		err = errors.New(z.WebSite + "ZC_TAGS_REGEX 为空")
		return
	}

	// 获取结果
	tagUrlConfig = strutil.After(tagUrlConfig, "{")
	tagUrlConfig = strutil.BeforeLast(tagUrlConfig, "}")
	tagUrlConfig = "{" + tagUrlConfig + "}"

	// 获取所有的分类ID
	var cfs []Category
	err = z.DB.Find(&cfs).Error
	if err != nil {
		err = errors.New(z.WebSite + " Category 查询数据库出错 ：" + err.Error())
		return
	}

	categoriesSlice := make([]string, 0)
	for _, v := range cfs {
		categoriesSlice = append(categoriesSlice, fmt.Sprintf("%d", v.CateID))
	}
	if len(categoriesSlice) == 0 {
		err = errors.New(z.WebSite + " 无分类")
		return
	}
	categories = strings.Join(categoriesSlice, ",")

	// 获取所有的用户ID
	var us []ZPBMember
	err = z.DB.Find(&us).Error
	if err != nil {
		err = errors.New(z.WebSite + " ZPBMember 查询数据库出错 ：" + err.Error())
		return
	}

	userIDSlice := make([]string, 0)
	for _, v := range us {
		userIDSlice = append(userIDSlice, fmt.Sprintf("%d", v.MemID))
	}
	userID = strings.Join(userIDSlice, ",")
	return
}
