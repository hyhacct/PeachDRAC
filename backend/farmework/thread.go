package farmework

import (
	"context"
	"errors"
	"log"
	"sync"
)

type Task interface {
	Run(ctx context.Context) error
}

type TaskFunc func(ctx context.Context) error

func (f TaskFunc) Run(ctx context.Context) error {
	return f(ctx)
}

type Runner struct {
	wg        sync.WaitGroup
	ctx       context.Context
	cancel    context.CancelFunc
	taskQueue chan Task
	errs      chan error
	forceQuit chan struct{}
}

type RunnerOption func(*Runner)

func WithTaskQueueSize(size int) RunnerOption {
	return func(r *Runner) {
		r.taskQueue = make(chan Task, size)
	}
}

func NewRunner(opts ...RunnerOption) *Runner {
	ctx, cancel := context.WithCancel(context.Background())
	r := &Runner{
		ctx:       ctx,
		cancel:    cancel,
		taskQueue: make(chan Task, 100), // default queue size
		errs:      make(chan error, 100),
		forceQuit: make(chan struct{}),
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// Submit 任务到 Runner
func (r *Runner) Submit(task Task) {
	select {
	case <-r.forceQuit:
		log.Println("task rejected: runner forced to quit")
	default:
		r.wg.Add(1) // 👈 每次提交任务都计数
		r.taskQueue <- task
	}
}

// Run 启动 worker
func (r *Runner) Run(workerCount int) {
	for i := 0; i < workerCount; i++ {
		go func(workerID int) {
			for {
				select {
				case <-r.ctx.Done():
					log.Printf("worker %d: received stop signal\n", workerID)
					return
				case <-r.forceQuit:
					log.Printf("worker %d: forced exit\n", workerID)
					return
				case task := <-r.taskQueue:
					if task == nil {
						continue
					}
					if err := task.Run(r.ctx); err != nil {
						r.errs <- err
					}
					r.wg.Done() // 👈 每个任务执行完毕后减少计数
				}
			}
		}(i)
	}
}

// GracefulStop 发出优雅停止信号
func (r *Runner) GracefulStop() {
	r.cancel()
}

// ForceQuit 强制退出所有任务
func (r *Runner) ForceQuit() {
	close(r.forceQuit)
}

// Wait 等待所有任务完成
func (r *Runner) Wait() error {
	r.wg.Wait()
	close(r.errs)
	if len(r.errs) > 0 {
		return errors.New("some tasks failed")
	}
	return nil
}

// Errors 返回所有错误（如果需要）
func (r *Runner) Errors() []error {
	var result []error
	for err := range r.errs {
		result = append(result, err)
	}
	return result
}
