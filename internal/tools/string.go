package tools

import (
	"strings"
)

// extractFirstWord 返回原始字符串中的第一个单词和剩余部分
func ExtractFirstWord(s string) (string, string) {
	// 使用strings.Fields分割字符串为单词切片
	// 注意：这也会去掉多余的空格，并且如果s只包含空格，将返回一个空切片
	words := strings.Fields(s)

	// 如果没有单词，返回空字符串和原始字符串
	if len(words) == 0 {
		return "", s
	}

	// 第一个单词是words[0]
	firstWord := words[0]

	// 剩余部分是去掉第一个单词后的原始字符串
	// 注意：这里我们使用strings.Join将切片（除了第一个单词）重新组合为字符串
	// 但由于我们只关心第一个单词后的内容，我们实际上可以简单地从原始字符串中切掉第一个单词
	remainder := strings.TrimPrefix(s, firstWord+" ") // 去掉第一个单词和后面的空格（如果有）

	// 如果原始字符串以第一个单词结束（即没有空格或其他字符），则remainder将是空的
	// 为了避免这种情况，我们检查remainder是否和去掉第一个单词的s相同，如果是，则将其设为空字符串
	if strings.HasPrefix(s, firstWord) && strings.TrimSpace(remainder) == "" {
		remainder = ""
	}

	return firstWord, remainder
}

// 将字符串的 \n 替换为\\n " 替换为\"
func ReplaceSpecialChar(s string) string {
	s = strings.Replace(s, "\n", "\\n", -1)
	s = strings.Replace(s, "\"", "\\\"", -1)
	return s
}
