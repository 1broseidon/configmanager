package internal

import (
	"fmt"
	"reflect"
	"strings"
)

// Flatten converts a nested map or struct into a flat map with dot notation keys.
func Flatten(data map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range data {
		flatten(v, k, result)
	}
	return result
}

// Recursive function to flatten nested maps and structs.
func flatten(data interface{}, prefix string, result map[string]interface{}) map[string]interface{} {
	rt := reflect.TypeOf(data)
	rv := reflect.ValueOf(data)

	switch rt.Kind() {
	case reflect.Map:
		for _, key := range rv.MapKeys() {
			strKey := fmt.Sprint(key.Interface())
			newPrefix := strKey
			if prefix != "" {
				newPrefix = prefix + "." + strKey
			}
			flatten(rv.MapIndex(key).Interface(), newPrefix, result)
		}
	case reflect.Struct:
		for i := 0; i < rt.NumField(); i++ {
			field := rt.Field(i)
			newPrefix := field.Name
			if prefix != "" {
				newPrefix = prefix + "." + field.Name
			}
			flatten(rv.Field(i).Interface(), newPrefix, result)
		}
	default:
		if prefix != "" {
			result[prefix] = data
		}
	}
	return result
}

// Unflatten restores a flat map with dot notation keys to a nested map.
func Unflatten(data map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range data {
		keys := strings.Split(k, ".")
		m := result
		for i, key := range keys {
			if i == len(keys)-1 {
				m[key] = v
			} else {
				if _, ok := m[key]; !ok {
					m[key] = make(map[string]interface{})
				}
				m = m[key].(map[string]interface{})
			}
		}
	}
	return result
}
