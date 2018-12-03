package tests

import (
	"encoding/json"
	"github.com/HackIllinois/api/common/configloader"
	"io/ioutil"
	"os"
	"testing"
)

func Setup(t *testing.T) {
	cfg_contents := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": map[string]interface{}{
			"key4": map[string]string{
				"key5": "value5",
				"key6": "value6",
			},
			"key7": "value7",
		},
	}

	cfg_json, err := json.Marshal(cfg_contents)

	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile("/tmp/testconfig.json", cfg_json, 0644)

	if err != nil {
		t.Fatal(err)
	}
}

func Teardown(t *testing.T) {
	err := os.Remove("/tmp/testconfig.json")

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Tests loading config from a file
*/
func TestLoadConfigFile(t *testing.T) {
	Setup(t)

	_, err := configloader.Load("file:///tmp/testconfig.json")

	if err != nil {
		t.Fatal(err)
	}

	Teardown(t)
}

/*
	Tests reading individual keys from a config
*/
func TestConfigGet(t *testing.T) {
	Setup(t)

	cfg, err := configloader.Load("file:///tmp/testconfig.json")

	if err != nil {
		t.Fatal(err)
	}

	value1, err := cfg.Get("key1")

	if err != nil {
		t.Fatal(err)
	}

	if value1 != "value1" {
		t.Errorf("Wrong value.\nExpected %v\ngot %v\n", "value1", value1)
	}

	value2, err := cfg.Get("key2")

	if err != nil {
		t.Fatal(err)
	}

	if value2 != "value2" {
		t.Errorf("Wrong value.\nExpected %v\ngot %v\n", "value2", value2)
	}

	Teardown(t)
}

/*
	Tests parsing complex keys into a struct
*/
func TestConfigParseInto(t *testing.T) {
	Setup(t)

	cfg, err := configloader.Load("file:///tmp/testconfig.json")

	if err != nil {
		t.Fatal(err)
	}

	value3 := struct {
		Value4 struct {
			Value5 string `json:"key5"`
			Value6 string `json:"key6"`
		} `json:"key4"`
		Value7 string `json:"key7"`
	}{}

	err = cfg.ParseInto("key3", &value3)

	if err != nil {
		t.Fatal(err)
	}

	if value3.Value4.Value5 != "value5" {
		t.Errorf("Wrong value.\nExpected %v\ngot %v\n", "value5", value3.Value4.Value5)
	}

	if value3.Value4.Value6 != "value6" {
		t.Errorf("Wrong value.\nExpected %v\ngot %v\n", "value6", value3.Value4.Value6)
	}

	if value3.Value7 != "value7" {
		t.Errorf("Wrong value.\nExpected %v\ngot %v\n", "value7", value3.Value7)
	}

	Teardown(t)
}
