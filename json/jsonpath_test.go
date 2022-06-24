package json

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJSONPath(t *testing.T) {
	type JS struct {
		Int                int                      `json_path:"int"`
		IntPtr             *int                     `json_path:"int"`
		IntInterface       interface{}              `json_path:"int"`
		Uint               uint                     `json_path:"uint"`
		Float              float64                  `json_path:"float"`
		String             string                   `json_path:"string"`
		StringPtr          *string                  `json_path:"string"`
		StringInterface    interface{}              `json_path:"string"`
		Bool               bool                     `json_path:"bool"`
		IntLevel2          int                      `json_path:"level2.int"`
		UintLevel2         uint                     `json_path:"level2.uint"`
		FloatLevel2        float64                  `json_path:"level2.float"`
		StringLevel2       string                   `json_path:"level2.string"`
		BoolLevel2         bool                     `json_path:"level2.bool"`
		Map                map[string]interface{}   `json_path:"map"`
		StringMap          map[string]string        `json_path:"string_map"`
		IntMap             map[string]int           `json_path:"int_map"`
		MapInterface       interface{}              `json_path:"map"`
		Array              []interface{}            `json_path:"array"`
		MapArray           []map[string]interface{} `json_path:"map_array"`
		MapIntArray        []map[string]int         `json_path:"map_int_array"`
		ArrayInterface     interface{}              `json_path:"array"`
		Array2             [][]interface{}          `json_path:"array2"`
		Array2Interface    interface{}              `json_path:"array2"`
		Array2int          [][]int                  `json_path:"array2"`
		Array2intInterface interface{}              `json_path:"array2"`
		Array3             [][][]interface{}        `json_path:"array3"`
		Array3Interface    interface{}              `json_path:"array3"`
		Array3int          [][][]int                `json_path:"array3"`
		Array3intInterface interface{}              `json_path:"array3"`
		StringArray        []string                 `json_path:"string_array"`
		Struct             struct {
			A int
			B string
			C float64
		} `json_path:"struct"`
		StructPtr *struct {
			A int
			B string
			C float64
		} `json_path:"struct"`
		StructPtr2 **struct {
			A int
			B string
			C float64
		} `json_path:"struct"`
		StructArray []struct {
			A int
			B string
			C float64
		} `json_path:"arraystruct"`
		Struct0 struct {
			Struct0A int     `json_path:"arraystruct.0.a"`
			Struct0B string  `json_path:"arraystruct.0.b"`
			Struct0C float64 `json_path:"arraystruct.0.c"`
		}
	}
	jp, err := NewJSONPath(
		[]byte(`
{
    "int": 1,
    "uint": 1,
    "float": 1.11,
    "string": "jsonpath",
    "bool": true,
    "map": {
        "a": 1,
        "b": "2",
        "c": 3
    },
    "int_map": {
        "a": 1,
        "b": 2,
        "c": 3
    },
    "string_map": {
        "a": "1",
        "b": "2",
        "c": "3"
    },
    "array": [
        1,
        "2",
        3
    ],
    "string_array": [
        "asdf",
        "ghjk",
        "zxcv"
    ],
    "string_array_null": [
        "abc",
        null,
        "efg"
    ],
    "map_array": [
        {
            "map11": 1
        },
        {
            "map21": 1,
            "map22": "2"
        }
    ],
    "map_int_array": [
        {
            "map11": 1
        },
        {
            "map21": 1,
            "map22": 2
        }
    ],
    "array2": [
        [
            1,
            2
        ],
        [
            2,
            3
        ],
        [
            3,
            4
        ]
    ],
    "array3": [
        [
            [
                1,
                2
            ]
        ],
        [
            [
                2,
                3
            ]
        ],
        [
            [
                3,
                4
            ]
        ]
    ],
	"struct": {
            "a": 1,
            "b": "1",
            "c": 1.11
        },
    "arraystruct": [
        {
            "a": 1,
            "b": "1",
            "c": 1.11
        },
        {
            "a": 2,
            "b": "2",
            "c": 2.22
        }
    ],
    "level2": {
        "int": 2,
        "uint": 2,
        "float": 2.22,
        "string": "jsonpath2",
        "bool": false
    }
}`),
	)

	assert.NotEqual(t, nil, jp)
	assert.Equal(t, nil, err)

	var ok bool
	_, ok = jp.Get2("int")
	assert.Equal(t, true, ok)

	_, ok = jp.Get2("missing_key")
	assert.Equal(t, false, ok)

	awm := jp.Get("map_array")
	assert.NotEqual(t, nil, awm)
	var awsval int
	awsval, _ = awm.Get(0).Get("map11").Int()
	assert.Equal(t, 1, awsval)
	awsval, _ = awm.Get(1).Get("map21").Int()
	assert.Equal(t, 1, awsval)
	awsval, _ = awm.Get(1).Get("map22").Int()
	assert.Equal(t, 2, awsval)

	i, _ := jp.Get("int").Int()
	assert.Equal(t, 1, i)

	f, _ := jp.Get("float").Float64()
	assert.Equal(t, 1.11, f)

	s, _ := jp.Get("string").String()
	assert.Equal(t, "jsonpath", s)

	b, _ := jp.Get("bool").Bool()
	assert.Equal(t, true, b)

	strs, err := jp.Get("string_array").StringArray()
	assert.Equal(t, nil, err)
	assert.Equal(t, strs[0], "asdf")
	assert.Equal(t, strs[1], "ghjk")
	assert.Equal(t, strs[2], "zxcv")

	strs2, err := jp.Get("string_array_null").StringArray()
	assert.Equal(t, nil, err)
	assert.Equal(t, strs2[0], "abc")
	assert.Equal(t, strs2[1], "")
	assert.Equal(t, strs2[2], "efg")

	gp, _ := jp.GetPath("level2", "string").String()
	assert.Equal(t, "jsonpath2", gp)

	gp2, _ := jp.GetPath("level2", "int").Int()
	assert.Equal(t, 2, gp2)

	// json struct
	js := JS{}
	err = jp.ParseWithJSONPath(&js)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, js.Int)
	ip := 1
	assert.Equal(t, &ip, js.IntPtr)
	assert.Equal(t, Number("1"), js.IntInterface)
	assert.Equal(t, uint(1), js.Uint)
	assert.Equal(t, 1.11, js.Float)
	assert.Equal(t, "jsonpath", js.String)
	strp := "jsonpath"
	assert.Equal(t, &strp, js.StringPtr)
	assert.Equal(t, "jsonpath", js.StringInterface)
	assert.Equal(t, true, js.Bool)
	assert.Equal(t, 2, js.IntLevel2)
	assert.Equal(t, uint(2), js.UintLevel2)
	assert.Equal(t, 2.22, js.FloatLevel2)
	assert.Equal(t, "jsonpath2", js.StringLevel2)
	assert.Equal(t, false, js.BoolLevel2)
	assert.Equal(t, map[string]interface{}{"a": Number("1"), "b": "2", "c": Number("3")}, js.Map)
	assert.Equal(t, map[string]string{"a": "1", "b": "2", "c": "3"}, js.StringMap)
	assert.Equal(t, map[string]int{"a": 1, "b": 2, "c": 3}, js.IntMap)
	assert.Equal(t, map[string]interface{}{"a": Number("1"), "b": "2", "c": Number("3")}, js.MapInterface)
	assert.Equal(t, []interface{}{Number("1"), "2", Number("3")}, js.Array)
	assert.Equal(
		t, []map[string]interface{}{
			{"map11": Number("1")},
			{"map21": Number("1"), "map22": "2"},
		}, js.MapArray,
	)
	assert.Equal(
		t, []map[string]int{
			{"map11": 1},
			{"map21": 1, "map22": 2},
		}, js.MapIntArray,
	)
	assert.Equal(t, []interface{}{Number("1"), "2", Number("3")}, js.ArrayInterface)
	assert.Equal(
		t,
		[][]interface{}{{Number("1"), Number("2")}, {Number("2"), Number("3")}, {Number("3"), Number("4")}},
		js.Array2,
	)
	assert.Equal(
		t,
		[]interface{}{
			[]interface{}{Number("1"), Number("2")},
			[]interface{}{Number("2"), Number("3")},
			[]interface{}{Number("3"), Number("4")},
		},
		js.Array2Interface,
	)
	assert.Equal(t, [][]int{{1, 2}, {2, 3}, {3, 4}}, js.Array2int)
	assert.Equal(
		t,
		[]interface{}{
			[]interface{}{Number("1"), Number("2")},
			[]interface{}{Number("2"), Number("3")},
			[]interface{}{Number("3"), Number("4")},
		},
		js.Array2intInterface,
	)
	assert.Equal(
		t,
		[][][]interface{}{{{Number("1"), Number("2")}}, {{Number("2"), Number("3")}}, {{Number("3"), Number("4")}}},
		js.Array3,
	)
	assert.Equal(
		t,
		[]interface{}{
			[]interface{}{[]interface{}{Number("1"), Number("2")}},
			[]interface{}{[]interface{}{Number("2"), Number("3")}},
			[]interface{}{[]interface{}{Number("3"), Number("4")}},
		},
		js.Array3Interface,
	)
	assert.Equal(t, [][][]int{{{1, 2}}, {{2, 3}}, {{3, 4}}}, js.Array3int)
	assert.Equal(
		t, []interface{}{
			[]interface{}{[]interface{}{Number("1"), Number("2")}},
			[]interface{}{[]interface{}{Number("2"), Number("3")}},
			[]interface{}{[]interface{}{Number("3"), Number("4")}},
		},
		js.Array3intInterface,
	)
	assert.Equal(t, []string{"asdf", "ghjk", "zxcv"}, js.StringArray)
	assert.Equal(
		t, struct {
			A int
			B string
			C float64
		}{
			A: 1,
			B: "1",
			C: 1.11,
		}, js.Struct,
	)
	assert.Equal(
		t, &struct {
			A int
			B string
			C float64
		}{
			A: 1,
			B: "1",
			C: 1.11,
		}, js.StructPtr,
	)
	sp := &struct {
		A int
		B string
		C float64
	}{
		A: 1,
		B: "1",
		C: 1.11,
	}
	assert.Equal(
		t, &sp, js.StructPtr2,
	)
	assert.Equal(
		t, []struct {
			A int
			B string
			C float64
		}{
			{
				A: 1,
				B: "1",
				C: 1.11,
			}, {
				A: 2,
				B: "2",
				C: 2.22,
			},
		}, js.StructArray,
	)
	assert.Equal(
		t, struct {
			Struct0A int     `json_path:"arraystruct.0.a"`
			Struct0B string  `json_path:"arraystruct.0.b"`
			Struct0C float64 `json_path:"arraystruct.0.c"`
		}{
			Struct0A: 1,
			Struct0B: "1",
			Struct0C: 1.11,
		}, js.Struct0,
	)
}

