package service

import (
	"errors"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/recognition/config"
	"github.com/HackIllinois/api/services/recognition/models"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

var db database.Database

func Initialize() error {
	if db != nil {
		db.Close()
		db = nil
	}

	var err error
	db, err = database.InitDatabase(config.RECOGNITION_DB_HOST, config.RECOGNITION_DB_NAME)

	if err != nil {
		return err
	}

	validate = validator.New()

	return nil
}

/*
	Returns the recognition with the given id
*/
func GetRecognition(id string) (*models.Recognition, error) {
	query := database.QuerySelector{
		"id": id,
	}

	var recognition models.Recognition
	err := db.FindOne("recognitions", query, &recognition)

	if err != nil {
		return nil, err
	}

	return &recognition, nil
}


/*
	Returns all the recognitions
*/
func GetAllRecognitions() (*models.RecognitionList, error) {
	recognitions := []models.Recognition{}
	// nil implies there are no filters on the query, therefore everything in the "recognitions" collection is returned.
	err := db.FindAll("recognitions", nil, &recognitions)

	if err != nil {
		return nil, err
	}

	recognition_list := models.RecognitionList{
		Recognitions: recognitions,
	}

	return &recognition_list, nil
}


/*
	Creates an recognition with the given id
*/
func CreateRecognition(id string, recognition models.Recognition) error {
	err := validate.Struct(recognition)

	if err != nil {
		return err
	}

	_, err = GetRecognition(id)

	if err != database.ErrNotFound {
		if err != nil {
			return err
		}
		return errors.New("Recognition already exists")
	}

	err = db.Insert("recognitions", &recognition)

	if err != nil {
		return err
	}

	return err
}
