package convert

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Convert(v interface{}, kind reflect.Kind) (interface{}, error) {
	switch kind {
	case reflect.Int:
		return Int(v)
	case reflect.Int32:
		return Int32(v)
	case reflect.Int64:
		return Int64(v)
	case reflect.Uint:
		return Uint(v)
	case reflect.Uint64:
		return Uint64(v)
	case reflect.String:
		return String(v)
	case reflect.Float64:
		return Float64(v)
	case reflect.Bool:
		return Bool(v)
	default:
		return nil, fmt.Errorf("not support type:%v", kind)
	}
}

func Int(v interface{}) (int, error) {
	if v == nil {
		return 0, nil
	}
	switch t := v.(type) {
	case int:
		return t, nil
	case int8:
		return int(t), nil
	case int16:
		return int(t), nil
	case int32:
		return int(t), nil
	case int64:
		return int(t), nil
	case uint:
		return int(t), nil
	case uint8:
		return int(t), nil
	case uint16:
		return int(t), nil
	case uint32:
		return int(t), nil
	case uint64:
		return int(t), nil
	case float32:
		return int(t), nil
	case float64:
		return int(t), nil
	case string:
		if len(t) == 0 {
			return 0, nil
		}
		i, err := strconv.Atoi(t)
		if err != nil {
			return 0, fmt.Errorf("string to int err:%w", err)
		}
		return i, nil
	case json.Number:
		i, err := t.Int64()
		if err != nil {
			return 0, fmt.Errorf("json number to int %w", err)
		}
		return int(i), nil
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("not support type %#v", v)
	}
}

func Int32(v interface{}) (int32, error) {
	if v == nil {
		return 0, nil
	}
	switch t := v.(type) {
	case int:
		return int32(t), nil
	case int8:
		return int32(t), nil
	case int16:
		return int32(t), nil
	case int32:
		return t, nil
	case int64:
		return int32(t), nil
	case uint:
		return int32(t), nil
	case uint8:
		return int32(t), nil
	case uint16:
		return int32(t), nil
	case uint32:
		return int32(t), nil
	case uint64:
		return int32(t), nil
	case float32:
		return int32(t), nil
	case float64:
		return int32(t), nil
	case string:
		if len(t) == 0 {
			return 0, nil
		}
		i, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("string to int32 err:%w", err)
		}
		return int32(i), nil
	case json.Number:
		i, err := t.Int64()
		if err != nil {
			return 0, fmt.Errorf("json number to int32 %w", err)
		}
		return int32(i), nil
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("not support type %#v", v)
	}
}

func Int64(v interface{}) (int64, error) {
	if v == nil {
		return 0, nil
	}
	switch t := v.(type) {
	case int:
		return int64(t), nil
	case int8:
		return int64(t), nil
	case int16:
		return int64(t), nil
	case int32:
		return int64(t), nil
	case int64:
		return t, nil
	case uint:
		return int64(t), nil
	case uint8:
		return int64(t), nil
	case uint16:
		return int64(t), nil
	case uint32:
		return int64(t), nil
	case uint64:
		return int64(t), nil
	case float32:
		return int64(t), nil
	case float64:
		return int64(t), nil
	case string:
		if len(t) == 0 {
			return 0, nil
		}
		i, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("string to int64 err:%w", err)
		}
		return i, nil

	case json.Number:
		i, err := t.Int64()
		if err != nil {
			return 0, fmt.Errorf("json number to int64 %w", err)
		}
		return i, nil

	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("not support type %#v", v)
	}
}

func Uint(v interface{}) (uint, error) {
	if v == nil {
		return 0, nil
	}
	switch t := v.(type) {
	case int:
		return uint(t), nil
	case int8:
		return uint(t), nil
	case int16:
		return uint(t), nil
	case int32:
		return uint(t), nil
	case int64:
		return uint(t), nil
	case uint:
		return t, nil
	case uint8:
		return uint(t), nil
	case uint16:
		return uint(t), nil
	case uint32:
		return uint(t), nil
	case uint64:
		return uint(t), nil
	case float32:
		return uint(t), nil
	case float64:
		return uint(t), nil
	case string:
		if len(t) == 0 {
			return 0, nil
		}
		i, err := strconv.Atoi(t)
		if err != nil {
			return 0, fmt.Errorf("string to uint err:%w", err)
		}
		return uint(i), nil
	case json.Number:
		i, err := t.Int64()
		if err != nil {
			return 0, fmt.Errorf("json number to uint %w", err)
		}
		return uint(i), nil
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("not support type %#v", v)
	}
}

