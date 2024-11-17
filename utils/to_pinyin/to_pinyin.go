package toPinyin

import (
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/mozillazg/go-pinyin"
)

func ToPinyin(text string) string {

	// 配置拼音选项
	args := pinyin.NewArgs()
	args.Style = pinyin.Normal // 输出完整拼音

	// 将文本分离为汉字部分和非汉字部分
	result := ""
	for _, r := range text {
		if r >= '\u4e00' && r <= '\u9fff' { // 判断是否为汉字
			py := pinyin.Pinyin(string(r), args)
			result += py[0][0] + " "
		} else {
			result += string(r) + " " // 非汉字直接拼接
		}
	}
	result = strutil.RemoveWhiteSpace(result, true)
	return result
}
