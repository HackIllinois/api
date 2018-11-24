package datastore

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type DataStoreDefinition struct {
	Name        string                `json:"name"`
	Type        string                `json:"type"`
	IsArray     bool                  `json:"IsArray"`
	Validations string                `json:"validations"`
	Fields      []DataStoreDefinition `json:"fields"`
}

type DataStore struct {
	Definition DataStoreDefinition
	rawData    map[string]interface{}
	Data       map[string]interface{}
}

func NewDataStore(definition DataStoreDefinition) DataStore {
	return DataStore{
		Definition: definition,
	}
}

func (datastore *DataStore) Validate() error {
	validate := validator.New()

	return validateField(datastore.Data, datastore.Definition, validate)
}

func validateField(data interface{}, definition DataStoreDefinition, validate *validator.Validate) error {

	fmt.Printf("%T: %v\n\n", data, data)

	err := validate.Var(data, definition.Validations)

	if err != nil {
		return err
	}

	switch definition.Type {
	case "object":
		for _, field := range definition.Fields {
			mappedData, ok := data.(map[string]interface{})

			if !ok {
				return errors.New("Definition contains field for non-mappable data")
			}

			err = validateField(mappedData[field.Name], field, validate)

			if err != nil {
				return err
			}
		}
	case "[]object":
		dataArray, ok := data.([]map[string]interface{})

		if !ok {
			return errors.New("Data format does not match definition")
		}

		for _, mappedData := range dataArray {
			for _, field := range definition.Fields {
				err = validateField(mappedData[field.Name], field, validate)

				if err != nil {
					return err
				}
			}
		}
	default:
	}

	return nil
}

func (datastore *DataStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(&datastore.Data)
}

func (datastore *DataStore) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &datastore.rawData)

	if err != nil {
		return err
	}

	data, err := buildDataFromDefinition(datastore.rawData, datastore.Definition)

	if err != nil {
		return err
	}

	var ok bool
	datastore.Data, ok = data.(map[string]interface{})

	if !ok {
		return errors.New("Invalid data unmarshalled")
	}

	return nil
}

func buildDataFromDefinition(rawData interface{}, definition DataStoreDefinition) (interface{}, error) {
	switch definition.Type {
	case "int":
		data, ok := rawData.(float64)

		if !ok {
			return nil, errors.New("Type mismatch in data and definition")
		}

		return int64(data), nil
	case "float":
		data, ok := rawData.(float64)

		if !ok {
			return nil, errors.New("Type mismatch in data and definition")
		}

		return data, nil
	case "string":
		data, ok := rawData.(string)

		if !ok {
			return nil, errors.New("Type mismatch in data and definition")
		}

		return data, nil
	case "boolean":
		data, ok := rawData.(bool)

		if !ok {
			return nil, errors.New("Type mismatch in data and definition")
		}

		return data, nil
	case "object":
		unfilteredData, ok := rawData.(map[string]interface{})

		if !ok {
			return nil, errors.New("Type mismatch in data and definition")
		}

		data := make(map[string]interface{})

		for _, field := range definition.Fields {
			unfilteredField, exists := unfilteredData[field.Name]

			if !exists {
				return nil, errors.New("Missing field in data")
			}

			var err error
			data[field.Name], err = buildDataFromDefinition(unfilteredField, field)

			if err != nil {
				return nil, err
			}
		}

		return data, nil
	case "[]int":
		data, ok := rawData.([]interface{})

		if !ok {
			return nil, errors.New("Type mismatch in data and definition")
		}

		intData := make([]int64, len(data))

		for i := 0; i < len(data); i++ {
			element, ok := data[i].(float64)

			if !ok {
				return nil, errors.New("Type mismatch in data and definition")
			}

			intData[i] = int64(element)
		}

		return intData, nil
	case "[]float":
		data, ok := rawData.([]interface{})

		if !ok {
			return nil, errors.New("Type mismatch in data and definition")
		}

		floatData := make([]float64, len(data))

		for i := 0; i < len(data); i++ {
			element, ok := data[i].(float64)

			if !ok {
				return nil, errors.New("Type mismatch in data and definition")
			}

			floatData[i] = element
		}

		return floatData, nil
	case "[]string":
		data, ok := rawData.([]interface{})

		if !ok {
			return nil, errors.New("Type mismatch in data and definition")
		}

		stringData := make([]string, len(data))

		for i := 0; i < len(data); i++ {
			element, ok := data[i].(string)

			if !ok {
				return nil, errors.New("Type mismatch in data and definition")
			}

			stringData[i] = element
		}

		return stringData, nil
	case "[]boolean":
		data, ok := rawData.([]interface{})

		if !ok {
			return nil, errors.New("Type mismatch in data and definition")
		}

		boolData := make([]bool, len(data))

		for i := 0; i < len(data); i++ {
			element, ok := data[i].(bool)

			if !ok {
				return nil, errors.New("Type mismatch in data and definition")
			}

			boolData[i] = element
		}

		return boolData, nil
	case "[]object":
		unfilteredData, ok := rawData.([]interface{})

		if !ok {
			return nil, errors.New("Type mismatch in data and definition")
		}

		data := make([]map[string]interface{}, len(unfilteredData))

		for i := 0; i < len(unfilteredData); i++ {
			element := make(map[string]interface{})

			for _, field := range definition.Fields {
				unfilteredDataElement, ok := unfilteredData[i].(map[string]interface{})

				if !ok {
					return nil, errors.New("Type mismatch in data and definition")
				}

				unfilteredField, exists := unfilteredDataElement[field.Name]

				if !exists {
					return nil, errors.New("Missing field in data")
				}

				var err error
				element[field.Name], err = buildDataFromDefinition(unfilteredField, field)

				if err != nil {
					return nil, err
				}
			}

			data[i] = element
		}

		return data, nil
	default:
		return nil, errors.New("Invalid type in definition")
	}
}
