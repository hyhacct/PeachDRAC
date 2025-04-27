package utils

import "github.com/google/uuid"

// 生成一串随机的 suid
func RandomSuid() string {
	return uuid.New().String()
}
