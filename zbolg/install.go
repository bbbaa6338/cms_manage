// Package zbolg
// @Description: zBolg的安装，安装成功会初始化配置：删除分类及文章，关闭评论，关闭登录验证码。新增用户
package zbolg

import (
	"cmsManage/utils"
	"cmsManage/utils/fake"
	"cmsManage/utils/reqRequest"
	toPinyin "cmsManage/utils/to_pinyin"
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/random"
	"github.com/duke-git/lancet/v2/slice"
	"time"
)

func (z *ZBolg) Install() (err error) {
	if err = z.installStep1(); err != nil {
		return
	}

	if err = z.installStep2(); err != nil {
		return
	}

	if err = z.installStep3(); err != nil {
		return
	}

	if err = z.installStep4(); err != nil {
		return
	}

	// 初始化配置
	if err = z.installInitConfig(); err != nil {
		return
	}

	// 设置用户
	if err = z.setAuthor(); err != nil {
		return
	}

	return
}

func (z *ZBolg) installStep1() (err error) {
	resp, err := z.Session.Get(reqRequest.RequestOption{
		Url: z.InstallUrl,
	})
	if err != nil {
		return
	}

	status := resp.StatusCode
	if status != 200 {
		err = errors.New(z.WebSite + "在 Install step1 请求失败, 状态码不对")
		return
	}
	return
}

func (z *ZBolg) installStep2() (err error) {
	resp, err := z.Session.Post(reqRequest.RequestOption{
		Url: z.InstallUrl,
		Params: map[string]string{
			"step": "2",
		},
		Data: map[string]string{
			"zbloglang": "zh-cn",
			"language":  "zh-cn",
			"next":      "下一步",
		},
	})
	if err != nil {
		return
	}

	status := resp.StatusCode
	if status != 200 {
		err = errors.New(z.WebSite + "在 Install step2 请求失败, 状态码不对")
		return
	}
	return
}

func (z *ZBolg) installStep3() (err error) {
	resp, err := z.Session.Post(reqRequest.RequestOption{
		Url: z.InstallUrl,
		Params: map[string]string{
			"step": "3",
		},
		Data: map[string]string{
			"zbloglang": "zh-cn",
			"next":      "下一步",
		},
	})
	if err != nil {
		return
	}

	status := resp.StatusCode
	if status != 200 {
		err = errors.New(z.WebSite + "在 Install step3 请求失败, 状态码不对")
		return
	}
	return
}

func (z *ZBolg) installStep4() (err error) {

	dbServe := fmt.Sprintf("%s:%d", z.Host, z.Port)
	data := map[string]string{
		"zbloglang":        "zh-cn",
		"fdbtype":          "mysql",
		"dbmysql_server":   dbServe,
		"dbmysql_name":     z.dbName,
		"dbmysql_username": z.dbName,
		"dbmysql_password": z.dbPassword,
		"dbmysql_pre":      "zbp_",
		"dbengine":         "InnoDB",
		"dbtype":           "mysqli",
		"dbsqlite_name":    "#%20924e44a0dd1194cb217cd8025232bd7a.db",
		"dbsqlite_pre":     "zbp_",
		"blogtitle":        "网站标题",
		"username":         zBolgUserName,
		"password":         zBolgPassword,
		"repassword":       zBolgPassword,
		"blogtheme":        "default|default",
		"next":             "下一步",
	}

	resp, err := z.Session.Post(reqRequest.RequestOption{
		Url: z.InstallUrl,
		Params: map[string]string{
			"step": "4",
		},
		Data: data,
	})

	if err != nil {
		return
	}

	status := resp.StatusCode
	if status != 200 {
		err = errors.New(z.WebSite + " 在 Install step4 请求失败, 状态码不对")
		return
	}
	return
}

func (z *ZBolg) installInitConfig() (err error) {

	config := map[string]string{
		"ZC_LOGIN_VERIFY_ENABLE": "b:0;", // 禁止登录需要验证码
		"ZC_DEBUG_MODE_WARNING":  "b:0;", // 禁止允许报Warning级别错误
		"ZC_COMMENT_TURNOFF":     "b:0;", // 禁止评论
		"ZC_COMMENT_AUDIT":       "b:1;", // 评论审核
	}

	for s, s2 := range config {
		if err = z.DB.Model(&ZbpConfig{}).Where("conf_Key = ?", s).Update("conf_Value", s2).Error; err != nil {
			err = errors.New(z.WebSite + " 在 " + s + " 失败" + err.Error())
			return err
		}
	}

	// 情况分类及文章表
	tables := []string{
		"zbp_category", "zbp_post",
	}

	var deleteTable = "TRUNCATE TABLE"
	for _, table := range tables {
		sql := fmt.Sprintf("%s %s", deleteTable, table)
		if err := z.DB.Exec(sql).Error; err != nil {
			err = errors.New(z.WebSite + " 在清空表 ：" + table + " 失败")
			return err
		}
	}
	return
}

func (z *ZBolg) setAuthor() (err error) {
	var author int
	if z.AuthorCount == 0 {
		author = random.RandInt(5, 9)
	}

	for i := 0; i < author; i++ {
		userName := fake.GenerateCNName() // 名字，中文

		alias := toPinyin.ToPinyin(userName)

		ll := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
		qqStart, _ := slice.Random(ll)
		qqRandom := random.RandInt(6, 8)
		qq := qqStart + random.RandNumeral(qqRandom)
		// 会员摘要
		memIntro := utils.RandomYiJuHua()

		newMember := ZPBMember{
			MemGuid:       "0977e68b7516fad1ba8c88abc874a52a", // GUID，使用你需要的值
			MemLevel:      3,                                  // 会员等级，例如 1
			MemStatus:     0,                                  // 会员状态，正常状态为 1
			MemName:       userName,                           // 会员姓名
			MemPassword:   "94738b674341e9d0a42c77915301c9fa", // 密码（一般需要加密处理）
			MemEmail:      qq + "@qq.com",                     // 邮箱
			MemHomePage:   "",                                 // 主页
			MemIP:         "192.168.1.1",                      // IP地址
			MemCreateTime: int(time.Now().Unix()),             // 创建时间（以 Unix 时间戳为例）
			MemPostTime:   int(time.Now().Unix()),             // 发帖时间
			MemUpdateTime: int(time.Now().Unix()),             // 更新时间
			MemAlias:      alias,                              // 昵称
			MemIntro:      memIntro,                           // 会员介绍
			MemArticles:   0,                                  // 文章数量
			MemPages:      0,                                  // 页数
			MemComments:   0,                                  // 评论数
			MemUploads:    0,                                  // 上传数
			MemTemplate:   "",                                 // 模板
			MemMeta:       "",                                 // Meta 信息（可以是JSON）
		}
		err = z.DB.Create(&newMember).Error
		if err != nil {
			err = errors.New(z.WebSite + "设置用户 失败" + err.Error())
			return
		}
	}

	return
}
