package service

import (
	"errors"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/common/utils"
	"github.com/HackIllinois/api/services/project/config"
	"github.com/HackIllinois/api/services/project/models"
	"github.com/go-playground/validator/v10"
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
	err := db.FindOne("projects", query, &project, nil)

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

	err = db.RemoveOne("projects", query, nil)

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
	err := db.FindAll("projects", nil, &projects, nil)

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
	query, err := database.CreateFilterQuery(parameters, models.Project{})

	if err != nil {
		return nil, err
	}

	projects := []models.Project{}
	filtered_projects := models.ProjectList{Projects: projects}
	err = db.FindAll("projects", query, &filtered_projects.Projects, nil)

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

	err = db.Insert("projects", &project, nil)

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

	err = db.Replace("projects", selector, &project, false, nil)

	return err
}

/*
	Returns the project favorites for the user with the given id
*/
func GetProjectFavorites(id string) (*models.ProjectFavorites, error) {
	query := database.QuerySelector{
		"id": id,
	}

	var project_favorites models.ProjectFavorites
	err := db.FindOne("favorites", query, &project_favorites, nil)

	if err != nil {
		if err == database.ErrNotFound {
			err = db.Insert("favorites", &models.ProjectFavorites{
				ID:       id,
				Projects: []string{},
			}, nil)

			if err != nil {
				return nil, err
			}

			err = db.FindOne("favorites", query, &project_favorites, nil)

			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return &project_favorites, nil
}

/*
	Adds the given project to the favorites for the user with the given id
*/
func AddProjectFavorite(id string, project string) error {
	selector := database.QuerySelector{
		"id": id,
	}

	_, err := GetProject(project)

	if err != nil {
		return errors.New("Could not find project with the given id.")
	}

	project_favorites, err := GetProjectFavorites(id)

	if err != nil {
		return err
	}

	if !utils.ContainsString(project_favorites.Projects, project) {
		project_favorites.Projects = append(project_favorites.Projects, project)
	}

	err = db.Replace("favorites", selector, project_favorites, false, nil)

	return err
}

/*
	Removes the given project from the favorites of the user with the given id
*/
func RemoveProjectFavorite(id string, project string) error {
	selector := database.QuerySelector{
		"id": id,
	}

	project_favorites, err := GetProjectFavorites(id)

	if err != nil {
		return err
	}

	project_favorites.Projects, err = utils.RemoveString(project_favorites.Projects, project)

	if err != nil {
		return errors.New("User's project favorites does not have specified project")
	}

	err = db.Replace("favorites", selector, project_favorites, false, nil)

	return err
}
