package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	Status = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "status",
			Namespace: "sonarr",
			Help:      "System Status",
		},
		[]string{"hostname"},
	)

	Series = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "series_total",
			Namespace: "sonarr",
			Help:      "Total number of series",
		},
		[]string{"hostname"},
	)

	Seasons = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "seasons_total",
			Namespace: "sonarr",
			Help:      "Total number of seasons for all series",
		},
		[]string{"hostname"},
	)

	Episodes = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "episodes_total",
			Namespace: "sonarr",
			Help:      "Total number of episodes downloaded for all series",
		},
		[]string{"hostname"},
	)

	// SeriesMonitored - Total number of Series monitored
	SeriesMonitored = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "series_monitored_total",
			Namespace: "sonarr",
			Help:      "Total number of monitored series",
		},
		[]string{"hostname"},
	)

	// SeriesUnmonitored - Total number of Seriess unmonitored
	SeriesUnmonitored = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "series_unmonitored_total",
			Namespace: "sonarr",
			Help:      "Total number of unmonitored series",
		},
		[]string{"hostname"},
	)

	History = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "history_total",
			Namespace: "sonarr",
			Help:      "Total number of records in history",
		},
		[]string{"hostname"},
	)

	Wanted = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "missing_episodes_total",
			Namespace: "sonarr",
			Help:      "Total number of missing episodes",
		},
		[]string{"hostname"},
	)

	Queue = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "queue_total",
			Namespace: "sonarr",
			Help:      "Total number of episodes in queue",
		},
		[]string{"hostname"},
	)

	RootFolder = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "rootfolder_freespace_bytes",
			Namespace: "sonarr",
			Help:      "Root folder space in bytes",
		},
		[]string{"hostname", "folder"},
	)

	Health = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "health_issues",
			Namespace: "sonarr",
			Help:      "Amount of health issues in Sonarr",
		},
		[]string{"hostname", "type"},
	)
)

// Init initializes all Prometheus metrics made available by Sonarr Exporter.
func Init() {
	prometheus.MustRegister(Status)
	prometheus.MustRegister(Series)
	prometheus.MustRegister(Seasons)
	prometheus.MustRegister(Episodes)
	prometheus.MustRegister(History)
	prometheus.MustRegister(Wanted)
	prometheus.MustRegister(Queue)
	prometheus.MustRegister(RootFolder)
	prometheus.MustRegister(Health)
}
