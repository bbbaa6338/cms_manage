// Package fake
// @Description: fake package, 用于生成网站需要的数据
package fake

import (
	"fmt"
	"github.com/duke-git/lancet/v2/random"
	"github.com/duke-git/lancet/v2/slice"
	"time"
)

var provinces = []string{
	"京", // 北京市
	"津", // 天津市
	"沪", // 上海市
	"渝", // 重庆市
	"冀", // 河北省
	"晋", // 山西省
	"辽", // 辽宁省
	"吉", // 吉林省
	"黑", // 黑龙江省
	"苏", // 江苏省
	"浙", // 浙江省
	"皖", // 安徽省
	"闽", // 福建省
	"赣", // 江西省
	"鲁", // 山东省
	"豫", // 河南省
	"鄂", // 湖北省
	"湘", // 湖南省
	"粤", // 广东省
	"琼", // 海南省
	"川", // 四川省
	"黔", // 贵州省
	"滇", // 云南省
	"陕", // 陕西省
	"甘", // 甘肃省
	"青", // 青海省
	"台", // 台湾省
	"蒙", // 内蒙古自治区
	"桂", // 广西壮族自治区
	"藏", // 西藏自治区
	"宁", // 宁夏回族自治区
	//"新",  // 新疆维吾尔自治区
	//"港",  // 香港特别行政区
	//"澳",  // 澳门特别行政区
}

type WebFake struct{}

// ICP
//
//	@Description: 生成备案号
//	@receiver w
//	@return string
func (w *WebFake) ICP() string {
	province, _ := slice.Random(provinces)
	year := time.Now().Year()

	rr := random.RandNumeral(random.RandInt(6, 7))

	l := random.RandInt(1, 3)
	hao := ""
	if l > 1 {
		hao = fmt.Sprintf("-%d", random.RandInt(1, 9))
	}
	icp := fmt.Sprintf("%sICP备%d%s号%s", province, year, rr, hao)
	return icp
}

// QQ
//
//	@Description: 生成 QQ 号
//	@receiver w
//	@return string
func (w *WebFake) QQ() string {
	qq := random.RandNumeral(8)
	// 如果 qq 的第一个字符为 0，则添加一个随机数
	if qq[0] == '0' {
		r := random.RandInt(1, 9)
		qq = fmt.Sprintf("%d%s", r, qq)
	}
	return qq
}

// WeiXin
//
//	@Description: 		生成微信号
//	@receiver w
//	@return string
func (w *WebFake) WeiXin() string {
	wx := random.RandNumeralOrLetter(random.RandInt(6, 9))
	return wx
}

// DouYin
//
//	@Description: 生成抖音号
//	@receiver w
//	@return string
func (w *WebFake) DouYin() string {
	dy := random.RandNumeralOrLetter(random.RandInt(6, 9))
	return dy
}

// Weibo
//
//	@Description: 生成微博号
//	@receiver w
//	@return string
func (w *WebFake) Weibo() string {
	wb := random.RandNumeralOrLetter(random.RandInt(6, 9))
	return wb
}

// CopyRight
//
//	@Description: 生成版权信息
//	@receiver w
//	@param str
//	@return cr
func (w *WebFake) CopyRight(str string) (cr string) {
	// 获取当前年份
	currentYear := time.Now().Year()
	// 随机获取前2到4年中的一个年份
	randomYearsAgo := random.RandInt(2, 5) // 随机获取2到4
	randomYear := currentYear - randomYearsAgo
	if str == "" {
		str = "本网站解释权归本站所有"
		cr = fmt.Sprintf(`Copyright<i class="fa fa-copyright"></i>%d-%d<a href="/">%s</a>`,
			currentYear, randomYear, str)
	} else {
		cr = fmt.Sprintf(`Copyright<i class="fa fa-copyright"></i>%d-%d<a href="/">%s</a>版权所有`,
			currentYear, randomYear, str)
	}
	return
}
