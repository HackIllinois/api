package tests

import (
	"encoding/json"
	"github.com/HackIllinois/api/common/datastore"
	"testing"
)

var smallJsonData string = `
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

var smallJsonDefinition string = `
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

var largeJsonData = `
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

var largeJsonDefiniton = `
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
	var definiton datastore.DataStoreDefinition
	err := json.Unmarshal([]byte(smallJsonDefinition), &definiton)

	if err != nil {
		t.Fatal(err)
	}

	store := datastore.NewDataStore(definiton)

	err = json.Unmarshal([]byte(smallJsonData), &store)

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

func TestDatastoreComplex(t *testing.T) {
	var definiton datastore.DataStoreDefinition
	err := json.Unmarshal([]byte(largeJsonDefiniton), &definiton)

	if err != nil {
		t.Fatal(err)
	}

	store := datastore.NewDataStore(definiton)

	err = json.Unmarshal([]byte(largeJsonData), &store)

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

	marshalledData, err := json.Marshal(&store)

	if err != nil {
		t.Fatal(err)
	}

	var rawData map[string]interface{}
	err = json.Unmarshal(marshalledData, &rawData)

	if err != nil {
		t.Fatal(err)
	}

	if rawData["firstName"] != "test" {
		t.Errorf("Wrong info.\nExpected %v\ngot %v\n", "test", rawData["firstName"])
	}

	if rawData["age"] != float64(19) {
		t.Errorf("Wrong info.\nExpected %v\ngot %v\n", float64(19), rawData["age"])
	}

	if rawData["isNovice"] != false {
		t.Errorf("Wrong info.\nExpected %v\ngot %v\n", false, rawData["isNovice"])
	}

	store.Data["createdAt"] = 1536178263
	store.Data["updatedAt"] = 1543011468

	err = store.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