func TestNewJSONPathWithData(t *testing.T) {
	jp := NewJSONPathWithData(map[string]interface{}{
		"i": 1,
		"m": map[string]interface{}{
			"1": 1,
			"2": "m2",
		},
		"s":  []int{1, 2, 3},
		"si": []interface{}{1, "2", 3},
	}, WithReflectSwitch(true))
	assert.Equal(t, 1, jp.Get("i").Interface())
	assert.Equal(t, map[string]interface{}{
		"1": 1,
		"2": "m2",
	}, jp.Get("m").Interface())
	assert.Equal(t, 1, jp.GetPath("m", "1").Interface())
	assert.Equal(t, "m2", jp.GetPath("m", "2").Interface())
	assert.Equal(t, "m2", jp.Get("m").Get("2").Interface())
	assert.Equal(t, []int{1, 2, 3}, jp.GetPath("s").Interface())
	assert.Equal(t, 2, jp.GetPath("s", "1").Interface())
	assert.Equal(t, []interface{}{1, "2", 3}, jp.Get("si").Interface())
	assert.Equal(t, "2", jp.GetPath("si", "1").Interface())
}

func TestNewJSONPathSet(t *testing.T) {
	jp, err := NewJSONPath([]byte(`{
	"i":1,
	"m":{
		"1": 1,
		"2": "m2"
	},
	"s":[1,2,3],
	"si":[1,"2",3]
}`))
	assert.Equal(t, nil, err)
	assert.Equal(t, Number("1"), jp.Get("i").Interface())
	jp.Set("i", 2)
	assert.Equal(t, 2, jp.Get("i").Interface())
	jp.SetPath([]interface{}{"i"}, 3)
	assert.Equal(t, 3, jp.Get("i").Interface())
	assert.Equal(t, Number("2"), jp.GetPath("s", "1").Interface())
	jp.SetPath([]interface{}{"s", "1"}, 4)
	assert.Equal(t, 4, jp.GetPath("s", "1").Interface())
	assert.Equal(t, "2", jp.GetPath("si", "1").Interface())
	jp.SetPath([]interface{}{"si", "1"}, 4)
	assert.Equal(t, 4, jp.GetPath("si", "1").Interface())
}

