package database

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

type QueryType string

const (
	LessThan    QueryType = "LessThan"
	In          QueryType = "In"
	GreaterThan QueryType = "GreaterThan"
	NotIn       QueryType = "NotIn"
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

		switch field.Type.Kind() {
		case reflect.Struct: // allows for 1 level of struct embedding
			// TODO: if you want more than 1 level, abstract this out into a local function and add a set backtrace limit (to prevent circular loops)
			for i := 0; i < field.Type.NumField(); i++ {
				tp := field.Type.Field(i)
				json_name := strings.ToLower(tp.Tag.Get("json"))
				expected_types[json_name] = tp.Type.String()
			}
		default:
			json_name := strings.ToLower(field.Tag.Get("json"))
			expected_types[json_name] = field.Type.String()
		}
	}

	return expected_types
}

func UpdateQuerySelectorInt64(qs QuerySelector, query_type QueryType, cast_values []int64) (QuerySelector, error) {
	switch query_type {
	case LessThan:
		qs["$lt"] = cast_values[0]
	case In:
		qs["$in"] = cast_values
	case GreaterThan:
		qs["$gt"] = cast_values[0]
	case NotIn:
		qs["$nin"] = cast_values
	default:
		return nil, errors.New("Invalid operation on integers")
	}
	return qs, nil
}

func UpdateQuerySelectorString(qs QuerySelector, query_type QueryType, cast_values []string) (QuerySelector, error) {
	switch query_type {
	case LessThan:
		qs["$lt"] = cast_values[0]
	case In:
		qs["$in"] = cast_values
	case GreaterThan:
		qs["$gt"] = cast_values[0]
	case NotIn:
		qs["$nin"] = cast_values
	default:
		return nil, errors.New("Invalid operation on strings")
	}
	return qs, nil
}

func UpdateQuerySelectorBool(qs QuerySelector, query_type QueryType, cast_values []bool) (QuerySelector, error) {
	switch query_type {
	case In:
		qs["$in"] = cast_values
	case NotIn:
		qs["$nin"] = cast_values
	default:
		return nil, errors.New("Invalid operation on booleans")
	}
	return qs, nil
}

func UpdateQuerySelectorStringSlice(qs QuerySelector, query_type QueryType, cast_values []string) (QuerySelector, error) {
	switch query_type {
	case In:
		qs["$all"] = cast_values
	default:
		return nil, errors.New("Invalid operation on string slices")
	}
	return qs, nil
}

func ParseQueryType(key string) (QueryType, string) {
	query_type := In

	if strings.HasSuffix(key, "Lt") {
		key = key[0 : len(key)-2]
		query_type = LessThan
	} else if strings.HasSuffix(key, "Gt") {
		key = key[0 : len(key)-2]
		query_type = GreaterThan
	} else if strings.HasSuffix(key, "Not") {
		key = key[0 : len(key)-3]
		query_type = NotIn
	}
	return query_type, key
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

		query_type, key := ParseQueryType(key)
		key = strings.ToLower(key)

		to_type, ok := expected_types[key]
		if !ok {
			return nil, errors.New("Invalid key " + key)
		}

		values := strings.Split(values[0], ",")

		// Each query selector might have several entries
		// ie less than 10, greater than 2, not 4
		qs, ok := query[key].(QuerySelector)
		if qs == nil {
			qs = QuerySelector{}
		}

		var err error

		// We must specifically handle each data type
		switch to_type {
		case "[]string":
			qs, err = UpdateQuerySelectorStringSlice(qs, query_type, values)
			query[key] = qs
		case "string":
			qs, err = UpdateQuerySelectorString(qs, query_type, values)
			query[key] = qs
		case "int", "int64":
			cast_values := make([]int64, len(values))
			for i, value := range values {
				value_int, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, err
				}
				cast_values[i] = value_int
			}
			qs, err = UpdateQuerySelectorInt64(qs, query_type, cast_values)
			query[key] = qs
		case "bool":
			cast_values := make([]bool, len(values))
			for i, value := range values {
				value_int, err := strconv.ParseBool(value)
				if err != nil {
					return nil, err
				}
				cast_values[i] = value_int
			}
			qs, err = UpdateQuerySelectorBool(qs, query_type, cast_values)
			query[key] = qs
		}

		if err != nil {
			return nil, err
		}
	}

	return query, nil
}
