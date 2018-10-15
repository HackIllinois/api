package service

import (
	"encoding/json"
	"errors"
	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/stat/config"
	"github.com/HackIllinois/api/services/stat/models"
	"net/http"
)

var db *database.MongoDatabase

func init() {
	db_connection, err := database.InitMongoDatabase(config.STAT_DB_HOST, config.STAT_DB_NAME)

	if err != nil {
		panic(err)
	}

	db = db_connection
}

/*
	Returns the service with the given name
*/
func GetService(name string) (*models.Service, error) {
	query := database.QuerySelector{
		"name": name,
	}

	var service models.Service
	err := db.FindOne("services", query, &service)

	if err != nil {
		return nil, err
	}

	return &service, nil
}

/*
	Registers the service with the given name
	The service will be created if it doesn't exist
*/
func RegisterService(name string, service models.Service) error {
	_, err := GetService(name)

	if err == database.ErrNotFound {
		err = db.Insert("services", &service)

		return err
	}

	if err != nil {
		return err
	}

	selector := database.QuerySelector{
		"name": name,
	}

	err = db.Update("services", selector, &service)

	return err
}

/*
	Returns all services that were registered
*/
func GetAllServices() ([]models.Service, error) {
	var services []models.Service
	err := db.FindAll("services", database.QuerySelector{}, &services)

	if err != nil {
		return nil, err
	}

	return services, nil
}

/*
	Retreive stats from a specified registered service
*/
func GetAggregatedStats(name string) (*models.Stat, error) {
	service, err := GetService(name)

	if err != nil {
		return nil, err
	}

	resp, err := http.Get(service.URL)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Could not retreive stats from registed service")
	}

	var stat models.Stat
	json.NewDecoder(resp.Body).Decode(&stat)

	return &stat, nil

}

/*
	Retreives stats from the registered services
	Returns a map of service name to stats
*/
func GetAllAggregatedStats() (*models.AggregatedStat, error) {
	stats := models.AggregatedStat{}

	services, err := GetAllServices()

	if err != nil {
		return nil, err
	}

	for _, service := range services {
		stat, err := GetAggregatedStats(service.Name)

		if err == nil {
			stats[service.Name] = *stat
		} else {
			stats[service.Name] = nil
		}
	}

	return &stats, nil
}
