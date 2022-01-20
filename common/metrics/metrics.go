package metrics

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func metricsWrapper(next http.HandlerFunc, endpoint string) http.Handler {
	labels := map[string]string{
		"endpoint": endpoint,
	}
	return promhttp.InstrumentHandlerDuration(Duration.MustCurryWith(labels),
		promhttp.InstrumentHandlerCounter(TotalRequests.MustCurryWith(labels),
			next))
}

func RegisterHandler(endpoint string, f func(http.ResponseWriter, *http.Request), method string, router *mux.Router) {
	chain := metricsWrapper(http.HandlerFunc(f), endpoint)
	router.Handle(endpoint, chain).Methods(method)
}
