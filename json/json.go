package json

import (
	"github.com/bytedance/sonic"
)

var DefaultConfig = sonic.Config{
	UseNumber: true,
}.Froze()

// Marshal ...
func Marshal(v any) ([]byte, error) {
	return DefaultConfig.Marshal(v)
}

// Unmarshal 处理json float64精度丢失
func Unmarshal(data []byte, v any) error {
	return DefaultConfig.Unmarshal(data, v)
}
