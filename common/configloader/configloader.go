package configloader

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

/*
	Used to load a key value configuration
*/
type ConfigLoader struct {
	configPath   string
	parsedConfig map[string]*json.RawMessage
}

/*
	Loads the configuration at the given path into a ConfigLoader struct
	Supported uri schemes are: s3, file, https
*/
func Load(config_path string) (*ConfigLoader, error) {
	uri, err := url.Parse(config_path)

	if err != nil {
		return nil, errors.New("Unable to parse '" + config_path + "' into URI structure")
	}

	var config_contents []byte

	switch uri.Scheme {
	case "s3":
		config_contents, err = loadFromS3(config_path)

		if err != nil {
			return nil, errors.New("Unable to load '" + config_path + "' as s3")
		}
	case "file":
		config_contents, err = loadFromFile(config_path)

		if err != nil {
			return nil, errors.New("Unable to load '" + config_path + "' as file")
		}
	case "https":
		config_contents, err = loadFromHttps(config_path)

		if err != nil {
			return nil, errors.New("Unable to load '" + config_path + "' as https")
		}
	default:
		return nil, errors.New("URI scheme of '" + config_path + "' unsupported. Must be s3, file, or https")
	}

	loader := ConfigLoader{
		configPath: config_path,
	}

	err = json.Unmarshal(config_contents, &loader.parsedConfig)

	if err != nil {
		return nil, errors.New("Unable to load contents of '" + config_path + "' into json")
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

	raw_value, exists := loader.parsedConfig[key]

	if !exists {
		return "", errors.New("Value for key '" + key + "' not set")
	}

	err := json.Unmarshal(*raw_value, &value)

	if err != nil {
		return "", errors.New("Value for key '" + key + "' couldn't be decoded")
	}

	return value, nil
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

	raw_value, exists := loader.parsedConfig[key]

	if !exists {
		return errors.New("Value for key '" + key + "' not set")
	}

	return json.Unmarshal(*raw_value, out)
}

/*
	Loads the data at a given s3 uri into a byte array
*/
func loadFromS3(config_path string) ([]byte, error) {
	uri, err := url.Parse(config_path)

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
		Key:    aws.String(uri.Path[1:]),
	})

	return buf.Bytes(), err
}

/*
	Loads the data at a given file uri into a byte array
*/
func loadFromFile(config_path string) ([]byte, error) {
	uri, err := url.Parse(config_path)

	if err != nil {
		return nil, err
	}

	return ioutil.ReadFile(uri.Path)
}

/*
	Loads the data at a given https uri into a byte array
*/
func loadFromHttps(config_path string) ([]byte, error) {
	resp, err := http.Get(config_path)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
