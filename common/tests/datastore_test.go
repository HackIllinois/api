package tests

import (
	"encoding/json"
	"fmt"
	"github.com/HackIllinois/api/common/datastore"
	"testing"
)

var jsonData string = `
{
	"intKey": 100,
	"stringKey": "value",
	"objectKey": {
		"stringKeys": [
			"value2",
			"value3"
		],
		"objectKeys": [
			{
				"thing1": 2,
				"thing2": "a"
			}
		]
	}
}
`

var jsonDefinition string = `
{
	"name": "topLevel",
	"type": "object",
	"validations": "required",
	"fields": [
		{
			"name": "intKey",
			"type": "int",
			"validations": "required",
			"fields": []
		},
		{
			"name": "stringKey",
			"type": "string",
			"validations": "required",
			"fields": []
		},
		{
			"name": "objectKey",
			"type": "object",
			"validations": "required",
			"fields": [
				{
					"name": "stringKeys",
					"type": "[]string",
					"validations": "required,dive,required,oneof=value2 value3",
					"fields": []
				},
				{
					"name": "objectKeys",
					"type": "[]object",
					"validations": "required",
					"fields": [
						{
							"name": "thing1",
							"type": "int",
							"validations": "required",
							"fields": []
						},
						{
							"name": "thing2",
							"type": "string",
							"validations": "required",
							"fields": []
						}
					]
				}
			]
		}
	]
}
`

func TestDatastore(t *testing.T) {
	var definiton datastore.DataStoreDefinition
	err := json.Unmarshal([]byte(jsonDefinition), &definiton)

	if err != nil {
		t.Fatal(err)
	}

	store := datastore.NewDataStore(definiton)

	err = json.Unmarshal([]byte(jsonData), &store)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%v\n", store)

	marshalledData, err := json.Marshal(&store)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%v\n", string(marshalledData))

	err = store.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
