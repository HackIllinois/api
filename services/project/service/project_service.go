package service

import (
	"errors"
	"strings"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/project/config"
	"github.com/HackIllinois/api/services/project/models"
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
	db, err = database.InitDatabase(config.PROJECT_DB_HOST, config.PROJECT_DB_NAME)

	if err != nil {
		return err
	}

	validate = validator.New()

	return nil
}

/*
	Returns the project with the given id
*/
func GetProject(id string) (*models.Project, error) {
	query := database.QuerySelector{
		"id": id,
	}

	var project models.Project
	err := db.FindOne("projects", query, &project)

	if err != nil {
		return nil, err
	}

	return &project, nil
}

/*
	Deletes the project with the given id.
	Removes the project from project trackers and every user's tracker.
	Returns the project that was deleted.
*/
func DeleteProject(id string) (*models.Project, error) {

	// Gets project to be able to return it later

	project, err := GetProject(id)

	if err != nil {
		return nil, err
	}

	query := database.QuerySelector{
		"id": id,
	}

	// Remove project from projects database

	err = db.RemoveOne("projects", query)

	if err != nil {
		return nil, err
	}

	return project, err
}

/*
	Returns all the projects
*/
func GetAllProjects() (*models.ProjectList, error) {
	projects := []models.Project{}
	// nil implies there are no filters on the query, therefore everything in the "projects" collection is returned.
	err := db.FindAll("projects", nil, &projects)

	if err != nil {
		return nil, err
	}

	project_list := models.ProjectList{
		Projects: projects,
	}

	return &project_list, nil
}

/*
	Returns all the projects
*/
func GetFilteredProjects(parameters map[string][]string) (*models.ProjectList, error) {
	query := make(map[string]interface{})

	for key, values := range parameters {
		if len(values) > 1 {
			return nil, errors.New("Multiple usage of key " + key)
		}

		key = strings.ToLower(key)
		query[key] = database.QuerySelector{"$in": strings.Split(values[0], ",")}
	}

	projects := []models.Project{}
	filtered_projects := models.ProjectList{Projects: projects}
	err := db.FindAll("projects", query, &filtered_projects.Projects)

	if err != nil {
		return nil, err
	}

	return &filtered_projects, nil
}

/*
	Creates a project with the given id
*/
func CreateProject(id string, project models.Project) error {
	err := validate.Struct(project)

	if err != nil {
		return err
	}

	_, err = GetProject(id)

	if err != database.ErrNotFound {
		if err != nil {
			return err
		}
		return errors.New("Project already exists")
	}

	err = db.Insert("projects", &project)

	return err
}

/*
	Updates the project with the given id
*/
func UpdateProject(id string, project models.Project) error {
	err := validate.Struct(project)

	if err != nil {
		return err
	}

	selector := database.QuerySelector{
		"id": id,
	}

	err = db.Update("projects", selector, &project)

	return err
}
