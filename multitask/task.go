package multitask

import (
	"context"
	"log"
	"sync"

	"github.com/SVz777/tk/collections"
)

type Do func(context.Context) (any, error)

type ITask interface {
	// Do 执行task
	Do()
	// Context 获取 task 的 context
	Context() context.Context
	// GetKey 获取 task key
	GetKey() string
	// SetResult 设置 task 结果
	SetResult(result any, err error)
	// GetResult 获取 task 结果
	GetResult() (any, error)
	// Done 返回一个完成标记的 chan
	Done() <-chan collections.Empty
}

type Task struct {
	ctx context.Context
	key string
	f   Do

	result any
	err    error
	done   chan collections.Empty
	once   sync.Once
}

func NewTask(ctx context.Context, key string, f Do) *Task {
	return &Task{ctx: ctx, key: key, f: f, done: make(chan collections.Empty, 1)}
}

func (task *Task) Context() context.Context {
	return task.ctx
}
func (task *Task) Do() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("task: %v err: %v\n", task.GetKey(), err)
		}
	}()
	task.SetResult(task.f(task.ctx))
	task.end()
}

func (task *Task) GetKey() string {
	return task.key
}

func (task *Task) GetResult() (any, error) {
	return task.result, task.err
}

func (task *Task) Done() <-chan collections.Empty {
	return task.done
}

func (task *Task) SetResult(result any, err error) {
	task.once.Do(func() {
		if err != nil {
			task.err = err
		} else {
			task.result, task.err = result, err
		}
	})
}

func (task *Task) end() {
	close(task.done)
}
