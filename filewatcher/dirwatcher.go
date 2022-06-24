package filewatcher

import (
	"io/ioutil"
	"log"
	"sync"
	"time"

	"github.com/SVz777/tk/collections"
)

type dirWatcher struct {
	sync.RWMutex
	path   string
	call   CallFunc
	files  map[string]time.Time
	closed chan collections.Empty
	opts   *Options
}

func NewDirWatcher(dirpath string, callback CallFunc, opt ...Option) (*dirWatcher, error) {
	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return nil, err
	}
	opts := GetDefaultOptions()
	opts.Update(opt...)
	dw := &dirWatcher{
		path:   dirpath,
		call:   callback,
		files:  make(map[string]time.Time, len(files)),
		closed: make(chan collections.Empty),
		opts:   opts,
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		dw.files[file.Name()] = file.ModTime()
	}
	return dw, nil
}

func (dw *dirWatcher) Run() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("dw fatal:", err)
		}
	}()
	tc := time.NewTicker(dw.opts.ScanInterval)
	defer tc.Stop()
	// all path watchers
	for {
		select {
		case <-tc.C:
			dw.Watch()
		case <-dw.closed:
			return
		}
	}
}

func (dw *dirWatcher) Path() string {
	return dw.path
}

func (dw *dirWatcher) Watch() {
	nowfiles, err := ioutil.ReadDir(dw.path)
	if err != nil {
		log.Println("dw read path error:", err)
		return
	}
	beforeFiles := collections.NewSet[string](collections.Keys(dw.files)...)
	dw.RWMutex.Lock()
	for _, file := range nowfiles {
		if file.IsDir() {
			continue
		}
		if modTime, ok := dw.files[file.Name()]; ok {
			// 删除访问过的文件
			beforeFiles.Delete(file.Name())
			if modTime.Before(file.ModTime()) {
				// 修改过
				if err := dw.call(file.Name(), Modify); err != nil {
					log.Println("dw callback error:", err)
					continue
				}
				dw.files[file.Name()] = file.ModTime()
			}
		} else {
			if err := dw.call(file.Name(), Create); err != nil {
				log.Println("dw callback error:", err)
				continue
			}
			dw.files[file.Name()] = file.ModTime()
		}

	}

	for file := range beforeFiles {
		// beforeFiles中都是没访问过的，也就是删除的
		if err := dw.call(file, Delete); err != nil {
			log.Println("dw callback error:", err)
			continue
		}
		delete(dw.files, file)
	}
	dw.RWMutex.Unlock()

}

func (dw *dirWatcher) Stop() {
	close(dw.closed)
}
