package header

import (
	"errors"
	"net/http"
	"reflect"
	"strings"
)

// 不使用reflect.Kind是因为考虑Time、Duration等特殊类型
const (
	Int      = "int"
	Int8     = "int8"
	Int16    = "int16"
	Int32    = "int32"
	Int64    = "int64"
	Uint     = "uint"
	Uint8    = "uint8"
	Uint16   = "uint16"
	Uint32   = "uint32"
	Uint64   = "uint64"
	Float32  = "float32"
	Float64  = "float64"
	String   = "string"
	Bool     = "bool"
	Time     = "time"
	Duration = "duration"

	IntSlice    = "[]int"
	StringSlice = "[]string"
)

var (
	parseFuncs = map[string]func(string) (interface{}, error){
		Int:      parseInt,
		Int8:     parseInt8,
		Int16:    parseInt16,
		Int32:    parseInt32,
		Int64:    parseInt64,
		Uint:     parseUint,
		Uint8:    parseUint8,
		Uint16:   parseUint16,
		Uint32:   parseUint32,
		Uint64:   parseUint64,
		Bool:     parseBool,
		String:   parseString,
		Time:     parseTime,
		Duration: parseDuration,
	}

	parseSliceFuncs = map[string]func([]string) (interface{}, error){
		StringSlice: parseStringSlice,
		IntSlice:    parseIntSlice,
	}
)

func Parse(h http.Header, target interface{}) error {
	// 判断target是否为指针
	vTg := reflect.ValueOf(target)
	if vTg.Kind() != reflect.Ptr {
		return errors.New("the target is not a pointer to a struct")
	}
	// 判断target的Elem是否为结构体
	vElem := vTg.Elem()
	if vElem.Kind() != reflect.Struct {
		return errors.New("element of the pointer is not a struct")
	}
	// 找到httpheader的tag，解析
	tElem := vElem.Type()
	for i, nf := 0, tElem.NumField(); i < nf; i++ {
		f := tElem.Field(i)
		tagStr := f.Tag.Get("httpheader")
		slice := strings.Split(tagStr, ";")
		var key, valueType string

		if len(slice) == 1 {
			key, valueType = slice[0], String
			// empty tag
			if key == "" {
				key = f.Name
			}
		} else {
			key, valueType = slice[0], slice[1]
		}
		value := h.Values(key)
		if len(value) == 0 {
			continue
		}
		// 假如结构体中的目标字段是个切片，则逐个解析，否则只解析value[0]
		if f.Type.Kind() == reflect.Slice {
			if parseFn, ok := parseSliceFuncs[valueType]; ok {
				if parsedValue, err := parseFn(value); err == nil {
					vElem.Field(i).Set(reflect.ValueOf(parsedValue))
				}
			}
			return nil
		}

		v0 := value[0]
		if parseFn, ok := parseFuncs[valueType]; ok {
			if parsedValue, err := parseFn(v0); err == nil {
				// valueType字符串需与类型一致
				// TODO: handle panic?
				vElem.Field(i).Set(reflect.ValueOf(parsedValue))
			}
		}
	}
	return nil
}
