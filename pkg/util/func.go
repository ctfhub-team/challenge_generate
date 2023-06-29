package util

import (
	"os"
)

/**
 * @Description: 获取map中所有key
 * @param m map对象
 * @return []string key数组
 */
func GetMapKey(m map[string]string) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

/**
 * @Description: 写入文件
 * @param filename 文件路径
 * @param content 文件内容
 * @param perm 文件权限 例如 0755
 */
func WriteFile(filename string, content string, perm os.FileMode) {
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, perm)
	file.WriteString(content)
}

func ReadFileByte(filename string) ([]byte, error) {
	//获得一个file
	f, err := os.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}
	return f, err
}
func ReadFileString(filename string) (string, error) {
	data, err := ReadFileByte(filename)
	return string(data), err
}
