package filewatcher

type Op uint32

func (op Op) String() string {
	switch op {
	case Create:
		return "Create"
	case Modify:
		return "Modify"
	case Delete:
		return "Delete"
	default:
		return ""
	}
}

const (
	Create Op = 1 << iota
	Modify
	Delete
)

type CallFunc func(file string, opType Op) error

type IWatcher interface {
	Run()
	Stop()
	Watch()
	Path() string
}
