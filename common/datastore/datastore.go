package datastore

import (
	"errors"
	"fmt"
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

func NewErrTypeMismatch(raw_data interface{}, expected string) error {
	return fmt.Errorf("Type mismatch in data and definition. Expected %s, got %T", expected, raw_data)
}

type ErrorInField struct {
	FieldName string
	Err       error
}

func (e ErrorInField) Error() string {
	return fmt.Sprintf("Error in field %s: %s", e.FieldName, e.Err)
}

/*
	Wraps the given error, to indicate which field it comes from
	If the provided error is of type ErrorInField, this function prepends the field name instead
	For example, an error 'Error in field: name: xyz' becomes 'Error in field: user.name: xyz'
*/
func NewErrInField(field_name string, err error) error {
	old_err_in_field, ok := err.(ErrorInField)
	if ok {
		field_name = fmt.Sprintf("%s.%s", field_name, old_err_in_field.FieldName)
		err = old_err_in_field.Err
	}
	return ErrorInField{FieldName: field_name, Err: err}
}

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
