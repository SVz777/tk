package filewatcher

import (
	"sync"
	"time"

	"github.com/SVz777/tk/collections"
)

type Watcher struct {
	sync.Mutex
	Watchers map[string]IWatcher
	closed   chan collections.Empty
	opts     *Options
}

func NewWatcher(opt ...Option) *Watcher {
	opts := GetDefaultOptions()
	opts.Update(opt...)
	return &Watcher{
		Watchers: make(map[string]IWatcher),
		closed:   make(chan collections.Empty),
		opts:     opts,
	}
}

func (w *Watcher) Run() {
	tc := time.NewTicker(w.opts.ScanInterval)
	defer tc.Stop()
	for {
		select {
		case <-w.closed:
			return
		case <-tc.C:
			for _, tw := range w.Watchers {
				tw.Watch()
			}
		}
	}

}

func (w *Watcher) AddWatcher(iw IWatcher) {
	w.Lock()
	w.Watchers[iw.Path()] = iw
	w.Unlock()
}

func (w *Watcher) AddFileWatcher(path string, callback CallFunc) error {
	fw, err := NewFileWatcher(path, callback)
	if err != nil {
		return err
	}
	w.AddWatcher(fw)
	return nil
}

func (w *Watcher) RemoveWatcher(path string) {
	w.Lock()
	delete(w.Watchers, path)
	w.Unlock()
}

func (w *Watcher) AddDirWatcher(path string, callback CallFunc) error {
	dw, err := NewDirWatcher(path, callback)
	if err != nil {
		return err
	}
	w.AddWatcher(dw)
	return nil
}

func (w *Watcher) Stop() {
	close(w.closed)
}
