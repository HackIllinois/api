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

func buildDataFromDefinition(raw_data interface{}, definition DataStoreDefinition) (interface{}, error) {
	switch definition.Type {
	case "int":
		return toInt(raw_data)
	case "float":
		return toFloat(raw_data)
	case "string":
		return toString(raw_data)
	case "boolean":
		return toBoolean(raw_data)
	case "object":
		return toObject(raw_data, definition)
	case "[]int":
		return toIntArray(raw_data)
	case "[]float":
		return toFloatArray(raw_data)
	case "[]string":
		return toStringArray(raw_data)
	case "[]boolean":
		return toBooleanArray(raw_data)
	case "[]object":
		return toObjectArray(raw_data, definition)
	default:
		return nil, errors.New("Invalid type in definition")
	}
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
