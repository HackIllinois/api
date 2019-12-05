package tests

import (
	"fmt"
	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/project/config"
	"github.com/HackIllinois/api/services/project/models"
	"github.com/HackIllinois/api/services/project/service"
	"os"
	"reflect"
	"testing"
	"time"
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

	db, err = database.InitDatabase(config.PROJECT_DB_HOST, config.PROJECT_DB_NAME)

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	return_code := m.Run()

	os.Exit(return_code)
}

var TestTime = time.Now().Unix()

/*
	Initialize db with a test project
*/
func SetupTestDB(t *testing.T) {
	project := models.Project{
		ID:      "testid",
		Name:    "testname",
		Mentors: []string{"testmentor"},
		Code:    "testcode",
		Tags:    []string{"tag1", "tag2"},
		Location: models.ProjectLocation{
			Description: "testlocationdescription",
			Latitude:    123.456,
			Longitude:   123.456,
		},
	}

	err := db.Insert("projects", &project)

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
	Service level test for getting all projects from db
*/
func TestGetAllProjectsService(t *testing.T) {
	SetupTestDB(t)

	project := models.Project{
		ID:      "testid2",
		Name:    "testname2",
		Mentors: []string{"testmentor2"},
		Code:    "testcode2",
		Tags:    []string{"tag2"},
		Location: models.ProjectLocation{
			Description: "testlocationdescription2",
			Latitude:    423.456,
			Longitude:   777.777,
		},
	}

	err := db.Insert("projects", &project)

	if err != nil {
		t.Fatal(err)
	}

	actual_project_list, err := service.GetAllProjects()

	if err != nil {
		t.Fatal(err)
	}

	expected_project_list := models.ProjectList{
		Projects: []models.Project{
			{
				ID:      "testid",
				Name:    "testname",
				Mentors: []string{"testmentor"},
				Code:    "testcode",
				Tags:    []string{"tag1", "tag2"},
				Location: models.ProjectLocation{
					Description: "testlocationdescription",
					Latitude:    123.456,
					Longitude:   123.456,
				},
			},
			{
				ID:      "testid2",
				Name:    "testname2",
				Mentors: []string{"testmentor2"},
				Code:    "testcode2",
				Tags:    []string{"tag2"},
				Location: models.ProjectLocation{
					Description: "testlocationdescription2",
					Latitude:    423.456,
					Longitude:   777.777,
				},
			},
		},
	}

	if !reflect.DeepEqual(actual_project_list, &expected_project_list) {
		t.Errorf("Wrong project list. Expected %v, got %v", expected_project_list, actual_project_list)
	}

	db.RemoveAll("projects", nil)

	actual_project_list, err = service.GetAllProjects()

	if err != nil {
		t.Fatal(err)
	}

	expected_project_list = models.ProjectList{
		Projects: []models.Project{},
	}

	if !reflect.DeepEqual(actual_project_list, &expected_project_list) {
		t.Errorf("Wrong project list. Expected %v, got %v", expected_project_list, actual_project_list)
	}

	CleanupTestDB(t)

}

/*
	Service level test for getting a filtered list of projects from the db
*/
func TestGetFilteredProjectsService(t *testing.T) {
	SetupTestDB(t)

	project := models.Project{
		ID:      "testid2",
		Name:    "testname2",
		Mentors: []string{"testmentor2"},
		Code:    "testcode2",
		Tags:    []string{"tag2"},
		Location: models.ProjectLocation{
			Description: "testlocationdescription2",
			Latitude:    423.456,
			Longitude:   777.777,
		},
	}

	err := db.Insert("projects", &project)

	if err != nil {
		t.Fatal(err)
	}

	// Filter to one project
	parameters := map[string][]string{
		"name": {"testname2"},
	}
	actual_project_list, err := service.GetFilteredProjects(parameters)

	if err != nil {
		t.Fatal(err)
	}

	expected_project_list := models.ProjectList{
		Projects: []models.Project{
			{
				ID:      "testid2",
				Name:    "testname2",
				Mentors: []string{"testmentor2"},
				Code:    "testcode2",
				Tags:    []string{"tag2"},
				Location: models.ProjectLocation{
					Description: "testlocationdescription2",
					Latitude:    423.456,
					Longitude:   777.777,
				},
			},
		},
	}

	if !reflect.DeepEqual(actual_project_list, &expected_project_list) {
		t.Errorf("Wrong project list. Expected %v, got %v", expected_project_list, actual_project_list)
	}

	// Filter to multiple (all) projects
	parameters = map[string][]string{}
	actual_project_list, err = service.GetFilteredProjects(parameters)

	if err != nil {
		t.Fatal(err)
	}

	expected_project_list = models.ProjectList{
		Projects: []models.Project{
			{
				ID:      "testid",
				Name:    "testname",
				Mentors: []string{"testmentor"},
				Code:    "testcode",
				Tags:    []string{"tag1", "tag2"},
				Location: models.ProjectLocation{
					Description: "testlocationdescription",
					Latitude:    123.456,
					Longitude:   123.456,
				},
			},
			{
				ID:      "testid2",
				Name:    "testname2",
				Mentors: []string{"testmentor2"},
				Code:    "testcode2",
				Tags:    []string{"tag2"},
				Location: models.ProjectLocation{
					Description: "testlocationdescription2",
					Latitude:    423.456,
					Longitude:   777.777,
				},
			},
		},
	}

	if !reflect.DeepEqual(actual_project_list, &expected_project_list) {
		t.Errorf("Wrong project list. Expected %v, got %v", expected_project_list, actual_project_list)
	}

	db.RemoveAll("projects", nil)

	// Filter again, with no projects remaining
	actual_project_list, err = service.GetFilteredProjects(parameters)

	if err != nil {
		t.Fatal(err)
	}

	expected_project_list = models.ProjectList{
		Projects: []models.Project{},
	}

	if !reflect.DeepEqual(actual_project_list, &expected_project_list) {
		t.Errorf("Wrong project list. Expected %v, got %v", expected_project_list, actual_project_list)
	}

	CleanupTestDB(t)

}

/*
	Service level test for getting project from db
*/
func TestGetProjectService(t *testing.T) {
	SetupTestDB(t)

	project, err := service.GetProject("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_project := models.Project{
		ID:      "testid",
		Name:    "testname",
		Mentors: []string{"testmentor"},
		Code:    "testcode",
		Tags:    []string{"tag1", "tag2"},
		Location: models.ProjectLocation{
			Description: "testlocationdescription",
			Latitude:    123.456,
			Longitude:   123.456,
		},
	}

	if !reflect.DeepEqual(project, &expected_project) {
		t.Errorf("Wrong project info. Expected %v, got %v", &expected_project, project)
	}

	CleanupTestDB(t)
}

/*
	Service level test for creating a project in the db
*/
func TestCreateProjectService(t *testing.T) {
	SetupTestDB(t)

	new_project := models.Project{
		ID:      "testid2",
		Name:    "testname2",
		Mentors: []string{"testmentor2"},
		Code:    "testcode2",
		Tags:    []string{"tag2"},
		Location: models.ProjectLocation{
			Description: "testlocationdescription2",
			Latitude:    423.456,
			Longitude:   777.777,
		},
	}

	err := service.CreateProject("testid2", new_project)

	if err != nil {
		t.Fatal(err)
	}

	project, err := service.GetProject("testid2")

	if err != nil {
		t.Fatal(err)
	}

	expected_project := models.Project{
		ID:      "testid2",
		Name:    "testname2",
		Mentors: []string{"testmentor2"},
		Code:    "testcode2",
		Tags:    []string{"tag2"},
		Location: models.ProjectLocation{
			Description: "testlocationdescription2",
			Latitude:    423.456,
			Longitude:   777.777,
		},
	}

	if !reflect.DeepEqual(project, &expected_project) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_project, project)
	}

	CleanupTestDB(t)
}

