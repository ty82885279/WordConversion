package word

import (
	"strings"
	"unicode"
)

//转换成大写
func ToUppper(s string) string {
	return strings.ToUpper(s)
}

//转换成小些
func ToLower(s string) string {

	return strings.ToLower(s)
}

//下划线单词转大写驼峰单词
func UnderscoreToUpperCamelCase(s string) string {

	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

//下划线单词转小写驼峰单词
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

//驼峰转下划线
func CamelCaseToUnderscore(s string) string {

	var output []rune
	for i, v := range s {

		if i == 0 {
			output = append(output, unicode.ToLower(v))
			continue
		}
		if unicode.IsUpper(v) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(v))
	}
	return string(output)
}
