package tests

import (
	"github.com/HackIllinois/api/services/auth/service"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

const URL_SCHEME = "http"
const URL_HOST = "hackillinois.com"
const URL_PATH = "register/now"
const CONSTRUCTED_BASIC_URL = "http://hackillinois.com/register/now"

const DIFFERENT_SCHEME = "ftp"
const HOST_WITH_PORT = "hackillinois.com:9800"
const CONSTRUCTED_COMPLEX_URL = "ftp://hackillinois.com:9800/register/now"

var QUERY_PARAMS = map[string]string{
	"major":           "CS",
	"graduating_year": "2021",
}

var EXPECTED_QUERY_VALUES, _ = url.ParseQuery("?graduating_year=2021&major=CS")

var QUERY_PARAMS_HASHTAG = map[string]string{
	"major":      "CS",
	"twitter":    "#hackillinois",
	"dead_param": "true",
}

var SINGLE_PARAM = map[string]string{
	"major": "CS",
}

const SINGLE_PARAM_QUERY_STRING = "?major=CS"

/*
   Test that a simple URL with no query params can be generated
*/
func TestConstructSimpleURL(t *testing.T) {
	result, err := service.ConstructSafeURL(URL_SCHEME, URL_HOST, URL_PATH, nil)

	if err != nil {
		t.Fatal(err)
	}

	if result != CONSTRUCTED_BASIC_URL {
		t.Errorf("URL not correctly constructed. Expected \"%v\", got \"%v\"", CONSTRUCTED_BASIC_URL, result)
	}
}

/*
   Test that a more complicated URL with a non http(s) scheme and a port
   can be generated.
*/
func TestConstructComplexURL(t *testing.T) {
	result, err := service.ConstructSafeURL(DIFFERENT_SCHEME, HOST_WITH_PORT, URL_PATH, nil)

	if err != nil {
		t.Fatal(err)
	}

	if result != CONSTRUCTED_COMPLEX_URL {
		t.Errorf("URL not correctly constructed. Expected \"%v\", got \"%v\"", CONSTRUCTED_COMPLEX_URL, result)
	}
}

/*
   Test that building a simple query into a URL works.
*/
func TestQueryStringBuilder(t *testing.T) {
	generatedURL := url.URL{
		Scheme: URL_SCHEME,
		Host:   URL_HOST,
		Path:   URL_PATH,
	}

	service.ConstructURLQuery(&generatedURL, QUERY_PARAMS)

	// If the query string separator is not in the URL, the test fails.
	if !strings.Contains(generatedURL.String(), "?") {
		t.Error("The `?` character was not present in a URL that has a query string")
	}

	queryString := "?" + strings.Split(generatedURL.String(), "?")[1]

	// We can't guarantee the order of the query string, so we reparse it to
	// later check for equality.
	reparsedQueryValues, err := url.ParseQuery(queryString)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(reparsedQueryValues, EXPECTED_QUERY_VALUES) {
		t.Error("Query string not correctly constructed.")
	}
}

/*
   Test that a URL with a scheme, host, path, and a query param is correctly constructed.
*/
func TestConstructFullURL(t *testing.T) {
	result, err := service.ConstructSafeURL(URL_SCHEME, URL_HOST, URL_PATH, SINGLE_PARAM)

	if err != nil {
		t.Fatal(err)
	}

	expectedURL := CONSTRUCTED_BASIC_URL + SINGLE_PARAM_QUERY_STRING
	if result != expectedURL {
		t.Errorf("URL not correctly constructed. Expected \"%v\", got \"%v\"", expectedURL, result)
	}
}

/*
   Test that a URL with a fragment (`#` character) gets caught and throws an error.
*/
func TestFragmentCaughtURL(t *testing.T) {
	_, err := service.ConstructSafeURL(DIFFERENT_SCHEME, HOST_WITH_PORT, URL_PATH, QUERY_PARAMS_HASHTAG)

	if err == nil {
		t.Error("The `#` in this URL was not caught!")
	}
}