func TestNewJSONPathSetFuzz(t *testing.T) {
	jp := NewJSONPathWithData(map[string]interface{}{
		"i": 1,
		"m": map[string]interface{}{
			"1": 1,
			"2": "m2",
		},
		"s":  []int{1, 2, 3},
		"si": []interface{}{1, "2", 3},
	}, WithReflectSwitch(true))
	assert.Equal(t, 1, jp.Get("i").Interface())
	jp.Set("i", 2)
	assert.Equal(t, 2, jp.Get("i").Interface())
	jp.SetPath([]interface{}{"i"}, 3)
	assert.Equal(t, 3, jp.Get("i").Interface())
	assert.Equal(t, 2, jp.GetPath("s", "1").Interface())
	jp.SetPath([]interface{}{"s", "1"}, 4)
	assert.Equal(t, 4, jp.GetPath("s", "1").Interface())
	assert.Equal(t, "2", jp.GetPath("si", "1").Interface())
	jp.SetPath([]interface{}{"si", "1"}, 4)
	assert.Equal(t, 4, jp.GetPath("si", "1").Interface())
}

func BenchmarkJSONPath_getComma(b *testing.B) {
	jp := NewJSONPathWithData(map[string]interface{}{
		"i": 1,
		"m": map[string]interface{}{
			"m1": 1,
			"m2": "m2",
		},
		"s":  []int{1, 2, 3},
		"si": []interface{}{1, "2", 3},
	})
	for i := 0; i < b.N; i++ {
		jp.GetPath("si", "1")
	}
}

