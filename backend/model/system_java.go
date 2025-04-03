package model

// SystemJava 表示一个 Java 安装的信息
type SystemJava struct {
	Path    string `json:"path"`    // java 可执行文件的绝对路径
	Version string `json:"version"` // Java 版本
}
