package datastore

import (
	"errors"
)

type DataStoreDefinition struct {
	Name        string                `json:"name"`
	Type        string                `json:"type"`
	Validations string                `json:"validations"`
	Fields      []DataStoreDefinition `json:"fields"`
}

type DataStore struct {
	Definition DataStoreDefinition
	Data       map[string]interface{}
}

func NewDataStore(definition DataStoreDefinition) DataStore {
	return DataStore{
		Definition: definition,
	}
}

var conversionFuncs map[string](func(interface{}, DataStoreDefinition) (interface{}, error))

func init() {
	conversionFuncs = make(map[string](func(interface{}, DataStoreDefinition) (interface{}, error)))
	conversionFuncs["int"] = toInt
	conversionFuncs["float"] = toFloat
	conversionFuncs["string"] = toString
	conversionFuncs["boolean"] = toBoolean
	conversionFuncs["object"] = toObject
	conversionFuncs["[]int"] = toIntArray
	conversionFuncs["[]float"] = toFloatArray
	conversionFuncs["[]string"] = toStringArray
	conversionFuncs["[]boolean"] = toBooleanArray
	conversionFuncs["[]object"] = toObjectArray
}

func buildDataFromDefinition(raw_data interface{}, definition DataStoreDefinition) (interface{}, error) {
	conversionFunc, exists := conversionFuncs[definition.Type]

	if !exists {
		return nil, errors.New("Invalid type in definition")
	}

	return conversionFunc(raw_data, definition)
}

func defaultValueForType(tpe string) interface{} {
	switch tpe {
	case "string":
		return ""
	case "int":
		return 0
	case "float":
		return 0.0
	case "boolean":
		return false
	case "object":
		return nil
	case "[]string":
		return nil
	case "[]int":
		return nil
	case "[]float":
		return nil
	case "[]boolean":
		return nil
	case "[]object":
		return nil
	default:
		return nil
	}
}
