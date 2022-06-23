package tests

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/project/config"
	"github.com/HackIllinois/api/services/project/models"
	"github.com/HackIllinois/api/services/project/service"
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
		ID:          "testid",
		Name:        "testname",
		Description: "testdesc",
		Mentors:     []string{"testmentor"},
		Number:      1,
		Tags:        []string{"tag1", "tag2"},
		Room:        "testroom",
	}

	err := db.Insert("projects", &project, nil)

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
	Service level test for getting all projects from db
*/
func TestGetAllProjectsService(t *testing.T) {
	SetupTestDB(t)

	project := models.Project{
		ID:          "testid2",
		Name:        "testname2",
		Description: "testdesc2",
		Mentors:     []string{"testmentor2"},
		Number:      2,
		Tags:        []string{"tag2"},
		Room:        "testroom2",
	}

	err := db.Insert("projects", &project, nil)

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
				ID:          "testid",
				Name:        "testname",
				Description: "testdesc",
				Mentors:     []string{"testmentor"},
				Number:      1,
				Tags:        []string{"tag1", "tag2"},
				Room:        "testroom",
			},
			{
				ID:          "testid2",
				Name:        "testname2",
				Description: "testdesc2",
				Mentors:     []string{"testmentor2"},
				Number:      2,
				Tags:        []string{"tag2"},
				Room:        "testroom2",
			},
		},
	}

	if !reflect.DeepEqual(actual_project_list, &expected_project_list) {
		t.Errorf("Wrong project list. Expected %v, got %v", expected_project_list, actual_project_list)
	}

	db.RemoveAll("projects", nil, nil)

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
		ID:          "testid2",
		Name:        "testname2",
		Description: "testdesc2",
		Mentors:     []string{"testmentor2"},
		Number:      2,
		Tags:        []string{"tag2"},
		Room:        "testroom2",
	}

	err := db.Insert("projects", &project, nil)

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
				ID:          "testid2",
				Name:        "testname2",
				Description: "testdesc2",
				Mentors:     []string{"testmentor2"},
				Number:      2,
				Tags:        []string{"tag2"},
				Room:        "testroom2",
			},
		},
	}

	if !reflect.DeepEqual(actual_project_list, &expected_project_list) {
		t.Errorf("Wrong project list. Expected %v, got %v", expected_project_list, actual_project_list)
	}

	// Filter to one project using number
	parameters = map[string][]string{
		"number": {"2"},
	}
	actual_project_list, err = service.GetFilteredProjects(parameters)

	if err != nil {
		t.Fatal(err)
	}

	expected_project_list = models.ProjectList{
		Projects: []models.Project{
			{
				ID:          "testid2",
				Name:        "testname2",
				Description: "testdesc2",
				Mentors:     []string{"testmentor2"},
				Number:      2,
				Tags:        []string{"tag2"},
				Room:        "testroom2",
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
				ID:          "testid",
				Name:        "testname",
				Description: "testdesc",
				Mentors:     []string{"testmentor"},
				Number:      1,
				Tags:        []string{"tag1", "tag2"},
				Room:        "testroom",
			},
			{
				ID:          "testid2",
				Name:        "testname2",
				Description: "testdesc2",
				Mentors:     []string{"testmentor2"},
				Number:      2,
				Tags:        []string{"tag2"},
				Room:        "testroom2",
			},
		},
	}

	if !reflect.DeepEqual(actual_project_list, &expected_project_list) {
		t.Errorf("Wrong project list. Expected %v, got %v", expected_project_list, actual_project_list)
	}

	db.RemoveAll("projects", nil, nil)

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
		ID:          "testid",
		Name:        "testname",
		Description: "testdesc",
		Mentors:     []string{"testmentor"},
		Number:      1,
		Tags:        []string{"tag1", "tag2"},
		Room:        "testroom",
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
		ID:          "testid2",
		Name:        "testname2",
		Description: "testdesc2",
		Mentors:     []string{"testmentor2"},
		Number:      5,
		Tags:        []string{"tag2"},
		Room:        "testroom2",
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
		ID:          "testid2",
		Name:        "testname2",
		Description: "testdesc2",
		Mentors:     []string{"testmentor2"},
		Number:      5,
		Tags:        []string{"tag2"},
		Room:        "testroom2",
	}

	if !reflect.DeepEqual(project, &expected_project) {
		t.Errorf("Wrong project info. Expected %v, got %v", expected_project, project)
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
		ID:          "testid",
		Name:        "testname2",
		Description: "testdesc2",
		Mentors:     []string{"testmentor", "testmentor2"},
		Number:      3,
		Tags:        []string{"tag1", "tag3"},
		Room:        "testroom2",
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
		ID:          "testid",
		Name:        "testname2",
		Description: "testdesc2",
		Mentors:     []string{"testmentor", "testmentor2"},
		Number:      3,
		Tags:        []string{"tag1", "tag3"},
		Room:        "testroom2",
	}

	if !reflect.DeepEqual(updated_project, &expected_project) {
		t.Errorf("Wrong project info. Expected %v, got %v", expected_project, updated_project)
	}

	CleanupTestDB(t)
}

/*
	Tests that getting project favorites works correctly at the service level
*/
func TestGetProjectFavorites(t *testing.T) {
	SetupTestDB(t)

	project_favorites, err := service.GetProjectFavorites("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_project_favorites := models.ProjectFavorites{
		ID:       "testid",
		Projects: []string{},
	}

	if !reflect.DeepEqual(project_favorites, &expected_project_favorites) {
		t.Errorf("Wrong tracker info. Expected %v, got %v", &expected_project_favorites, project_favorites)
	}

	CleanupTestDB(t)
}

/*
	Tests that adding project favorites works correctly at the service level
*/
func TestAddProjectFavorite(t *testing.T) {
	SetupTestDB(t)

	err := service.AddProjectFavorite("testid", "testid")

	if err != nil {
		t.Fatal(err)
	}

	project_favorites, err := service.GetProjectFavorites("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_project_favorites := models.ProjectFavorites{
		ID:       "testid",
		Projects: []string{"testid"},
	}

	if !reflect.DeepEqual(project_favorites, &expected_project_favorites) {
		t.Errorf("Wrong tracker info. Expected %v, got %v", &expected_project_favorites, project_favorites)
	}

	CleanupTestDB(t)
}

/*
	Tests that removing project favorites works correctly at the service level
*/
func TestRemoveProjectFavorite(t *testing.T) {
	SetupTestDB(t)

	err := service.AddProjectFavorite("testid", "testid")

	if err != nil {
		t.Fatal(err)
	}

	err = service.RemoveProjectFavorite("testid", "testid")

	if err != nil {
		t.Fatal(err)
	}

	project_favorites, err := service.GetProjectFavorites("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_project_favorites := models.ProjectFavorites{
		ID:       "testid",
		Projects: []string{},
	}

	if !reflect.DeepEqual(project_favorites, &expected_project_favorites) {
		t.Errorf("Wrong tracker info. Expected %v, got %v", &expected_project_favorites, project_favorites)
	}

	CleanupTestDB(t)
}
