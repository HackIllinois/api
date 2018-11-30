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

var ErrInvalidDefinition = errors.New("DataStore definition is invalid")
var ErrInvalidData = errors.New("Invalid data unmarshalled")
var ErrTypeMismatch = errors.New("Type mismatch in data and definition")

var conversionFuncs map[string](func(interface{}, DataStoreDefinition) (interface{}, error))
var defaultValues map[string]interface{}

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

	defaultValues = make(map[string]interface{})
	defaultValues["int"] = 0
	defaultValues["float"] = 0.0
	defaultValues["string"] = ""
	defaultValues["boolean"] = false
	defaultValues["object"] = nil
	defaultValues["[]int"] = nil
	defaultValues["[]float"] = nil
	defaultValues["[]string"] = nil
	defaultValues["[]boolean"] = nil
	defaultValues["[]object"] = nil
}
