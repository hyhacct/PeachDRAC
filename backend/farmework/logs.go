package farmework

import (
	"PeachDRAC/backend/constants"
	"log"
	"os"
	"path/filepath"
	"sync"
)

// Logger 封装标准库的 log.Logger，提供线程安全的日志记录
type Logger struct {
	infoLogger    *log.Logger
	successLogger *log.Logger
	errorLogger   *log.Logger
	file          *os.File
	mutex         sync.Mutex
}

var (
	ModuleLogs *Logger
)

// NewLogger 创建并初始化 Logger，日志输出到指定文件
func NewLogger() {
	logFilePath := constants.PathDefaultLog
	// 确保日志目录存在
	dir := filepath.Dir(logFilePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}

	// 打开日志文件（追加模式，若不存在则创建）
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	// 创建 Logger 实例
	logger := &Logger{
		file: file,
	}

	// 初始化不同级别的 Logger，带时间戳和级别前缀
	logger.infoLogger = log.New(file, "Level=INFO ", log.Ldate|log.Ltime)
	logger.successLogger = log.New(file, "Level=SUCCESS ", log.Ldate|log.Ltime)
	logger.errorLogger = log.New(file, "Level=ERROR ", log.Ldate|log.Ltime)

	ModuleLogs = logger
}

// Info 记录 INFO 级别日志，支持多参数
func (l *Logger) Info(args ...interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.infoLogger.Println(args...)
}

// Success 记录 SUCCESS 级别日志，支持多参数
func (l *Logger) Success(args ...interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.successLogger.Println(args...)
}

// Error 记录 ERROR 级别日志，支持多参数
func (l *Logger) Error(args ...interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.errorLogger.Println(args...)
}

// Close 关闭日志文件
func (l *Logger) Close() error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return l.file.Close()
}