func BenchmarkJSONPath_getReflect(b *testing.B) {
	jp := NewJSONPathWithData(map[string]interface{}{
		"i": 1,
		"m": map[string]interface{}{
			"m1": 1,
			"m2": "m2",
		},
		"s":  []int{1, 2, 3},
		"si": []interface{}{1, "2", 3},
	}, WithReflectSwitch(true))
	for i := 0; i < b.N; i++ {
		jp.GetPath("si", "1")
	}
}

func BenchmarkJSONPath_setComma(b *testing.B) {
	jp := NewJSONPathWithData(map[string]interface{}{
		"i": 1,
		"m": map[string]interface{}{
			"m1": 1,
			"m2": "m2",
		},
		"s":  []int{1, 2, 3},
		"si": []interface{}{1, "2", 3},
	})
	for i := 0; i < b.N; i++ {
		jp.SetPath([]interface{}{"si"}, 1)
	}
}

func BenchmarkJSONPath_setReflect(b *testing.B) {
	jp := NewJSONPathWithData(map[string]interface{}{
		"i": 1,
		"m": map[string]interface{}{
			"m1": 1,
			"m2": "m2",
		},
		"s":  []int{1, 2, 3},
		"si": []interface{}{1, "2", 3},
	}, WithReflectSwitch(true))
	for i := 0; i < b.N; i++ {
		jp.SetPath([]interface{}{"si"}, 1)
	}
}
