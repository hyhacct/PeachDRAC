package modules

import (
	"PeachDRAC/backend/constants"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Logs struct {
}

// LogLevel 定义日志级别
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

var (
	logFile   *os.File
	logger    *log.Logger
	once      sync.Once
	logMutex  sync.Mutex
	logLevels = map[LogLevel]string{
		DEBUG: "DEBUG",
		INFO:  "INFO",
		WARN:  "WARN",
		ERROR: "ERROR",
	}
)

func NewLogsService() *Logs {
	return &Logs{}
}

// InitLogger 初始化日志系统
func (c *Logs) InitLogger() error {
	var err error
	once.Do(func() {
		// 确保日志目录存在
		logDir := filepath.Dir(constants.PathLog)
		if err = os.MkdirAll(logDir, 0755); err != nil {
			return
		}

		// 打开日志文件，如果不存在则创建，追加写入
		logFile, err = os.OpenFile(constants.PathLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return
		}

		// 初始化logger
		logger = log.New(logFile, "", 0)
	})
	return err
}

// LogMessage 记录日志消息
func (c *Logs) LogMessage(level LogLevel, format string, args ...interface{}) {
	if logger == nil {
		return
	}

	logMutex.Lock()
	defer logMutex.Unlock()

	// 格式化时间
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")

	// 格式化消息
	message := fmt.Sprintf(format, args...)

	// 完整的日志条目
	logEntry := fmt.Sprintf("[%s] [%s] %s", timestamp, logLevels[level], message)

	// 写入日志
	logger.Println(logEntry)
}

// CloseLogger 关闭日志文件
func (c *Logs) CloseLogger() {
	if logFile != nil {
		logFile.Close()
	}
}

// 便捷的日志记录方法
func (c *Logs) Debug(format string, args ...interface{}) {
	c.LogMessage(DEBUG, format, args...)
}

func (c *Logs) Info(format string, args ...interface{}) {
	c.LogMessage(INFO, format, args...)
}

func (c *Logs) Warn(format string, args ...interface{}) {
	c.LogMessage(WARN, format, args...)
}

func (c *Logs) Error(format string, args ...interface{}) {
	c.LogMessage(ERROR, format, args...)
}

// RotateLog 日志文件轮转
func (c *Logs) RotateLog(logPath string) error {
	logMutex.Lock()
	defer logMutex.Unlock()

	// 关闭当前日志文件
	if logFile != nil {
		logFile.Close()
	}

	// 生成新的日志文件名（带时间戳）
	timestamp := time.Now().Format("20060102150405")
	ext := filepath.Ext(logPath)
	newPath := fmt.Sprintf("%s.%s%s", logPath[:len(logPath)-len(ext)], timestamp, ext)

	// 重命名当前日志文件
	if err := os.Rename(logPath, newPath); err != nil {
		return fmt.Errorf("rotate log file failed: %v", err)
	}

	// 创建新的日志文件
	var err error
	logFile, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("create new log file failed: %v", err)
	}

	// 更新logger
	logger = log.New(logFile, "", 0)
	return nil
}
