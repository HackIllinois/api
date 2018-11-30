package datastore

import (
	"errors"
	"gopkg.in/go-playground/validator.v9"
)

func (datastore *DataStore) Validate() error {
	validate := validator.New()

	return validateField(datastore.Data, datastore.Definition, validate)
}

func validateField(data interface{}, definition DataStoreDefinition, validate *validator.Validate) error {
	err := validate.Var(data, definition.Validations)

	if err != nil {
		return err
	}

	switch definition.Type {
	case "object":
		for _, field := range definition.Fields {
			mapped_data, ok := data.(map[string]interface{})

			if !ok {
				return errors.New("Definition contains field for non-mappable data")
			}

			err = validateField(mapped_data[field.Name], field, validate)

			if err != nil {
				return err
			}
		}
	case "[]object":
		data_array, ok := data.([]map[string]interface{})

		if !ok {
			return errors.New("Data format does not match definition")
		}

		for _, mapped_data := range data_array {
			for _, field := range definition.Fields {
				err = validateField(mapped_data[field.Name], field, validate)

				if err != nil {
					return err
				}
			}
		}
	default:
	}

	return nil
}
