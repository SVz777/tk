package json

type Options struct {
	ReflectSwitch bool
	Tag           string
	Convert       bool
	ErrorHandler  func(key string, err error) bool
}

func (opts *Options) Update(opt ...Option) {
	for _, o := range opt {
		o(opts)
	}
}

type Option func(opts *Options)

func GetOptions(opt ...Option) *Options {
	opts := &Options{
		ReflectSwitch: false,
		Tag:           "json_path",
		Convert:       false,
	}
	opts.Update(opt...)
	return opts
}

func WithReflectSwitch(reflectSwitch bool) Option {
	return func(opts *Options) {
		opts.ReflectSwitch = reflectSwitch
	}
}

func WithTag(tag string) Option {
	return func(opts *Options) {
		opts.Tag = tag
	}
}

func WithErrorHandler(f func(string, error) bool) Option {
	return func(opts *Options) {
		opts.ErrorHandler = f
	}
}

func WithConvert(flag bool) Option {
	return func(opts *Options) {
		opts.Convert = flag
	}
}
