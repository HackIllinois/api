package config

import (
	"os"
)

var UPLOAD_DB_HOST = os.Getenv("UPLOAD_DB_HOST")
var UPLOAD_DB_NAME = os.Getenv("UPLOAD_DB_NAME")

var UPLOAD_PORT = os.Getenv("UPLOAD_PORT")

var S3_REGION = os.Getenv("S3_REGION")
var S3_BUCKET = os.Getenv("S3_BUCKET")
