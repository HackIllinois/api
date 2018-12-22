package service

import (
	"errors"
	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/stat/config"
	"github.com/HackIllinois/api/services/stat/models"
	"net/http"
)

var db database.Database

func init() {
	db_connection, err := database.InitDatabase(config.STAT_DB_HOST, config.STAT_DB_NAME)

	if err != nil {
		panic(err)
	}

	db = db_connection
}

/*
	Retrieve stats from the specified service
*/
func GetAggregatedStats(service string) (*models.Stat, error) {
	endpoint, exists := config.STAT_ENDPOINTS[service]

	if !exists {
		return nil, errors.New("Could not find endpoint for requested stats")
	}

	var stat models.Stat
	status, err := apirequest.Get(endpoint, &stat)

	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New("Could not retreive stats from service")
	}

	return &stat, nil
}

/*
	Retreives stats from all services
	Returns a map of service name to stats
*/
func GetAllAggregatedStats() (*models.AggregatedStat, error) {
	stats := models.AggregatedStat{}

	for service, _ := range config.STAT_ENDPOINTS {
		stat, err := GetAggregatedStats(service)

		if err == nil {
			stats[service] = *stat
		} else {
			stats[service] = nil
		}
	}

	return &stats, nil
}
