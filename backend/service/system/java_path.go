package system

import (
	"PeachDRAC/backend/model"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// FindJavaInstalls 查找系统中的 Java 安装
func (s *SystemService) FindJavaInstalls() model.ConfigRespond {
	var installs []model.SystemJava
	seen := make(map[string]bool) // 用于去重

	// 1. 检查 JAVA_HOME
	if javaHome := os.Getenv("JAVA_HOME"); javaHome != "" {
		javaPath := filepath.Join(javaHome, "bin", "java")
		if runtime.GOOS == "windows" {
			javaPath += ".exe"
		}
		if version, err := getJavaVersion(javaPath); err == nil && !seen[javaPath] {
			installs = append(installs, model.SystemJava{Path: javaPath, Version: version})
			seen[javaPath] = true
		} else if err != nil {
			log.Printf("警告: 无法获取 JAVA_HOME 中的 Java 版本: %v", err)
		}
	}

	// 2. 检查 PATH 中的 java
	if pathJava, err := exec.LookPath("java"); err == nil {
		if version, err := getJavaVersion(pathJava); err == nil && !seen[pathJava] {
			installs = append(installs, model.SystemJava{Path: pathJava, Version: version})
			seen[pathJava] = true
		} else if err != nil {
			log.Printf("警告: 无法获取 PATH 中的 Java 版本: %v", err)
		}
	}

	// 3. 系统特定方法
	var err error
	switch runtime.GOOS {
	case "windows":
		err = findWindowsJavaInstalls(&installs, seen)
	case "linux", "darwin":
		err = findUnixJavaInstalls(&installs, seen)
	default:
		return model.Error("不支持的操作系统 => " + runtime.GOOS)
	}

	if err != nil {
		log.Printf("警告: 系统特定方法查找 Java 失败: %v", err)
	}

	if len(installs) == 0 {
		return model.Error("未发现有效的 Java 安装")
	}

	return model.Success(installs)
}

// findWindowsJavaInstalls 在 Windows 系统上查找 Java 安装
func findWindowsJavaInstalls(installs *[]model.SystemJava, seen map[string]bool) error {
	cmd := exec.Command("where", "java")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("执行 where 命令失败: %w", err)
	}

	paths := strings.Split(strings.TrimSpace(string(output)), "\n")
	for _, path := range paths {
		path = strings.TrimSpace(path)
		if path == "" {
			continue
		}
		if version, err := getJavaVersion(path); err == nil && !seen[path] {
			*installs = append(*installs, model.SystemJava{Path: path, Version: version})
			seen[path] = true
		} else if err != nil {
			log.Printf("警告: 无法获取 Java 版本 (路径: %s): %v", path, err)
		}
	}
	return nil
}

// findUnixJavaInstalls 在 Unix-like 系统上查找 Java 安装
func findUnixJavaInstalls(installs *[]model.SystemJava, seen map[string]bool) error {
	cmd := exec.Command("which", "-a", "java")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("执行 which 命令失败: %w", err)
	}

	paths := strings.Split(strings.TrimSpace(string(output)), "\n")
	for _, path := range paths {
		path = strings.TrimSpace(path)
		if path == "" {
			continue
		}
		// 解析符号链接到真实路径
		if realPath, err := filepath.EvalSymlinks(path); err == nil {
			path = realPath
		} else {
			log.Printf("警告: 无法解析符号链接 (路径: %s): %v", path, err)
		}
		if version, err := getJavaVersion(path); err == nil && !seen[path] {
			*installs = append(*installs, model.SystemJava{Path: path, Version: version})
			seen[path] = true
		} else if err != nil {
			log.Printf("警告: 无法获取 Java 版本 (路径: %s): %v", path, err)
		}
	}
	return nil
}

// getJavaVersion 获取指定 Java 可执行文件的版本信息
func getJavaVersion(javaPath string) (string, error) {
	cmd := exec.Command(javaPath, "-version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("执行 java -version 失败: %w", err)
	}

	lines := strings.Split(string(output), "\n")
	if len(lines) == 0 {
		return "", fmt.Errorf("无法获取版本信息: 输出为空")
	}

	version := strings.TrimSpace(lines[0])
	if version == "" {
		return "", fmt.Errorf("无法获取版本信息: 版本字符串为空")
	}

	return version, nil
}
