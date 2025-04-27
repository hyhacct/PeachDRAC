package utils

import (
	"fmt"
	"os"
	"path"
)

// 检查并且创建 jnlp 启动目录,并响应绝对路径
func MkdirJnlp() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("取运行目录失败 %v", err)
	}
	dir_path := path.Join(dir, "etc", "jnlp")
	if _, err := os.Stat(dir_path); os.IsNotExist(err) {
		os.MkdirAll(dir_path, 0755)
	}
	return dir_path, nil
}
