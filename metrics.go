package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	currentCount = 0

	httpHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "configstore_http_hit_total",
			Help: "Total number of http hits.",
		},
	)

	postConfigHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "configstore_post_config_hit_total",
			Help: "Total number of create config hits.",
		},
	)

	getConfigVerHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "configstore_get_config_ver_hit_total",
			Help: "Total number of get all config versions hits.",
		},
	)

	postConfigVerHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "configstore_post_config_ver_hit_total",
			Help: "Total number of add new config version hits.",
		},
	)

	getConfigHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "configstore_get_config_hit_total",
			Help: "Total number of get one config hits.",
		},
	)

	delConfigHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "configstore_del_config_hit_total",
			Help: "Total number of delete config hits.",
		},
	)

	postGroupHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "configstore_post_group_hit_total",
			Help: "Total number of post new group hits.",
		},
	)

	postGroupVerHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "configstore_post_group_ver_hit_total",
			Help: "Total number of adding new group version hits.",
		},
	)

	getGroupHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "configstore_get_group_hit_total",
			Help: "Total number of get group hits.",
		},
	)

	delGroupHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "configstore_del_group_hit_total",
			Help: "Total number of delete group hits.",
		},
	)

	getGroupConfigHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "configstore_get_group_config_hit_total",
			Help: "Total number of get group configs by label hits.",
		},
	)

	addGroupConfigHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "configstore_add_group_config_hit_total",
			Help: "Total number of add new config to a group hits.",
		},
	)

	metricsList = []prometheus.Collector{
		postConfigHits, getConfigVerHits, postConfigVerHits, getConfigHits,
		delConfigHits, postGroupHits, postGroupVerHits, getGroupHits, delGroupHits,
		getGroupConfigHits, addGroupConfigHits, httpHits,
	}

	prometheusRegistry = prometheus.NewRegistry()
)

func init() {
	prometheusRegistry.MustRegister(metricsList...)
}

func metricsHandler() http.Handler {
	return promhttp.HandlerFor(prometheusRegistry, promhttp.HandlerOpts{})
}

func countPostConfig(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		postConfigHits.Inc()
		f(w, r) // original function call
	}
}

func countGetConfigVer(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		getConfigVerHits.Inc()
		f(w, r) // original function call
	}
}

func countPostConfigVer(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		postConfigVerHits.Inc()
		f(w, r) // original function call
	}
}

func countGetConfig(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		getConfigHits.Inc()
		f(w, r) // original function call
	}
}

func countDelConfig(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		delConfigHits.Inc()
		f(w, r) // original function call
	}
}

func countPostGroup(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		postGroupHits.Inc()
		f(w, r) // original function call
	}
}

func countPostGroupVer(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		postGroupVerHits.Inc()
		f(w, r) // original function call
	}
}

func countGetGroup(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		getGroupHits.Inc()
		f(w, r) // original function call
	}
}

func countDelGroup(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		delGroupHits.Inc()
		f(w, r) // original function call
	}
}

func countGetGroupConfigs(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		getGroupConfigHits.Inc()
		f(w, r) // original function call
	}
}

func countAddGroupConfig(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		addGroupConfigHits.Inc()
		f(w, r) // original function call
	}
}
