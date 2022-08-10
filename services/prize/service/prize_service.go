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
	err := db.FindOne("prizes", query, &prize)

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

	err = db.Insert("prizes", &prize)

	return err
}

/*
	Creates a prize
*/
func UpdatePrize(prize_id string, prize models.Prize) error {
	err := validate.Struct(prize)

	if err != nil {
		return err
	}

	selector := database.QuerySelector{
		"id": prize_id,
	}

	err = db.Update("prizes", selector, &prize)

	return err
}

func GetPoints(user_id string) (*models.UserPoints, error) {
	selector := database.QuerySelector{
		"id": user_id,
	}

	var points models.UserPoints
	err := db.FindOne("userpoints", selector, &points)

	if err != nil {
		return nil, errors.New("Could not find given user id")
	}

	return &points, nil
}

func DeletePrize(prize_id string) error {

	selector := database.QuerySelector{
		"id": prize_id,
	}

	err := db.RemoveOne("prizes", selector)

	return err
}

func CreateUserPoints(user_id string, user_points *models.UserPoints) error {
	selector := database.QuerySelector{
		"id": user_id,
	}

	var user_pts models.UserPoints
	err := db.FindOne("userpoints", selector, &user_pts)

	if err != database.ErrNotFound {
		if err != nil {
			return err
		}
		return errors.New("User points already exists")
	}

	user_points.ID = user_id
	user_points.Points = 0

	err = db.Insert("userpoints", &user_points)
	if err != nil {
		return errors.New("Could not create user points")
	}

	return nil
}

func AwardPoints(points int, user_id string) (*models.UserPoints, error) {
	selector := database.QuerySelector{
		"id": user_id,
	}

	var user_points models.UserPoints
	err := db.FindOne("userpoints", selector, &user_points)

	if err != nil {
		// Could not find given user id, assuming zero points and creating new entry
		err = CreateUserPoints(user_id, &user_points)

		if err != nil {
			return nil, errors.New("Could not find or create user points")
		}
	}

	user_points.Points += points

	err = db.Update("userpoints", selector, &user_points)

	if err != nil {
		return nil, errors.New("Could not update user points")
	}

	return &user_points, nil
}

func RedeemPrize(prize_item_id string, user_id string) error {
	prize_query := database.QuerySelector{
		"id": prize_item_id,
	}

	var prize models.Prize
	err := db.FindOne("prizes", prize_query, &prize)

	if err != nil {
		return errors.New("Item trying to redeem doesn't exist")
	}

	points_query := database.QuerySelector{
		"id": user_id,
	}

	var user_points models.UserPoints
	err = db.FindOne("userpoints", points_query, &user_points)

	if err != nil {
		return errors.New("User does not have any existing points")
	}

	if prize.Quantity <= 0 {
		return errors.New("Item is out of stock")
	}

	if prize.Value > user_points.Points {
		return errors.New("User has insufficient points")
	}

	user_points.Points -= prize.Value // Note: prize.Value can be negative
	prize.Quantity -= 1

	// TODO: Associate redeemed items with profile

	err = db.Update("prizes", prize_query, &prize)

	if err != nil {
		return errors.New("Prize item could not be updated")
	}

	selector := database.QuerySelector{
		"id": user_points.ID,
	}
	// A rare case, but if this update fails, the user is not charged for the
	//  item and the quantity of the prize item decrements since they are
	//  individual commits.
	err = db.Update("userpoints", selector, user_points)

	if err != nil {
		return errors.New("Profile could not be updated")
	}

	return nil
}
