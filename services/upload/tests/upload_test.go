package tests

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/upload/config"
	"github.com/HackIllinois/api/services/upload/models"
	"github.com/HackIllinois/api/services/upload/service"
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

	db, err = database.InitDatabase(config.UPLOAD_DB_HOST, config.UPLOAD_DB_NAME)

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	return_code := m.Run()

	os.Exit(return_code)
}

/*
	Initialize database with test upload info
*/
func SetupTestDB(t *testing.T) {
	err := db.Insert("blobstore", &models.Blob{
		ID:   "testid",
		Data: "testdata",
	}, nil)

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Drop test db
*/
func CleanupTestDB(t *testing.T) {
	err := db.DropDatabase(nil)

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Service level test for getting a blob from db
*/
func TestGetBlobService(t *testing.T) {
	SetupTestDB(t)

	blob, err := service.GetBlob("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_blob := &models.Blob{
		ID:   "testid",
		Data: "testdata",
	}

	if !reflect.DeepEqual(blob, expected_blob) {
		t.Errorf("Wrong blob. Expected %v, got %v", expected_blob, blob)
	}

	CleanupTestDB(t)
}

/*
	Service level test for creating a blob
*/
func TestCreateBlobService(t *testing.T) {
	SetupTestDB(t)

	new_blob := models.Blob{
		ID:   "testid2",
		Data: "testdata2",
	}

	err := service.CreateBlob(new_blob)

	if err != nil {
		t.Fatal(err)
	}

	retrieved_blob, err := service.GetBlob("testid2")

	if err != nil {
		t.Fatal(err)
	}

	expected_blob := &models.Blob{
		ID:   "testid2",
		Data: "testdata2",
	}

	if !reflect.DeepEqual(retrieved_blob, expected_blob) {
		t.Errorf("Wrong blob. Expected %v, got %v", expected_blob, retrieved_blob)
	}

	CleanupTestDB(t)
}

/*
	Service level test for updating a blob
*/
func TestUpdateBlobService(t *testing.T) {
	SetupTestDB(t)

	updated_blob := models.Blob{
		ID:   "testid",
		Data: "testdata2",
	}

	err := service.UpdateBlob(updated_blob)

	if err != nil {
		t.Fatal(err)
	}

	retrieved_blob, err := service.GetBlob("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_blob := &models.Blob{
		ID:   "testid",
		Data: "testdata2",
	}

	if !reflect.DeepEqual(retrieved_blob, expected_blob) {
		t.Errorf("Wrong blob. Expected %v, got %v", expected_blob, retrieved_blob)
	}

	CleanupTestDB(t)
}

func TestDeleteBlobService(t *testing.T) {
	SetupTestDB(t)

	new_blob := models.Blob{
		ID:   "testid2",
		Data: "testdata2",
	}

	err := service.CreateBlob(new_blob)

	if err != nil {
		t.Fatal(err)
	}

	expected_blob := &models.Blob{
		ID:   "testid2",
		Data: "testdata2",
	}

	returned_blob, err := service.DeleteBlob("testid2")

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(returned_blob, expected_blob) {
		t.Errorf("Wrong blob. Expected %v, got %v", expected_blob, returned_blob)
	}

	retrieved_blob, err := service.GetBlob("testid2")

	if retrieved_blob != nil || err.Error() != "Error: NOT_FOUND" {
		t.Errorf("Blob deletion failed")
	}

	CleanupTestDB(t)
}
