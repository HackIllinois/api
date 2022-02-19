package datastore

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func (datastore *DataStore) Validate() error {
	validate := validator.New()

	return validateField(datastore.Data, datastore.Definition, validate)
}

func validateField(data interface{}, definition DataStoreDefinition, validate *validator.Validate) error {
	err := validate.Var(data, definition.Validations)

	if err != nil {
		return fmt.Errorf("Key '%v' with value '%v' failed validation '%v'", definition.Name, data, definition.Validations)
	}

	switch definition.Type {
	case "object":
		mapped_data, ok := data.(map[string]interface{})

		if !ok {
			return NewErrTypeMismatch(data, "map[string]interface{}")
		}

		return validateFieldArray(mapped_data, definition, validate)
	case "[]object":
		data_array, ok := data.([]map[string]interface{})

		if !ok {
			return NewErrTypeMismatch(data, "[]map[string]interface{}")
		}

		for _, mapped_data := range data_array {
			err = validateFieldArray(mapped_data, definition, validate)

			if err != nil {
				return err
			}
		}

		return nil
	default:
		return nil
	}
}

func validateFieldArray(data map[string]interface{}, definition DataStoreDefinition, validate *validator.Validate) error {
	for _, field := range definition.Fields {
		err := validateField(data[field.Name], field, validate)

		if err != nil {
			return err
		}
	}

	return nil
}