func Uint64(v interface{}) (uint64, error) {
	if v == nil {
		return 0, nil
	}
	switch t := v.(type) {
	case int:
		return uint64(t), nil
	case int8:
		return uint64(t), nil
	case int16:
		return uint64(t), nil
	case int32:
		return uint64(t), nil
	case int64:
		return uint64(t), nil
	case uint:
		return uint64(t), nil
	case uint8:
		return uint64(t), nil
	case uint16:
		return uint64(t), nil
	case uint32:
		return uint64(t), nil
	case uint64:
		return t, nil
	case float32:
		return uint64(t), nil
	case float64:
		return uint64(t), nil
	case string:
		if len(t) == 0 {
			return 0, nil
		}
		i, err := strconv.ParseUint(t, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("string to uint64 err:%w", err)
		}
		return i, nil

	case json.Number:
		i, err := t.Int64()
		if err != nil {
			return 0, fmt.Errorf("json number to uint64 %w", err)
		}
		return uint64(i), nil

	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("not support type %#v", v)
	}
}

func String(v interface{}) (string, error) {
	if v == nil {
		return "", nil
	}
	switch t := v.(type) {
	case int:
		return strconv.Itoa(t), nil
	case int8:
		return strconv.FormatInt(int64(t), 10), nil
	case int16:
		return strconv.FormatInt(int64(t), 10), nil
	case int32:
		return strconv.FormatInt(int64(t), 10), nil
	case int64:
		return strconv.FormatInt(t, 10), nil
	case uint:
		return strconv.FormatUint(uint64(t), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(t), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(t), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(t), 10), nil
	case uint64:
		return strconv.FormatUint(t, 10), nil
	case float32:
		return strconv.FormatFloat(float64(t), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64), nil
	case string:
		return t, nil
	case json.Number:
		return t.String(), nil
	case bool:
		if t {
			return "true", nil
		}
		return "false", nil
	case map[string]interface{}, []interface{}:
		if b, err := json.Marshal(t); err != nil {
			return "", fmt.Errorf("map to string %w", err)
		} else {
			return string(b), nil
		}
	default:
		return "", fmt.Errorf("not support type %#v", v)
	}
}

func Float64(v interface{}) (float64, error) {
	if v == nil {
		return 0, nil
	}
	switch t := v.(type) {
	case int:
		return float64(t), nil
	case int8:
		return float64(t), nil
	case int16:
		return float64(t), nil
	case int32:
		return float64(t), nil
	case int64:
		return float64(t), nil
	case uint:
		return float64(t), nil
	case uint8:
		return float64(t), nil
	case uint16:
		return float64(t), nil
	case uint32:
		return float64(t), nil
	case uint64:
		return float64(t), nil
	case float32:
		return float64(t), nil
	case float64:
		return t, nil
	case string:
		if len(t) == 0 {
			return 0, nil
		}
		f, err := strconv.ParseFloat(t, 64)
		if err != nil {
			return 0, fmt.Errorf("string to float64 err:%w", err)
		}
		return f, nil

	case json.Number:
		f, err := t.Float64()
		if err != nil {
			return 0, fmt.Errorf("string to float64 err:%w", err)
		}
		return f, nil

	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("not support type %#v", v)
	}
}

func Bool(v interface{}) (bool, error) {
	if v == nil {
		return false, nil
	}
	switch t := v.(type) {
	case int:
		return t != 0, nil
	case int8:
		return t != 0, nil
	case int16:
		return t != 0, nil
	case int32:
		return t != 0, nil
	case int64:
		return t != 0, nil
	case uint:
		return t != 0, nil
	case uint8:
		return t != 0, nil
	case uint16:
		return t != 0, nil
	case uint32:
		return t != 0, nil
	case uint64:
		return t != 0, nil
	case float32:
		return t != 0, nil
	case float64:
		return t != 0, nil
	case string:
		if len(t) == 0 {
			return false, nil
		}
		return strings.ToLower(t) == "true", nil

	case json.Number:
		f, err := t.Float64()
		if err != nil {
			return false, fmt.Errorf("string to float64 err:%w", err)
		}
		return f != 0, nil
	case bool:
		return t, nil
	default:
		return false, fmt.Errorf("not support type %#v", v)
	}
}
