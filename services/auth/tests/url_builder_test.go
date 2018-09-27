package tests

import (
	"github.com/HackIllinois/api/services/auth/service"
	"testing"
    "strings"
    "net/url"
)

const URL_SCHEME = "http"
const URL_HOST = "hackillinois.com"
const URL_PATH = "register/now"
const CONSTRUCTED_BASIC_URL = "http://hackillinois.com/register/now"

const DIFFERENT_SCHEME = "ftp"
const HOST_WITH_PORT = "hackillinois.com:9800"
const CONSTRUCTED_COMPLEX_URL = "ftp://hackillinois.com:9800/register/now"

var QUERY_PARAMS = map[string]string{
    "major": "CS",
    "graduating_year": "2021",
}

const EXPECTED_QUERY_STRING = "?graduating_year=2021&major=CS"

var QUERY_PARAMS_HASHTAG = map[string]string{
    "major": "CS",
    "twitter": "#hackillinois",
    "dead_param": "true",
}
const EXPECTED_QUERY_STRING_HASHTAG = "?major=CS&twitter=#hackillinois&dead_param=true"

/*
   Test that a simple URL with no query params can be generated
*/
func TestConstructSimpleURL(t *testing.T) {
	result, err := service.ConstructSafeURL(URL_SCHEME, URL_HOST, URL_PATH, nil)

	if result != CONSTRUCTED_BASIC_URL {
		t.Errorf("URL not correctly constructed. Expected \"%v\", got \"%v\"", CONSTRUCTED_BASIC_URL, result)
	}

	if err != nil {
		t.Fatal(err)
	}
}

/*
   Test that a more complicated URL with a non http(s) scheme and a port
   can be generated.
*/
func TestConstructComplexURL(t *testing.T) {
	result, err := service.ConstructSafeURL(DIFFERENT_SCHEME, HOST_WITH_PORT, URL_PATH, nil)

	if result != CONSTRUCTED_COMPLEX_URL {
		t.Errorf("URL not correctly constructed. Expected \"%v\", got \"%v\"", CONSTRUCTED_COMPLEX_URL, result)
	}

	if err != nil {
		t.Fatal(err)
	}
}


/*
    Test that building a simple query into a URL works.
*/
func TestQueryStringBuilder(t *testing.T) {
    url := url.URL{
        Scheme: URL_SCHEME,
        Host: URL_HOST,
        Path: URL_PATH,
    }

    service.ConstructURLQuery(&url, QUERY_PARAMS)

    queryString := "?" + strings.Split(url.String(), "?")[1]

    if queryString != EXPECTED_QUERY_STRING {
		t.Errorf("Query string not correctly constructed. Expected \"%v\", got \"%v\"", EXPECTED_QUERY_STRING, queryString)
    }
}

/*
    Test that a URL with a scheme, host, path, and query params is correctly constructed.
*/
func TestConstructFullURL(t *testing.T) {
	result, err := service.ConstructSafeURL(URL_SCHEME, URL_HOST, URL_PATH, QUERY_PARAMS)

    expectedURL := CONSTRUCTED_BASIC_URL + EXPECTED_QUERY_STRING
    if result != expectedURL {
		t.Errorf("URL not correctly constructed. Expected \"%v\", got \"%v\"", expectedURL, result)
    }

	if err != nil {
		t.Fatal(err)
	}
}

/*
    Test that a URL with a fragment (`#` character) gets caught and throws an error.
*/
func TestFragmentCaughtURL(t *testing.T) {
	_, err := service.ConstructSafeURL(DIFFERENT_SCHEME, HOST_WITH_PORT, URL_PATH, QUERY_PARAMS_HASHTAG)

    if err != service.HASHTAG_INVALID_ERR {
        t.Error("The `#` in this URL was not caught!")
    }
}
