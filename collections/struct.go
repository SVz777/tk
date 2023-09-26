package collections

import (
	"fmt"
	"reflect"
)

// GetValueWithFieldPath  根据path列表从 data 结构体中获取数据 data为*struct or struct
func GetValueWithFieldPath(data any, path ...string) (any, error) {
	if len(path) <= 0 {
		return nil, fmt.Errorf("path <= 0")
	}
	d := reflect.ValueOf(data)

	for idx, p := range path {
		if d.Kind() == reflect.Interface {
			// interface 转成实际类型
			d = d.Elem()
		}
		if d.Kind() == reflect.Ptr {
			// 指针解引用
			d = d.Elem()
		}
		d = d.FieldByName(p)
		if !d.IsValid() {
			return nil, fmt.Errorf("%v is nil", path[:idx])
		}
	}

	return d.Interface(), nil

}

// SetValueWithFieldPath 根据path列表从 设置 data 结构体中对应数据 data必须为*struct
func SetValueWithFieldPath(data any, value any, path ...string) error {
	if len(path) <= 0 {
		return fmt.Errorf("path <= 0")
	}
	d := reflect.ValueOf(data)
	for idx, p := range path {
		if d.Kind() == reflect.Interface {
			d = d.Elem()
		}
		if d.Kind() == reflect.Ptr {
			d = d.Elem()
		}
		d = d.FieldByName(p)
		if !d.IsValid() {
			return fmt.Errorf("%v is nil", path[:idx])
		}
	}

	// 判断类型是不是一致
	vvr := reflect.ValueOf(value)
	if d.Kind() == reflect.Ptr {
		// 指针取值
		d = d.Elem()
	}
	if d.Kind() != reflect.Interface && d.Kind() != vvr.Kind() {
		return fmt.Errorf("kind not equal d:%s v:%s", d.Kind(), vvr.Kind())
	}
	if !d.CanSet() {
		return fmt.Errorf("%v:%v can't set", path, d.String())
	}
	d.Set(vvr)
	return nil
}
