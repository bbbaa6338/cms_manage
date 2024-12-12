// Package zbolg
// @Description: 网站设置
package zbolg

import (
	"errors"
	"fmt"
)

func (z *ZBolg) WebSetting(title, subTitle, copyright string) (err error) {
	//if title == "" || subTitle == "" || copyright == "" {
	//	err = errors.New("TDK 参数不能为空")
	//	return
	//}
	// 查询基础的数据
	keys := []string{"ZC_BLOG_NAME", "ZC_BLOG_SUBNAME", "ZC_BLOG_COPYRIGHT"}

	var cf []ZbpConfig
	err = z.DB.Where("conf_Key IN ?", keys).Find(&cf).Error
	if err != nil {
		err = errors.New(z.WebSite + " 查询数据库出错 ：" + err.Error())
		return
	}

	if len(keys) != len(cf) {
		err = errors.New(z.WebSite + " 查询数据库出错, 未找到所有的配置")
		return
	}

	for _, config := range cf {
		switch config.ConfKey {
		case "ZC_BLOG_NAME":
			config.ConfValue = fmt.Sprintf(`s:%d:"%s";`, len(title), title)
			err = z.DB.Save(&config).Error
			if err != nil {
				err = errors.New(z.WebSite + " 更新 title 失败 ：" + err.Error())
				return
			}
		case "ZC_BLOG_SUBNAME":
			config.ConfValue = fmt.Sprintf(`s:%d:"%s";`, len(subTitle), subTitle)
			err = z.DB.Save(&config).Error
			if err != nil {
				err = errors.New(z.WebSite + " 更新 subTitle 失败 ：" + err.Error())
				return
			}
		case "ZC_BLOG_COPYRIGHT":
			config.ConfValue = fmt.Sprintf(`s:%d:"%s";`, len(copyright), copyright)
			err = z.DB.Save(&config).Error
			if err != nil {
				err = errors.New(z.WebSite + " 更新 copyRight 失败 ：" + err.Error())
				return
			}
		default:
			err = errors.New(z.WebSite + " 未知的配置项 ：" + config.ConfKey)
			return
		}
	}
	return
}
