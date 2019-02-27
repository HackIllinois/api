package service

import (
	"errors"
	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/services/stat/config"
	"github.com/HackIllinois/api/services/stat/models"
	"net/http"
)

func Initialize() error {
	return nil
}

/*
	Retrieve stats from the specified service
*/
func GetAggregatedStats(service string) (*models.Stat, error) {
	endpoint, exists := config.STAT_ENDPOINTS[service]

	if !exists {
		return nil, errors.New("Could not find endpoint for requested statistics.")
	}

	var stat models.Stat
	status, err := apirequest.Get(endpoint, &stat)

	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New("Could not retrieve stats from service.")
	}

	return &stat, nil
}

/*
	Attempts to retrieve stats from the specified service and outputs this
	information to the given channel
*/
func GetAggregatedStatsAsync(service string, stat_chan chan models.AsyncStat) {
	stat, err := GetAggregatedStats(service)
	stat_chan <- models.AsyncStat{
		Service: service,
		Stat:    stat,
		Error:   err,
	}
}

/*
	Retreives stats from all services
	Returns a map of service name to stats
*/
func GetAllAggregatedStats() (*models.AggregatedStat, error) {
	stats := models.AggregatedStat{}

	stat_chan := make(chan models.AsyncStat)

	for service := range config.STAT_ENDPOINTS {
		go GetAggregatedStatsAsync(service, stat_chan)
	}

	for i := 0; i < len(config.STAT_ENDPOINTS); i++ {
		async_stat := <-stat_chan

		service := async_stat.Service
		stat := async_stat.Stat
		err := async_stat.Error

		if err == nil {
			stats[service] = *stat
		} else {
			stats[service] = nil
		}
	}

	return &stats, nil
}
