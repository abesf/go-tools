package string

import "strings"
//fomat 12345678 to 12,345,678
func NumFormat(str string) string {
	numStr := strings.Split(str, ".")[0] //如果有小数获取整数部分
	length := len(str)
	if length < 4 {
		return str
	}
	count := (length - 1) / 3 //取于-有多少组三位数
	for i := 0; i < count; i++ {
		numStr = numStr[:length-(i+1)*3] + "," + numStr[length-(i+1)*3:]
	}
	return numStr
}