package json

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/SVz777/tk/convert"
)

// Path json path的封装
type Path struct {
	opts *Options
	data interface{}
}

// NewJSONPath ...
func NewJSONPath(jsonData []byte, opt ...Option) (*Path, error) {
	opts := GetOptions(opt...)
	j := &Path{
		opts: opts,
	}
	err := j.UnmarshalJSON(jsonData)
	if err != nil {
		return nil, err
	}
	return j, nil
}

// NewJSONPathWithData ...
func NewJSONPathWithData(data interface{}, opt ...Option) *Path {
	return &Path{
		opts: GetOptions(opt...),
		data: data,
	}
}

// MarshalJSON json.Marshaler
func (j *Path) MarshalJSON() ([]byte, error) {
	return Marshal(&j.data)
}

// UnmarshalJSON json.Unmarshaler
func (j *Path) UnmarshalJSON(p []byte) error {
	return Unmarshal(p, &j.data)
}

// IsNil 值是否为空
func (j *Path) IsNil() bool {
	return j.data == nil
}

// genNew 生成一份新的只修改 data
func (j *Path) genNew(data interface{}) *Path {
	return &Path{
		opts: j.opts,
		data: data,
	}
}

// Get 获取key 对应值，不存在 IsNil 为 true
func (j *Path) Get(key interface{}) *Path {
	v, _ := j.Get2(key)
	return v
}

// Get2 获取key 对应值，取到了第二个值为true
func (j *Path) Get2(key interface{}) (*Path, bool) {
	if j.opts.ReflectSwitch {
		return j.get2Reflect(key)
	} else {
		return j.get2Comma(key)
	}
}

func (j *Path) get2Comma(key interface{}) (*Path, bool) {
	switch data := j.data.(type) {
	case map[string]interface{}:
		k, ok := key.(string)
		if !ok {
			return j.genNew(nil), false
		}
		v, ok := data[k]
		if !ok {
			return j.genNew(nil), false
		}
		return j.genNew(v), true
	case []interface{}:
		k, err := convert.Int(key)
		if err != nil {
			return j.genNew(nil), false
		}
		if len(data) <= k {
			return j.genNew(nil), false
		}
		return j.genNew(data[k]), true
	default:
		return j.genNew(nil), false
	}
}

func (j *Path) get2Reflect(key interface{}) (*Path, bool) {
	rv := reflect.ValueOf(j.data)
	switch rv.Kind() {
	case reflect.Map:
		k := reflect.ValueOf(key)
		if rv.Type().Key() != k.Type() {
			return j.genNew(nil), false
		}
		v := rv.MapIndex(k)
		if !v.IsValid() {
			return j.genNew(nil), false
		}
		return j.genNew(v.Interface()), true
	case reflect.Slice, reflect.Array:
		k, err := convert.Int(key)
		if err != nil {
			return j.genNew(nil), false
		}
		v := rv.Index(k)
		if !v.IsValid() {
			return j.genNew(nil), false
		}
		return j.genNew(v.Interface()), true
	default:
		return j.genNew(nil), false
	}
}

// GetPath 根据path 获取
func (j *Path) GetPath(path ...string) *Path {
	t := j
	for _, p := range path {
		t = t.Get(p)
	}
	return t
}

// Set 设置值
func (j *Path) Set(key interface{}, value interface{}) bool {
	if j.opts.ReflectSwitch {
		return j.setReflect(key, value)
	} else {
		return j.setComma(key, value)
	}
}

func (j *Path) setComma(key interface{}, value interface{}) bool {
	switch data := j.data.(type) {
	case map[string]interface{}:
		k, ok := key.(string)
		if !ok {
			return false
		}
		data[k] = value
		return true
	case []interface{}:
		k, err := convert.Int(key)
		if err != nil {
			return false
		}
		if k >= len(data) {
			return false
		}
		data[k] = value
		return true
	default:
		return false
	}
}

