package datastore

import (
	"errors"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
)

func (datastore *DataStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(&datastore.Data)
}

func (datastore *DataStore) UnmarshalJSON(b []byte) error {
	var raw_data map[string]interface{}
	err := json.Unmarshal(b, &raw_data)

	if err != nil {
		return err
	}

	data, err := buildDataFromDefinition(raw_data, datastore.Definition)

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

func (datastore *DataStore) GetBSON() (interface{}, error) {
	return datastore.Data, nil
}

func (datastore *DataStore) SetBSON(raw bson.Raw) error {
	return raw.Unmarshal(&datastore.Data)
}
