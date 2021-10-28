package datastore

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

func (datastore *DataStore) Validate() error {
	validate := validator.New()

	return validateField(datastore.Data, datastore.Definition, validate, false)
}

func (datastore *DataStore) ValidateNonEmpty() error {
	validate := validator.New()

	return validateField(datastore.Data, datastore.Definition, validate, true)
}

func validateField(
	data interface{},
	definition DataStoreDefinition,
	validate *validator.Validate,
	ignore_empty bool,
) error {
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

		return validateFieldArray(mapped_data, definition, validate, ignore_empty)
	case "[]object":
		data_array, ok := data.([]map[string]interface{})

		if !ok {
			return NewErrTypeMismatch(data, "[]map[string]interface{}")
		}

		for _, mapped_data := range data_array {
			err = validateFieldArray(mapped_data, definition, validate, ignore_empty)

			if err != nil {
				return err
			}
		}

		return nil
	default:
		return nil
	}
}

func validateFieldArray(
	data map[string]interface{},
	definition DataStoreDefinition,
	validate *validator.Validate,
	ignore_empty bool,
) error {
	for _, field := range definition.Fields {
		if ignore_empty {
			if _, ok := data[field.Name]; !ok {
				continue
			}
		}

		err := validateField(data[field.Name], field, validate, ignore_empty)

		if err != nil {
			return err
		}
	}

	return nil
}
