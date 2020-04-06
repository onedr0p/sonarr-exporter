package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Status - System Status
	Status = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "status",
			Namespace: "sonarr",
			Help:      "System Status",
		},
		[]string{"hostname"},
	)

	// Series - Total number of Series
	Series = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "series_total",
			Namespace: "sonarr",
			Help:      "Total number of series",
		},
		[]string{"hostname"},
	)

	// Seasons - Total number of seasons
	Seasons = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "season_total",
			Namespace: "sonarr",
			Help:      "Total number of seasons for all series",
		},
		[]string{"hostname"},
	)

	// Episodes - Total number of downloaded episodes
	Episodes = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "episode_download_total",
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

	// EpisodeQualities - Total number of Episodes by Quality
	EpisodeQualities = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "episode_quality_total",
			Namespace: "sonarr",
			Help:      "Total number of downloaded episodes by quality",
		},
		[]string{"hostname", "quality"},
	)

	// EpisodeFileSize - Total size of all Episodes
	EpisodeFileSize = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "episode_bytes",
			Namespace: "sonarr",
			Help:      "Total file size of all episodes in bytes",
		},
		[]string{"hostname"},
	)

	// History - Total number of records in History
	History = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "history_total",
			Namespace: "sonarr",
			Help:      "Total number of records in history",
		},
		[]string{"hostname"},
	)

	// Wanted - Total number of missing/wanted Episodes
	Wanted = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "episode_missing_total",
			Namespace: "sonarr",
			Help:      "Total number of missing episodes",
		},
		[]string{"hostname"},
	)

	// Queue - Total number of episodes in Queue
	Queue = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "queue_total",
			Namespace: "sonarr",
			Help:      "Total number of episodes in queue",
		},
		[]string{"hostname"},
	)

	// RootFolder - Space by Root Folder
	RootFolder = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "rootfolder_freespace_bytes",
			Namespace: "sonarr",
			Help:      "Root folder space in bytes",
		},
		[]string{"hostname", "folder"},
	)

	// Health - Health issues with type and message
	Health = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "health_issues",
			Namespace: "sonarr",
			Help:      "Health issues in Sonarr",
		},
		[]string{"hostname", "type", "message", "wikiurl"},
	)
)

// Init initializes all Prometheus metrics made available by Sonarr Exporter.
func Init() {
	prometheus.MustRegister(Status)
	prometheus.MustRegister(Series)
	prometheus.MustRegister(Seasons)
	prometheus.MustRegister(Episodes)
	prometheus.MustRegister(SeriesMonitored)
	prometheus.MustRegister(SeriesUnmonitored)
	prometheus.MustRegister(EpisodeFileSize)
	prometheus.MustRegister(EpisodeQualities)
	prometheus.MustRegister(History)
	prometheus.MustRegister(Wanted)
	prometheus.MustRegister(Queue)
	prometheus.MustRegister(RootFolder)
	prometheus.MustRegister(Health)
}
