package multitask

import (
	"github.com/SVz777/tk/collections"
)

type ITaskPool interface {
	Do(ITask)
	Serve()
	Stop()
	Done() chan collections.Empty
}

func NewTaskPool(opt ...Option) ITaskPool {
	opts := GetDefaultOptions()
	for _, o := range opt {
		o(opts)
	}
	if opts.HashFunc == nil {
		return newNormalTaskPool(opts)
	}
	return newHashTaskPool(opts)
}

// normalTaskPool 随机分配core执行task
type normalTaskPool struct {
	ch     chan ITask
	core   int
	closed chan collections.Empty
}

func newNormalTaskPool(opts *Options) *normalTaskPool {
	tp := normalTaskPool{
		ch:   make(chan ITask, opts.CoreNum),
		core: opts.CoreNum,
	}
	return &tp
}

func (tp *normalTaskPool) Do(task ITask) {
	tp.ch <- task
}

func (tp *normalTaskPool) Serve() {
	for ii := 0; ii < tp.core; ii++ {
		go start(tp, tp.ch)
	}
	tp.closed = make(chan collections.Empty)
}

func (tp *normalTaskPool) Stop() {
	close(tp.closed)
}

func (tp *normalTaskPool) Done() chan collections.Empty {
	return tp.closed
}

// hashTaskPool 同key task会在同一个core上执行，保证有序
type hashTaskPool struct {
	chs    []chan ITask
	core   int
	hash   func(string) int
	closed chan collections.Empty
}

func newHashTaskPool(opts *Options) *hashTaskPool {
	tp := hashTaskPool{
		chs:  make([]chan ITask, opts.CoreNum),
		core: opts.CoreNum,
		hash: opts.HashFunc,
	}
	for idx := range tp.chs {
		tp.chs[idx] = make(chan ITask, opts.CoreNum)
	}
	return &tp
}

func (tp *hashTaskPool) Do(task ITask) {
	tp.chs[tp.hash(task.GetKey())%tp.core] <- task
}

func (tp *hashTaskPool) Serve() {
	for ii := 0; ii < tp.core; ii++ {
		go start(tp, tp.chs[ii])
	}
	tp.closed = make(chan collections.Empty)
}

func (tp *hashTaskPool) Stop() {
	close(tp.closed)
}

func (tp *hashTaskPool) Done() chan collections.Empty {
	return tp.closed
}

func start(pool ITaskPool, c chan ITask) {
	for {
		select {
		case task := <-c:
			select {
			case <-task.Context().Done():
				// context 结束了不执行task
				continue
			default:
				task.Do()
			}
		case <-pool.Done():
			return
		}
	}
}
