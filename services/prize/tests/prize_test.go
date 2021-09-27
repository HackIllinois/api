package tests

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/prize/config"
	"github.com/HackIllinois/api/services/prize/models"
	"github.com/HackIllinois/api/services/prize/service"
)

var db database.Database

func TestMain(m *testing.M) {
	err := config.Initialize()

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)

	}

	err = service.Initialize()

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	db, err = database.InitDatabase(config.PRIZE_DB_HOST, config.PRIZE_DB_NAME)

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	return_code := m.Run()

	os.Exit(return_code)
}

var TestTime = time.Now().Unix()

/*
	Initialize db with a test prize
*/
func SetupTestDB(t *testing.T) {
	user_id := "testid"

	prize := models.Prize{
		ID:       user_id,
		Name:     "testprizename",
		Value:    233,
		Quantity: 1337,
		ImageUrl: "https://imgs.smoothradio.com/images/191589?crop=16_9&width=660&relax=1&signature=Rz93ikqcAz7BcX6SKiEC94zJnqo=",
	}

	points := models.UserPoints{
		ID:     "testuserid",
		Points: 500,
	}

	err := db.Insert("userpoints", &points)

	if err != nil {
		t.Fatal(err)
	}

	err = db.Insert("prizes", &prize)

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Drop test db
*/
func CleanupTestDB(t *testing.T) {
	err := db.DropDatabase()

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Service level test for getting prize from db
*/
func TestGetPrizeService(t *testing.T) {
	SetupTestDB(t)

	prize, err := service.GetPrize("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_prize := models.Prize{
		ID:       "testid",
		Name:     "testprizename",
		Value:    233,
		Quantity: 1337,
		ImageUrl: "https://imgs.smoothradio.com/images/191589?crop=16_9&width=660&relax=1&signature=Rz93ikqcAz7BcX6SKiEC94zJnqo=",
	}

	if !reflect.DeepEqual(prize, &expected_prize) {
		t.Errorf("Wrong prize info. Expected %v, got %v", &expected_prize, prize)
	}

	CleanupTestDB(t)
}

/*
	Service level test for creating a prize in the db
*/
func TestCreatePrizeService(t *testing.T) {
	SetupTestDB(t)

	new_prize := models.Prize{
		ID:       "testid2",
		Name:     "testfirstname",
		Value:    5000,
		Quantity: 30,
		ImageUrl: "https://imgs.smoothradio.com/images/191589?crop=16_9&width=660&relax=1&signature=Rz93ikqcAz7BcX6SKiEC94zJnqo=",
	}

	err := service.CreatePrize(new_prize)

	if err != nil {
		t.Fatal(err)
	}

	prize, err := service.GetPrize(new_prize.ID)

	if err != nil {
		t.Fatal(err)
	}

	expected_prize := models.Prize{
		ID:       "testid2",
		Name:     "testfirstname",
		Value:    5000,
		Quantity: 30,
		ImageUrl: "https://imgs.smoothradio.com/images/191589?crop=16_9&width=660&relax=1&signature=Rz93ikqcAz7BcX6SKiEC94zJnqo=",
	}

	if !reflect.DeepEqual(prize, &expected_prize) {
		t.Errorf("Wrong prize info. Expected %v, got %v", expected_prize, prize)
	}

	CleanupTestDB(t)
}

/*
	Service level test for deleting a prize in the db
*/
func TestDeletePrizeService(t *testing.T) {
	SetupTestDB(t)

	prize_id := "testid"

	// Try to delete the profile

	err := service.DeletePrize(prize_id)

	if err != nil {
		t.Fatal(err)
	}

	// Try to find the profile in the profiles db
	prize, err := service.GetPrize(prize_id)

	if err == nil {
		t.Errorf("Found prize %v in prizes database.", prize)
	}

	CleanupTestDB(t)
}

/*
	Service level test for updating a profile in the db
*/
func TestUpdatePrizeService(t *testing.T) {
	SetupTestDB(t)

	prize := models.Prize{
		ID:       "testid",
		Name:     "another prize!",
		Value:    1234,
		Quantity: 9999,
		ImageUrl: "https://imgs.smoothradio.com/images/191589?crop=16_9&width=660&relax=1&signature=Rz93ikqcAz7BcX6SKiEC94zJnqo=",
	}

	err := service.UpdatePrize(prize.ID, prize)

	if err != nil {
		t.Fatal(err)
	}

	updated_prize, err := service.GetPrize(prize.ID)

	if err != nil {
		t.Fatal(err)
	}

	expected_prize := models.Prize{
		ID:       "testid",
		Name:     "another prize!",
		Value:    1234,
		Quantity: 9999,
		ImageUrl: "https://imgs.smoothradio.com/images/191589?crop=16_9&width=660&relax=1&signature=Rz93ikqcAz7BcX6SKiEC94zJnqo=",
	}

	if !reflect.DeepEqual(updated_prize, &expected_prize) {
		t.Errorf("Wrong prize info. Expected %v, got %v", expected_prize, updated_prize)
	}

	CleanupTestDB(t)
}

/*
	Service level test for redeeming a prize
*/
func TestAwardPointsService(t *testing.T) {
	SetupTestDB(t)

	updated_points, err := service.AwardPoints(100, "testuserid")

	if err != nil {
		t.Fatal(err)
	}

	expected_points := models.UserPoints{
		ID:     "testuserid",
		Points: 600,
	}

	if !reflect.DeepEqual(updated_points, &expected_points) {
		t.Errorf("Wrong prize info. Expected %v, got %v", expected_points, updated_points)
	}

	CleanupTestDB(t)
}

/*
	Service level test for redeeming a prize
*/
func TestRedeemPrizeService(t *testing.T) {
	SetupTestDB(t)

	err := service.RedeemPrize("testid", "testuserid")

	if err != nil {
		t.Fatal(err)
	}

	updated_prize, err := service.GetPrize("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_prize := models.Prize{
		ID:       "testid",
		Name:     "testprizename",
		Value:    233,
		Quantity: 1336,
		ImageUrl: "https://imgs.smoothradio.com/images/191589?crop=16_9&width=660&relax=1&signature=Rz93ikqcAz7BcX6SKiEC94zJnqo=",
	}

	if !reflect.DeepEqual(updated_prize, &expected_prize) {
		t.Errorf("Wrong prize info. Expected %v, got %v", expected_prize, updated_prize)
	}

	updated_points, err := service.GetPoints("testuserid")

	if err != nil {
		t.Fatal(err)
	}

	expected_points := models.UserPoints{
		ID:     "testuserid",
		Points: 267,
	}

	if !reflect.DeepEqual(updated_points, &expected_points) {
		t.Errorf("Wrong user points info. Expected %v, got %v", expected_points, updated_points)
	}

	CleanupTestDB(t)
}
