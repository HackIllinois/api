package apiserver

import (
	"github.com/thoas/stats"
)

var HEALTHY_RESPONSE_TIME float64 = 3.0
var HEALTHY_OK_RESPONSE_PERCENT float64 = 40.0
var MINIMUM_NUM_RESPONSES int = 100

/*
	Determines if a service is healthy based on current health stats
*/
func IsHealthy(health_stats *stats.Data) bool {

	// Not enough data to mark service as unhealthy
	if health_stats.TotalCount < MINIMUM_NUM_RESPONSES {
		return true
	}

	if health_stats.AverageResponseTimeSec > HEALTHY_RESPONSE_TIME {
		return false
	}

	ok_response_count := 0
	total_response_count := 0

	for status_code, count := range health_stats.TotalStatusCodeCount {
		total_response_count += count

		if len(status_code) > 0 && (status_code[0] == '2' || status_code[0] == '3') {
			ok_response_count += count
		}
	}

	if float64(ok_response_count)/float64(total_response_count) < HEALTHY_OK_RESPONSE_PERCENT {
		return false
	}

	return true
}