/*
	Service level test for deleting a project in the db
*/
func TestDeleteProjectService(t *testing.T) {
	SetupTestDB(t)

	project_id := "testid"

	// Try to delete the project

	_, err := service.DeleteProject(project_id)

	if err != nil {
		t.Fatal(err)
	}

	// Try to find the project in the projects db
	project, err := service.GetProject(project_id)

	if err == nil {
		t.Errorf("Found project %v in projects database.", project)
	}

	CleanupTestDB(t)
}

/*
	Service level test for updating a project in the db
*/
func TestUpdateProjectService(t *testing.T) {
	SetupTestDB(t)

	project := models.Project{
		ID:      "testid",
		Name:    "testname2",
		Mentors: []string{"testmentor", "testmentor2"},
		Code:    "testcode2",
		Tags:    []string{"tag1", "tag3"},
		Location: models.ProjectLocation{
			Description: "testlocationdescription2",
			Latitude:    444.456,
			Longitude:   123.456,
		},
	}

	err := service.UpdateProject("testid", project)

	if err != nil {
		t.Fatal(err)
	}

	updated_project, err := service.GetProject("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_project := models.Project{
		ID:      "testid",
		Name:    "testname2",
		Mentors: []string{"testmentor", "testmentor2"},
		Code:    "testcode2",
		Tags:    []string{"tag1", "tag3"},
		Location: models.ProjectLocation{
			Description: "testlocationdescription2",
			Latitude:    444.456,
			Longitude:   123.456,
		},
	}

	if !reflect.DeepEqual(updated_project, &expected_project) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_project, updated_project)
	}

	CleanupTestDB(t)
}
