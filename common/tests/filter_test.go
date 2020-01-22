package tests

import (
	"reflect"
	"testing"

	"github.com/HackIllinois/api/common/database"
)

type TestStruct struct {
	Value1 string `json:"value1"`
	Value2 string `json:"value2"`
}

func TestFilterBasic(t *testing.T) {

	params := map[string][]string{
		"value1": {"test"},
		"value2": {"foo,bar"},
	}

	expected_query := map[string]interface{}{
		"value1": database.QuerySelector{"$in": []string{"test"}},
		"value2": database.QuerySelector{"$in": []string{"foo", "bar"}},
	}

	query, err := database.CreateFilterQuery(params, TestStruct{})
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(query, expected_query) {
		t.Errorf("Incorrect query.\nExpected %v\ngot %v\n", expected_query, query)
	}
}

func TestFilterMissing(t *testing.T) {

	// value2 missing
	params := map[string][]string{
		"value1": {"test1,test2,test3"},
	}

	expected_query := map[string]interface{}{
		"value1": database.QuerySelector{"$in": []string{"test1", "test2", "test3"}},
	}

	query, err := database.CreateFilterQuery(params, TestStruct{})
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(query, expected_query) {
		t.Errorf("Incorrect query.\nExpected %v\ngot %v\n", expected_query, query)
	}
}

type TestStruct2 struct {
	Hack     string `json:"hackValue"`
	Illinois int64  `json:"illinoisValue"`
	Test     bool   `json:"testValue"`
}

func TestFilterCasting(t *testing.T) {

	params := map[string][]string{
		"hackValue":     {"foo,bar"},
		"illinoisValue": {"55,63"},
		"testValue":     {"true,false,true"},
	}

	expected_query := map[string]interface{}{
		"hackvalue":     database.QuerySelector{"$in": []string{"foo", "bar"}},
		"illinoisvalue": database.QuerySelector{"$in": []int64{55, 63}},
		"testvalue":     database.QuerySelector{"$in": []bool{true, false, true}},
	}

	query, err := database.CreateFilterQuery(params, TestStruct2{})
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(query, expected_query) {
		t.Errorf("Incorrect query.\nExpected %v\ngot %v\n", expected_query, query)
	}
}
