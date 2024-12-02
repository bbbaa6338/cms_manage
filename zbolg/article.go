package zbolg

import (
	"cmsManage/utils/reqRequest"
	toPinyin "cmsManage/utils/to_pinyin"
	"fmt"
	"github.com/duke-git/lancet/v2/datetime"
	"github.com/duke-git/lancet/v2/random"
	"github.com/duke-git/lancet/v2/slice"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"
)

func (z *ZBolg) ArticleUpdate(site, cateId, authorId string, title, articleContent, tags, intro string) (err error) {
	tagInfo, err := z.TagSave(tags)
	if err != nil {
		fmt.Println("保存标签失败", err)
	}
	//tagUrl : http://abcdz.cc/tags/aodiaodi/
	replaceTag := make(map[string]string)
	for s, s2 := range tagInfo {
		ht := fmt.Sprintf(`<a href="http://%s/%s/%s/" title="%s">%s</a>`, site, "tags", s2, s, s)
		replaceTag[s] = ht
	}

	articleContent = z.ReplaceInTags(articleContent, replaceTag)

	// 别名
	alias := fmt.Sprintf(random.RandNumeral(13))
	alias = site[:1] + alias

	postTime := time.Now().Format("2006-01-02 15:04:05")

	formData := map[string]string{
		"ID":       "0",
		"Type":     "0",
		"Title":    title,
		"Content":  articleContent,
		"Alias":    alias,
		"Tag":      tags,
		"Intro":    intro,
		"CateID":   cateId,
		"Status":   "0",
		"Template": "single",
		"AuthorID": authorId,
		"PostTime": postTime,
		"IsLock":   "1",
	}

	response, err := z.Session.Post(reqRequest.RequestOption{
		Url:  fmt.Sprintf("http://%s/zb_system/cmd.php?act=ArticlePst&csrfToken=", site) + z.csrfToken,
		Data: formData,
		Headers: map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		},
	})
	if err != nil {
		err = fmt.Errorf(site + " 更新文章失败" + err.Error())
		return
	}

	statusCode := response.StatusCode
	if statusCode != 200 {
		err = fmt.Errorf("更新文章失败, 状态码异常：%d", statusCode)
		return
	}
	return
}

// TagSave
//
//	@Description: 保存标签
//	@receiver z
//	@param tags	string	// 标签
//	@return tagInfo // 标签信息, key 为标签名，value 为标签 别名，因为在设置标签链接的时候，默认为标签别名
//	@return err
func (z *ZBolg) TagSave(tags string) (tagInfo map[string]string, err error) {

	tagSlice := strings.Split(tags, ",")

	newTag := make([]string, len(tagSlice)) // 新增的标签

	alreadyExistsTagAlias := make([]string, len(tagSlice)) // 已经存在的标签别名

	// 查询标签是否存在
	var tag []ZPBTag
	err = z.DB.Where("tag_Name in ?", tagSlice).Find(&tag).Error
	if err != nil {
		return
	}

	tagInfo = make(map[string]string, len(tagSlice))
	for _, zpbTag := range tag {
		alias := zpbTag.TagAlias
		tagName := zpbTag.TagName
		alreadyExistsTagAlias = append(alreadyExistsTagAlias, tagName)
		tagInfo[tagName] = alias
	}

	// 获取新增的标签
	for _, s := range tagSlice {
		if slice.Contain(alreadyExistsTagAlias, s) {
			continue
		}
		newTag = append(newTag, s)
	}

	for _, s := range newTag {
		if len(s) == 0 || utf8.RuneCountInString(s) > 9 {
			continue
		}
		alias := toPinyin.ToPinyin(s)

		var zpbTag ZPBTag
		zpbTag.TagName = s
		zpbTag.TagAlias = alias
		nowTime := int(datetime.Timestamp())
		zpbTag.TagOrder = 0
		zpbTag.TagCount = 0
		zpbTag.TagCreateTime = nowTime
		zpbTag.TagUpdateTime = nowTime
		zpbTag.TagPostTime = nowTime
		err = z.DB.Create(&zpbTag).Error
		if err != nil {
			fmt.Println("创建标签失败", err)
			continue
		}
		tagInfo[s] = alias
	}
	return
}

// ReplaceInTags
//
//	@Description: 替换 content 中 >< 之间的内容，新增 tag 的链接
//	@receiver z
//	@param content
//	@param replacements
//	@return string
func (z *ZBolg) ReplaceInTags(content string, replacements map[string]string) string {
	// 正则表达式：匹配 >< 中的内容
	regex := regexp.MustCompile(`>([^<]+)<`)

	// 排序替换规则：按 key 的长度降序排列，确保长关键词优先
	sortedKeys := make([]string, 0, len(replacements))
	for key := range replacements {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Slice(sortedKeys, func(i, j int) bool {
		return len(sortedKeys[i]) > len(sortedKeys[j]) // 按长度降序排列
	})

	// 替换逻辑：逐段处理 >内容<
	content = regex.ReplaceAllStringFunc(content, func(match string) string {
		// 提取 >< 中的文字内容，去掉 > 和 <
		text := strings.Trim(match, "><")

		// 按标点符号切割文本
		segments := z.splitByPunctuation(text)

		// 遍历每个分段进行替换
		for i, segment := range segments {
			for _, key := range sortedKeys {
				// 如果包含当前 key，进行替换，仅替换一次
				if strings.Contains(segment, key) {
					segments[i] = strings.Replace(segment, key, replacements[key], 1)
					break // 替换成功后跳过其他规则
				}
			}
		}

		// 合并分段，保留原标点符号，重新拼接成字符串
		updatedText := strings.Join(segments, "")

		// 返回重新拼接的字符串
		return fmt.Sprintf(">%s<", updatedText)
	})

	return content
}

// splitByPunctuation 按标点符号切割文本
func (z *ZBolg) splitByPunctuation(text string) []string {
	// 定义常见的标点符号
	punctuation := []string{"。", "，", "！", "？", "；"}
	for _, p := range punctuation {
		text = strings.ReplaceAll(text, p, fmt.Sprintf("%s%s", p, "|||")) // 标点后加分隔符
	}
	return strings.Split(text, "|||") // 按分隔符切割
}
