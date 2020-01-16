package database

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

/*
	Returns a map for the given model, where each key is the JSON field name and
	each value is a string representation of that field's type.

	For example, a struct which holds a string Name value has types["name"] = "string"
*/
func GetFieldTypes(model interface{}) map[string]string {
	expected_types := make(map[string]string)

	v := reflect.ValueOf(model).Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		json_name := strings.ToLower(field.Tag.Get("json"))
		expected_types[json_name] = field.Type.Name()
	}

	return expected_types
}

/*
	Turns a series of string URL parameters into a query for a particular data type
	Returns a map of generated QuerySelectors
*/
func CreateFilterQuery(
	parameters map[string][]string, model interface{}) (map[string]interface{}, error) {

	expected_types := GetFieldTypes(model)

	query := make(map[string]interface{})

	for key, values := range parameters {
		if len(values) > 1 {
			return nil, errors.New("Multiple usage of key " + key)
		}

		key = strings.ToLower(key)
		values := strings.Split(values[0], ",")

		to_type, ok := expected_types[key]
		if !ok {
			return nil, errors.New("Invalid key " + key)
		}

		// We must specifically handle each data type
		switch to_type {
		case "string":
			query[key] = QuerySelector{"$in": values}
		case "int64":
			cast_values := make([]int64, len(values))
			for i, value := range values {
				value_int, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, err
				}
				cast_values[i] = value_int
			}
			query[key] = QuerySelector{"$in": cast_values}
		case "bool":
			cast_values := make([]bool, len(values))
			for i, value := range values {
				value_int, err := strconv.ParseBool(value)
				if err != nil {
					return nil, err
				}
				cast_values[i] = value_int
			}
			query[key] = QuerySelector{"$in": cast_values}
		}
	}

	return query, nil
}