func (j *Path) setReflect(key interface{}, value interface{}) bool {
	rv := reflect.ValueOf(j.data)
	rt := rv.Type()
	switch rv.Kind() {
	case reflect.Map:
		k := reflect.ValueOf(key)
		if rt.Key() != k.Type() {
			return false
		}
		v := reflect.ValueOf(value)
		if !v.CanConvert(rt.Elem()) {
			return false
		}
		rv.SetMapIndex(k, v)
		return true
	case reflect.Slice, reflect.Array:
		k, err := convert.Int(key)
		if err != nil {
			return false
		}
		if k >= rv.Len() {
			return false
		}
		v := reflect.ValueOf(value)
		if !v.CanConvert(rt.Elem()) {
			return false
		}
		rv.Index(k).Set(v)
		return false
	default:
		return false
	}
}

// SetPath 根据path设置值
func (j *Path) SetPath(path []interface{}, value interface{}) bool {
	if len(path) <= 0 {
		return false
	}
	t := j
	for _, p := range path[:len(path)-1] {
		t = t.Get(p)
	}
	return t.Set(path[len(path)-1], value)
}

// Interface 获取data值
func (j *Path) Interface() interface{} {
	return j.data
}

// Int 将值转为 int
func (j *Path) Int() (int, error) {
	return convert.Int(j.data)
}

// Int32 将值转为 int32
func (j *Path) Int32() (int32, error) {
	return convert.Int32(j.data)
}

// Int64 将值转为 int64
func (j *Path) Int64() (int64, error) {
	return convert.Int64(j.data)
}

// Uint 将值转为 uint
func (j *Path) Uint() (uint, error) {
	return convert.Uint(j.data)
}

// UInt64 将值转为 Uint64
func (j *Path) UInt64() (uint64, error) {
	return convert.Uint64(j.data)
}

// String 将值转为 string
func (j *Path) String() (string, error) {
	return convert.String(j.data)
}

// Float64 将值转为 float64
func (j *Path) Float64() (float64, error) {
	return convert.Float64(j.data)
}

// Bool 将值转为 bool
func (j *Path) Bool() (bool, error) {
	return convert.Bool(j.data)
}

// Map 将值断言为 map[string]interface
func (j *Path) Map() (map[string]interface{}, error) {
	if m, ok := (j.data).(map[string]interface{}); ok {
		return m, nil
	}
	return nil, fmt.Errorf("type assertion to map[string]interface{} failed")
}

// Array 将值断言为 []interface
func (j *Path) Array() ([]interface{}, error) {
	if a, ok := (j.data).([]interface{}); ok {
		return a, nil
	}
	return nil, fmt.Errorf("type assertion to []interface{} failed")
}

// StringArray 将值转为 []string
func (j *Path) StringArray() ([]string, error) {
	a, err := j.Array()
	if err != nil {
		return nil, err
	}
	ret := make([]string, 0, len(a))
	for idx, s := range a {
		v, err := convert.String(s)
		if err != nil {
			return nil, fmt.Errorf("%v:%v convert err: %w", idx, s, err)
		}
		ret = append(ret, v)
	}
	return ret, nil
}

// ParseWithJSONPath 根据data的tag(json_path)定义来填充data
func (j *Path) ParseWithJSONPath(data interface{}) error {
	vr := reflect.ValueOf(data)
	if vr.Kind() != reflect.Ptr {
		// data必须是指针
		return fmt.Errorf("data must be pointer")
	}
	vr = vr.Elem()
	if vr.Kind() != reflect.Struct {
		// *data必须是结构体
		return fmt.Errorf("data must be *struct")
	}
	return j.parseWithJSONPath(vr)
}

