package multitask

type Options struct {
	CoreNum  int
	HashFunc func(string) int
}
type Option func(o *Options)

func GetDefaultOptions() *Options {
	return &Options{
		CoreNum: 4,
	}
}

func WithCoreNum(num int) Option {
	return func(o *Options) {
		o.CoreNum = num
	}
}

func WithHashFunc(f func(string) int) Option {
	return func(o *Options) {
		o.HashFunc = f
	}
}
