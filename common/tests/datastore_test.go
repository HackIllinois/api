package tests

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/HackIllinois/api/common/datastore"
	"go.mongodb.org/mongo-driver/bson"
)

var small_json_data string = `
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

var small_invalid_json_data string = `
{
	"intKey": "test",
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

var small_invalid_json_data_2 string = `
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
				"thing2": 5
			}
		]
	}
}
`

var small_json_definition string = `
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

var large_json_data = `
{
    "id": "google000000000000000000001",
    "firstName": "test",
    "lastName": "user",
    "email": "testemail@gmail.com",
    "shirtSize": "M",
    "diet": "NONE",
    "age": 19,
    "graduationYear": 2020,
    "transportation": "NONE",
    "school": "University of Illinois at Urbana-Champaign",
    "major": "Computer Science",
    "gender": "MALE",
    "professionalInterest": "INTERNSHIP",
    "github": "test gihubusername",
    "linkedin": "some-account",
    "interests": "Software",
    "isNovice": false,
    "isPrivate": false,
    "phoneNumber": "555-555-5555",
    "longforms": [
        {
            "response": "This is a longform."
        }
    ],
    "extraInfos": [
        {
            "response": "This is an extra info."
        }
    ],
    "osContributors": [
        {
            "contactInfo": "person@gmail.com",
            "name": "Person"
        }
    ],
    "collaborators": [
        {
            "github": "persongithub"
        }
    ]
}
`

var large_json_definition = `
{
    "name": "user_registration",
    "type": "object",
    "validations": "required",
    "fields": [
        {
            "name": "id",
            "type": "string",
            "validations": "required",
            "fields": []
        },
        {
            "name": "firstName",
            "type": "string",
            "validations": "required",
            "fields": []
        },
        {
            "name": "lastName",
            "type": "string",
            "validations": "required",
            "fields": []
        },
        {
            "name": "email",
            "type": "string",
            "validations": "required,email",
            "fields": []
        },
        {
            "name": "shirtSize",
            "type": "string",
            "validations": "required,oneof=S M L XL",
            "fields": []
        },
        {
            "name": "diet",
            "type": "string",
            "validations": "required,oneof=NONE VEGAN VEGETARIAN",
            "fields": []
        },
        {
            "name": "age",
            "type": "int",
            "validations": "required",
            "fields": []
        },
        {
            "name": "graduationYear",
            "type": "int",
            "validations": "required",
            "fields": []
        },
        {
            "name": "transportation",
            "type": "string",
            "validations": "required,oneof=NONE BUS",
            "fields": []
        },
        {
            "name": "school",
            "type": "string",
            "validations": "required",
            "fields": []
        },
        {
            "name": "major",
            "type": "string",
            "validations": "required",
            "fields": []
        },
        {
            "name": "gender",
            "type": "string",
            "validations": "required,oneof=MALE FEMALE NONBINARY OTHER",
            "fields": []
        },
        {
            "name": "professionalInterest",
            "type": "string",
            "validations": "required,oneof=INTERNSHIP FULLTIME BOTH NONE",
            "fields": []
        },
        {
            "name": "github",
            "type": "string",
            "validations": "required",
            "fields": []
        },
        {
            "name": "linkedin",
            "type": "string",
            "validations": "required",
            "fields": []
        },
        {
            "name": "interests",
            "type": "string",
            "validations": "required",
            "fields": []
        },
        {
            "name": "isNovice",
            "type": "boolean",
            "validations": "required|isdefault",
            "fields": []
        },
        {
            "name": "isPrivate",
            "type": "boolean",
            "validations": "required|isdefault",
            "fields": []
        },
        {
            "name": "phoneNumber",
            "type": "string",
            "validations": "required",
            "fields": []
        },
        {
            "name": "longforms",
            "type": "[]object",
            "validations": "required",
            "fields": [
                {
                    "name": "response",
                    "type": "string",
                    "validations": "required",
                    "fields": []
                }
            ]
        },
        {
            "name": "extraInfos",
            "type": "[]object",
            "validations": "required",
            "fields": [
                {
                    "name": "response",
                    "type": "string",
                    "validations": "required",
                    "fields": []
                }
            ]
        },
        {
            "name": "osContributors",
            "type": "[]object",
            "validations": "required",
            "fields": [
                {
                    "name": "contactInfo",
                    "type": "string",
                    "validations": "required",
                    "fields": []
                },
                {
                    "name": "name",
                    "type": "string",
                    "validations": "required",
                    "fields": []
                }
            ]
        },
        {
            "name": "collaborators",
            "type": "[]object",
            "validations": "required",
            "fields": [
                {
                    "name": "github",
                    "type": "string",
                    "validations": "required",
                    "fields": []
                }
            ]
        },
        {
            "name": "createdAt",
            "type": "int",
            "validations": "required",
            "fields": []
        },
        {
            "name": "updatedAt",
            "type": "int",
            "validations": "required",
            "fields": []
        }
    ]
}
`

func TestDatastoreBasic(t *testing.T) {
	var definition datastore.DataStoreDefinition
	err := json.Unmarshal([]byte(small_json_definition), &definition)

	if err != nil {
		t.Fatal(err)
	}

	store := datastore.NewDataStore(definition)

	err = json.Unmarshal([]byte(small_json_data), &store)

	if err != nil {
		t.Fatal(err)
	}

	_, err = json.Marshal(&store)

	if err != nil {
		t.Fatal(err)
	}

	err = store.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

func TestDatastoreAssertions(t *testing.T) {
	var definition datastore.DataStoreDefinition
	err := json.Unmarshal([]byte(small_json_definition), &definition)

	if err != nil {
		t.Fatal(err)
	}

	store := datastore.NewDataStore(definition)
	err = json.Unmarshal([]byte(small_invalid_json_data), &store)

	expected_err_inner := errors.New("Type mismatch in data and definition. Expected float64, got string")
	expected_err := datastore.ErrorInField{
		FieldName: "intKey",
		Err:       expected_err_inner,
	}

	if fmt.Sprint(err) != fmt.Sprint(expected_err) {
		t.Errorf("Wrong field name.\nExpected %s\ngot %s\n", expected_err, err)
	}

	store = datastore.NewDataStore(definition)
	err = json.Unmarshal([]byte(small_invalid_json_data_2), &store)

	expected_err_inner = errors.New("Type mismatch in data and definition. Expected string, got float64")
	expected_err = datastore.ErrorInField{
		FieldName: "objectKey.objectKeys.thing2",
		Err:       expected_err_inner,
	}

	if fmt.Sprint(err) != fmt.Sprint(expected_err) {
		t.Errorf("Wrong field name.\nExpected %s\ngot %s\n", expected_err, err)
	}
}

func TestDatastoreComplex(t *testing.T) {
	var definition datastore.DataStoreDefinition
	err := json.Unmarshal([]byte(large_json_definition), &definition)

	if err != nil {
		t.Fatal(err)
	}

	store := datastore.NewDataStore(definition)

	err = json.Unmarshal([]byte(large_json_data), &store)

	if err != nil {
		t.Fatal(err)
	}

	if store.Data["firstName"] != "test" {
		t.Errorf("Wrong info.\nExpected %v\ngot %v\n", "test", store.Data["firstName"])
	}

	if store.Data["age"] != int64(19) {
		t.Errorf("Wrong info.\nExpected %v\ngot %v\n", int64(19), store.Data["age"])
	}

	if store.Data["isNovice"] != false {
		t.Errorf("Wrong info.\nExpected %v\ngot %v\n", false, store.Data["isNovice"])
	}

	marshalled_data, err := json.Marshal(&store)

	if err != nil {
		t.Fatal(err)
	}

	var raw_data map[string]interface{}
	err = json.Unmarshal(marshalled_data, &raw_data)

	if err != nil {
		t.Fatal(err)
	}

	if raw_data["firstName"] != "test" {
		t.Errorf("Wrong info.\nExpected %v\ngot %v\n", "test", raw_data["firstName"])
	}

	if raw_data["age"] != float64(19) {
		t.Errorf("Wrong info.\nExpected %v\ngot %v\n", float64(19), raw_data["age"])
	}

	if raw_data["isNovice"] != false {
		t.Errorf("Wrong info.\nExpected %v\ngot %v\n", false, raw_data["isNovice"])
	}

	store.Data["createdAt"] = 1536178263
	store.Data["updatedAt"] = 1543011468

	err = store.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

func TestBSONConversions(t *testing.T) {
	var definition datastore.DataStoreDefinition
	err := json.Unmarshal([]byte(large_json_definition), &definition)

	if err != nil {
		t.Fatal(err)
	}

	store := datastore.NewDataStore(definition)

	err = json.Unmarshal([]byte(large_json_data), &store)

	if err != nil {
		t.Fatal(err)
	}

	bson_marshalled, err := bson.Marshal(&store)

	if err != nil {
		t.Fatal(err)
	}

	var unmarshalled_store datastore.DataStore
	err = bson.Unmarshal(bson_marshalled, &unmarshalled_store)

	if err != nil {
		t.Fatal(err)
	}

	if unmarshalled_store.Data["firstName"] != store.Data["firstName"] {
		t.Errorf("Wrong info.\nExpected %v\ngot %v\n", store.Data["firstName"], unmarshalled_store.Data["firstName"])
	}

	if unmarshalled_store.Data["age"] != store.Data["age"] {
		t.Errorf("Wrong info.\nExpected %v\ngot %v\n", store.Data["age"], unmarshalled_store.Data["age"])
	}

	if unmarshalled_store.Data["isNovice"] != store.Data["isNovice"] {
		t.Errorf("Wrong info.\nExpected %v\ngot %v\n", store.Data["isNovice"], unmarshalled_store.Data["isNovice"])
	}
}
