package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 原始字符串
	input := `
		123123.com
		运行环境 [php74]
		123123.com
		HTTP
		永不过期
		1121aa.com
		运行环境 [php74]
		1121aa.com
		HTTP
		永不过期
		111aa.com
		运行环境 [php74]
		111aa.com
		HTTP
		永不过期
		11aa.com
		运行环境 [php74]
		11aa.com
		HTTP
		永不过期
		bb.com
		运行环境 [php74]
		bb.com
		HTTP
		永不过期
		aa.com
	`

	// 使用正则表达式提取所有域名
	re := regexp.MustCompile(`[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	matches := re.FindAllString(input, -1)

	// 去重处理
	uniqueSites := removeDuplicates(matches)

	// 输出结果
	fmt.Println("Extracted Sites:", uniqueSites)
}

// 去重函数
func removeDuplicates(slice []string) []string {
	seen := make(map[string]struct{})
	result := []string{}
	for _, value := range slice {
		if _, exists := seen[value]; !exists {
			seen[value] = struct{}{}
			result = append(result, value)
		}
	}
	return result
}
