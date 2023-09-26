package multitask

import (
	"context"
	"fmt"
	"sync"
)

type ITaskManager interface {
	Add(key string, f Do)
	GetAllTasks() map[string]ITask
	GetTaskResult(key string) (result any)
	GetTaskErr(key string) (err error)
	Do() error
}
type taskManager struct {
	ctx   context.Context
	wg    sync.WaitGroup
	tasks map[string]ITask
}

func NewTaskManager(ctx context.Context) ITaskManager {
	return &taskManager{
		ctx:   ctx,
		wg:    sync.WaitGroup{},
		tasks: make(map[string]ITask),
	}
}

func (tm *taskManager) Add(key string, f Do) {
	tm.tasks[key] = NewTask(tm.ctx, key, f)
}

func (tm *taskManager) GetAllTasks() map[string]ITask {
	return tm.tasks
}

func (tm *taskManager) GetTaskResult(key string) (result any) {
	if v, ok := tm.tasks[key]; ok {
		result, _ = v.GetResult()
	}
	return
}

func (tm *taskManager) GetTaskErr(key string) (err error) {
	if v, ok := tm.tasks[key]; ok {
		_, err = v.GetResult()
	}
	return
}

func (tm *taskManager) Do() error {
	if len(tm.tasks) <= 0 {
		return nil
	}
	tm.wg.Add(len(tm.tasks))
	for key, t := range tm.tasks {
		go func(ctx context.Context, key string, t ITask) {
			defer tm.wg.Done()
			go t.Do()
			select {
			case <-t.Done():
				return
			case <-ctx.Done():
				t.SetResult(nil, fmt.Errorf("ctx err: %w", ctx.Err()))
				return
			}
		}(tm.ctx, key, t)
	}
	tm.wg.Wait()
	var err error
	for key, task := range tm.tasks {
		_, taskErr := task.GetResult()
		if taskErr != nil {
			if err == nil {
				err = fmt.Errorf("[%v:%v]", key, taskErr)
			} else {
				err = fmt.Errorf("[%v:%v]||%w", key, taskErr, err)
			}
		}
	}
	return err
}