func (j *Path) parseWithJSONPath(v reflect.Value) error {
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}
	fn := v.NumField()
	for ii := 0; ii < fn; ii++ {
		vf := v.Field(ii)
		tf := v.Type().Field(ii)
		tagSrc := tf.Tag.Get(j.opts.Tag)
		if tagSrc == "" {
			if tf.Type.Kind() == reflect.Struct ||
				(vf.Kind() == reflect.Ptr && vf.Type().Elem().Kind() == reflect.Struct) {
				err := j.parseWithJSONPath(vf)
				if err != nil {
					return fmt.Errorf("sub struct %s parse err: %w", tf.Name, err)
				}
			}
			continue
		}
		if !vf.CanSet() {
			return fmt.Errorf("%s can't set", tf.Name)
		}
		tagName := strings.Split(tagSrc, ",")[0]
		tagPath := strings.Split(tagName, ".")
		jsonValue := j.GetPath(tagPath...)
		if jsonValue.IsNil() {
			continue
		}
		value, err := j.parseValue(tf.Name, vf.Type(), jsonValue)
		if err != nil {
			if j.opts.IgnoreSingleFieldError {
				continue
			}
			return fmt.Errorf("getvalue error: %w", err)
		}

		if !value.Type().AssignableTo(vf.Type()) {
			return fmt.Errorf("%s can't %s:%v set", tf.Name, value.String(), value)
		}
		vf.Set(value)
	}
	return nil
}

func (j *Path) parseValue(fieldName string, tf reflect.Type, jsonValue *Path) (reflect.Value, error) {
	switch tf.Kind() {
	case reflect.Ptr:
		pv, err := j.parseValue(fieldName, tf.Elem(), jsonValue)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("%s parse slice err: %w", fieldName, err)
		}
		trueValue := reflect.New(tf.Elem())
		trueValue.Elem().Set(pv)
		return trueValue, nil

	case reflect.Interface:
		return reflect.ValueOf(jsonValue.Interface()), nil

	case reflect.Map:
		itemKind := tf.Elem().Kind()
		if itemKind == reflect.Interface {
			v, err := jsonValue.Map()
			if err != nil {
				return reflect.Value{}, fmt.Errorf("%s parse map err: %w", fieldName, err)
			}
			return reflect.ValueOf(v), nil
		} else {
			mValue, err := jsonValue.Map()
			if err != nil {
				return reflect.Value{}, fmt.Errorf("%s parse map err: %w", fieldName, err)
			}
			trueValue := reflect.MakeMapWithSize(tf, len(mValue))
			for k := range mValue {
				iv, err := j.parseValue(fieldName, tf.Elem(), jsonValue.Get(k))
				if err != nil {
					return reflect.Value{}, fmt.Errorf("%s parse map err: %w", fieldName, err)
				}
				trueValue.SetMapIndex(reflect.ValueOf(k), iv)
			}
			return trueValue, nil
		}

	case reflect.Struct:
		trueValue := reflect.New(tf)
		b, _ := jsonValue.MarshalJSON()
		err := json.Unmarshal(b, trueValue.Interface())
		if err != nil {
			return reflect.Value{}, fmt.Errorf("%s parse struct err: %w", fieldName, err)
		}
		return trueValue.Elem(), nil

	case reflect.Slice:
		itemKind := tf.Elem().Kind()
		if itemKind == reflect.Interface {
			aValue, err := jsonValue.Array()
			if err != nil {
				return reflect.Value{}, fmt.Errorf("%s parse slice err: %w", fieldName, err)
			}
			return reflect.ValueOf(aValue), nil
		} else if itemKind == reflect.String {
			aValue, err := jsonValue.StringArray()
			if err != nil {
				return reflect.Value{}, fmt.Errorf("%s parse slice err: %w", fieldName, err)
			}
			return reflect.ValueOf(aValue), nil
		} else {
			aValue, err := jsonValue.Array()
			if err != nil {
				return reflect.Value{}, fmt.Errorf("%s parse slice err: %w", fieldName, err)
			}
			trueValue := reflect.MakeSlice(tf, len(aValue), len(aValue))
			for idx := range aValue {
				iv, err := j.parseValue(fieldName, tf.Elem(), jsonValue.Get(idx))
				if err != nil {
					return reflect.Value{}, fmt.Errorf("%s parse slice err: %w", fieldName, err)
				}
				trueValue.Index(idx).Set(iv)
			}
			return trueValue, nil
		}

	default:
		if j.opts.Convert {
			iv, err1 := convert.Convert(jsonValue.Interface(), tf.Kind())
			if err1 != nil {
				return reflect.Value{}, fmt.Errorf("%s parse default err: %w", fieldName, err1)
			}
			return reflect.ValueOf(iv), nil
		}
		return reflect.ValueOf(jsonValue.Interface()), nil
	}
}
