package filewatcher

import (
	"time"

	"github.com/SVz777/tk/multitask"
)

type Options struct {
	ScanInterval time.Duration
	TpOption     []multitask.Option
}

type Option func(opts *Options)

func (opts *Options) Update(opt ...Option) {
	for _, o := range opt {
		o(opts)
	}
}

func GetDefaultOptions() *Options {
	return &Options{
		ScanInterval: 2 * time.Second,
	}
}

func WithScanInterval(scanInterval time.Duration) Option {
	return func(opts *Options) {
		opts.ScanInterval = scanInterval
	}
}

func WithTpOptions(tpOption ...multitask.Option) Option {
	return func(opts *Options) {
		opts.TpOption = tpOption
	}
}
