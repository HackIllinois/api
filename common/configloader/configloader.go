package configloader

import (
	"os"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var ErrNotSet = errors.New("The value for the given key was not set")
var ErrDecodeFailed = errors.New("The value for the given key could not be decoded")
var ErrLoadFailed = errors.New("Unable to load config")

/*
	Used to load a key value configuration
*/
type ConfigLoader struct {
	configPath string
	parsedConfig map[string]*json.RawMessage
}

/*
	Loads the configuration at the given path into a ConfigLoader struct
	Supported uri schemes are: s3, file, https
*/
func Load(configPath string) (*ConfigLoader, error) {
	uri, err := url.Parse(configPath)

	if err != nil {
		return nil, ErrLoadFailed
	}

	var configContents []byte

	switch uri.Scheme {
	case "s3":
		configContents, err = loadFromS3(configPath)
	case "file":
		configContents, err = loadFromFile(configPath)
	case "https":
		configContents, err = loadFromHttps(configPath)
	default:
		return nil, ErrLoadFailed
	}

	if err != nil {
		return nil, ErrLoadFailed
	}

	loader := ConfigLoader {
		configPath: configPath,
	}

	err = json.Unmarshal(configContents, &loader.parsedConfig)

	if err != nil {
		return nil, ErrLoadFailed
	}

	return &loader, nil
}

/*
	Returns the value associated with a given key as a string
	Environment variables will override configuration
*/
func (loader *ConfigLoader) Get(key string) (string, error) {
	value, exists := os.LookupEnv(key)

	if exists {
		return value, nil
	}

	rawValue, exists := loader.parsedConfig[key]

	if !exists {
		return "", ErrNotSet
	}

	bytes, err := json.Marshal(rawValue)

	if err != nil {
		return "", ErrDecodeFailed
	}

	return string(bytes), nil
}

/*
	Parses the value the given key into the given interface{}
	Environment variables will override configutation
*/
func (loader *ConfigLoader) ParseInto(key string, out interface{}) error {
	value, exists := os.LookupEnv(key)

	if exists {
		return json.Unmarshal([]byte(value), out)
	}

	rawValue, exists := loader.parsedConfig[key]

	if !exists {
		return ErrNotSet
	}

	return json.Unmarshal(*rawValue, out)
}

/*
	Loads the data at a given s3 uri into a byte array
*/
func loadFromS3(configPath string) ([]byte, error) {
	uri, err := url.Parse(configPath)

	if err != nil {
		return nil, err
	}

	region, exists := os.LookupEnv("S3_REGION")

	if !exists {
		region = "us-east-1"
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	downloader := s3manager.NewDownloader(sess)

	buf := &aws.WriteAtBuffer{}

	_, err = downloader.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String(uri.Host),
		Key: aws.String(uri.Path[1 : ]),
	})

	return buf.Bytes(), err
}

/*
	Loads the data at a given file uri into a byte array
*/
func loadFromFile(configPath string) ([]byte, error) {
	uri, err := url.Parse(configPath)

	if err != nil {
		return nil, err
	}

	return ioutil.ReadFile(uri.Path)
}

/*
	Loads the data at a given https uri into a byte array
*/
func loadFromHttps(configPath string) ([]byte, error) {
	resp, err := http.Get(configPath)

	if err != nil{
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
