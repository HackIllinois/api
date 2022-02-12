package controller

import (
	"encoding/json"
	"net/http"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/common/metrics"
	"github.com/HackIllinois/api/common/utils"
	"github.com/HackIllinois/api/services/project/models"
	"github.com/HackIllinois/api/services/project/service"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/internal/metrics/", promhttp.Handler()).Methods("GET")

	metrics.RegisterHandler("/favorite/", GetProjectFavorites, "GET", router)
	metrics.RegisterHandler("/favorite/", AddProjectFavorite, "POST", router)
	metrics.RegisterHandler("/favorite/", RemoveProjectFavorite, "DELETE", router)

	metrics.RegisterHandler("/filter/", GetFilteredProjects, "GET", router)
	metrics.RegisterHandler("/{id}/", GetProject, "GET", router)
	metrics.RegisterHandler("/{id}/", DeleteProject, "DELETE", router)
	metrics.RegisterHandler("/", CreateProject, "POST", router)
	metrics.RegisterHandler("/", UpdateProject, "PUT", router)
	metrics.RegisterHandler("/", GetAllProjects, "GET", router)
}

/*
	Endpoint to get the current user's project favorites
*/
func GetProjectFavorites(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	favorites, err := service.GetProjectFavorites(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get user's project favourites."))
		return
	}

	json.NewEncoder(w).Encode(favorites)
}

/*
	Endpoint to add a project favorite for the current user
*/
func AddProjectFavorite(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	var project_favorite_modification models.ProjectFavoriteModification
	json.NewDecoder(r.Body).Decode(&project_favorite_modification)

	err := service.AddProjectFavorite(id, project_favorite_modification.ProjectID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not add a project favorite for the current user."))
		return
	}

	favorites, err := service.GetProjectFavorites(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get updated user project favorites."))
		return
	}

	json.NewEncoder(w).Encode(favorites)
}

/*
	Endpoint to remove a project favorite for the current user
*/
func RemoveProjectFavorite(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	var project_favorite_modification models.ProjectFavoriteModification
	json.NewDecoder(r.Body).Decode(&project_favorite_modification)

	err := service.RemoveProjectFavorite(id, project_favorite_modification.ProjectID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not remove a project favorite for the current user."))
		return
	}

	favorites, err := service.GetProjectFavorites(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not fetch updated project favourites for the user (post-removal)."))
		return
	}

	json.NewEncoder(w).Encode(favorites)
}

/*
	Endpoint to get the project with the specified id
*/
func GetProject(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	project, err := service.GetProject(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not fetch the project details."))
		return
	}

	json.NewEncoder(w).Encode(project)
}

/*
	Endpoint to delete a project with the specified id.
	It removes the project from the project trackers, and every user's tracker.
	On successful deletion, it returns the project that was deleted.
*/
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	project, err := service.DeleteProject(id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not delete either the project, project trackers, or user trackers, or an intermediary subroutine failed."))
		return
	}

	json.NewEncoder(w).Encode(project)
}

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	project_list, err := service.GetAllProjects()

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get all projects."))
		return
	}

	json.NewEncoder(w).Encode(project_list)
}

/*
	Endpoint to get projects based on filters
*/
func GetFilteredProjects(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	project, err := service.GetFilteredProjects(parameters)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not fetch filtered list of projects."))
		return
	}

	json.NewEncoder(w).Encode(project)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	json.NewDecoder(r.Body).Decode(&project)

	project.ID = utils.GenerateUniqueID()

	err := service.CreateProject(project.ID, project)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not create new project."))
		return
	}

	updated_project, err := service.GetProject(project.ID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get updated project."))
		return
	}

	json.NewEncoder(w).Encode(updated_project)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	json.NewDecoder(r.Body).Decode(&project)

	err := service.UpdateProject(project.ID, project)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not update the project."))
		return
	}

	updated_project, err := service.GetProject(project.ID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get updated project details."))
		return
	}

	json.NewEncoder(w).Encode(updated_project)
}
