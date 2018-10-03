package ologger

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// Log log your object using `oLog` tag
//   e.g: fieldA string `oLog:"1"`
// logLevel: Log when `LogLevel >= struct oLog`
// indentString: Direct pass to json.MarshalIndent / json.Marshal
// obj: Your object
func Log(logLevel int, indentString string, obj interface{}) string {
	result := ""
	fields := extractOLogField(logLevel, obj)

	if indentString != "" {
		temp, _ := json.MarshalIndent(fields, "", indentString)
		result = string(temp)
	} else {
		temp, _ := json.Marshal(fields)
		result = string(temp)
	}
	return string(result)
}

func extractOLogField(logLevel int, obj interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	var typeSwitchFunc func(t reflect.Type, v reflect.Value)
	typeSwitchFunc = func(t reflect.Type, v reflect.Value) {
		switch t.Kind() {
		case reflect.Ptr:
			typeSwitchFunc(t.Elem(), v.Elem())
		case reflect.Struct:
			result[""] = extractStruct(logLevel, t, v)
		case reflect.Map, reflect.Slice:
			result[""] = extractValue(logLevel, v)
		default:
			result[""] = obj
		}
	}
	typeSwitchFunc(reflect.TypeOf(obj), reflect.ValueOf(obj))

	return result
}

func extractStruct(logLevel int, t reflect.Type, v reflect.Value) map[string]interface{} {
	result := map[string]interface{}{}

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		tag := field.Tag.Get("oLog")
		isLog := isNeedToLog(logLevel, tag)

		if isLog {
			result[field.Name] = extractFieldValue(logLevel, value)
		}
	}

	return result
}

func extractFieldValue(logLevel int, value reflect.Value) interface{} {
	var result interface{}
	switch value.Kind() {
	case reflect.Struct:
		result = extractStruct(logLevel, value.Type(), value)
	default:
		result = extractValue(logLevel, value)
	}
	return result
}

// ref https://play.golang.org/p/C4F1BHXAGNR
func extractValue(logLevel int, v reflect.Value) interface{} {
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		len := v.Len()
		s := make([]interface{}, len)
		for i := 0; i < len; i++ {
			s[i] = extractFieldValue(logLevel, v.Index(i))
		}
		return s
	case reflect.Bool:
		return v.Bool()
	case reflect.Chan:
		return v
	case reflect.Complex128, reflect.Complex64:
		return v.Complex()
	case reflect.Float32, reflect.Float64:
		return v.Float()
	case reflect.Func:
		return v
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8:
		return v.Int()
	case reflect.Interface:
		return v.Interface()
	case reflect.Invalid:
		return nil
	case reflect.Map:
		m := map[string]interface{}{} //FIXME: it will deallocated if type is map[interface{}]interface{}
		for _, k := range v.MapKeys() {
			key := fmt.Sprint(extractFieldValue(logLevel, k))
			value := extractFieldValue(logLevel, v.MapIndex(k))
			m[key] = value
		}
		return m
	case reflect.Ptr:
		return extractFieldValue(logLevel, v.Elem())
	case reflect.String:
		return v.String()
	case reflect.Struct:
		return v
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8, reflect.Uintptr:
		return v.Uint()
	}
	return nil
}

func isNeedToLog(logLevel int, tag string) bool {
	tagLogLevel, err := strconv.Atoi(tag)
	result := (err == nil && logLevel >= tagLogLevel)
	return result
}
