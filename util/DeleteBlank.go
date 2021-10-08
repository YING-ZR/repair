package util

//删除空格
func DeleteTailBlank(str string) string {
	spaceNum := 0
	for i := len(str)-1; i >= 0; i-- {  // 去除字符串尾部的所有空格
		if str[i] == ' ' {
			spaceNum++
		} else {
			break
		}
	}
	return str[:len(str)-spaceNum]
}
