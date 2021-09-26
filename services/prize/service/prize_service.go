package service

import (
	"errors"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/prize/config"
	"github.com/HackIllinois/api/services/prize/models"
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
	db, err = database.InitDatabase(config.PRIZE_DB_HOST, config.PRIZE_DB_NAME)

	if err != nil {
		return err
	}

	validate = validator.New()

	return nil
}

/*
	TODO: Returns all prizes
*/
func GetAllPrizes() {

}

/*
	Returns the prize with the given id
*/
func GetPrize(prize_id string) (*models.Prize, error) {
	query := database.QuerySelector{
		"id": prize_id,
	}

	var prize models.Prize
	err := db.FindOne("prize", query, &prize)

	if err != nil {
		return nil, err
	}

	return &prize, nil
}

/*
	Creates a prize
*/
func CreatePrize(prize models.Prize) error {
	_, err := GetPrize(prize.ID)

	if err != database.ErrNotFound {
		if err != nil {
			return err
		}
		return errors.New("Prize already exists")
	}

	err = db.Insert("prize", &prize)

	return err
}

/*
	Creates a prize
*/
func UpdatePrize(prize models.Prize) error {
	err := validate.Struct(prize)

	if err != nil {
		return err
	}

	selector := database.QuerySelector{
		"id": prize.ID,
	}

	err = db.Update("prize", selector, &prize)

	return err
}

func DeletePrize(prize_id string) error {

	selector := database.QuerySelector{
		"id": prize_id,
	}

	err := db.RemoveOne("prize", selector)

	return err
}
