package tests

import (
	"github.com/HackIllinois/api/services/auth/service"
	"testing"
)

const URL_SCHEME = "http"
const URL_HOST = "hackillinois.com"
const URL_PATH = "register/now"
const CONSTRUCTED_BASIC_URL = "http://hackillinois.com/register/now"

const DIFFERENT_SCHEME = "ftp"
const HOST_WITH_PORT = "hackillinois.com:9800"
const CONSTRUCTED_COMPLEX_URL = "ftp://hackillinois.com:9800/register/now"

/*
   Test that a simple URL with no query params can be generated
*/
func TestConstructSimpleURL(t *testing.T) {
	result, err := service.ConstructSafeURL(URL_SCHEME, URL_HOST, URL_PATH, nil)

	if result != CONSTRUCTED_BASIC_URL {
		t.Errorf("URL not correctly instructed. Expected \"%v\", got \"%v\"", CONSTRUCTED_BASIC_URL, result)
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
		t.Errorf("URL not correctly instructed. Expected \"%v\", got \"%v\"", CONSTRUCTED_COMPLEX_URL, result)
	}

	if err != nil {
		t.Fatal(err)
	}
}
