package filewatcher

import (
	"log"
	"os"
	"time"

	"github.com/SVz777/tk/collections"
)

type fileWatcher struct {
	path    string
	call    CallFunc
	modTime time.Time
	closed  chan collections.Empty
	opts    *Options
}

func NewFileWatcher(filepath string, callback CallFunc, opt ...Option) (IWatcher, error) {
	info, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}
	opts := GetDefaultOptions()
	opts.Update(opt...)
	fw := &fileWatcher{
		path:   filepath,
		call:   callback,
		closed: make(chan collections.Empty),
		opts:   opts,
	}
	fw.modTime = info.ModTime()
	return fw, nil
}

func (fw *fileWatcher) Run() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("fw fatal:", err)
		}
	}()
	tc := time.NewTicker(fw.opts.ScanInterval)
	defer tc.Stop()
	// all path watchers
	for {
		select {
		case <-tc.C:
			fw.Watch()
		case <-fw.closed:
			return
		}
	}
}

func (fw *fileWatcher) Path() string {
	return fw.path
}

func (fw *fileWatcher) Watch() {
	info, err := os.Stat(fw.path)
	if err != nil {
		log.Println("fw os stat error:", err)
		return
	}
	if fw.modTime.Before(info.ModTime()) {
		// 修改过
		if err := fw.call(fw.path, Modify); err != nil {
			log.Println("fw callback error:", err)
		}
		fw.modTime = info.ModTime()
	}

}

func (fw *fileWatcher) Stop() {
	close(fw.closed)
}
