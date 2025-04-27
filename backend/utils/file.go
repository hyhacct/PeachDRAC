package utils

import (
	"fmt"
	"os"
	"time"
)

// 将文件内容写入到指定路径下,并且等待完全写入完成
func FileWriteToPath(path string, content []byte) error {
	err := os.WriteFile(path, content, 0644)
	if err != nil {
		return fmt.Errorf("写入文件失败: %v", err)
	}
	time.Sleep(time.Second * 1) // 等待1秒,确保文件完全写入
	return nil
}
