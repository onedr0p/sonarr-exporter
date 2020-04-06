package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// SystemStatus - System Status
	SystemStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "system_status",
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

	// Monitored - Total number of Series monitored
	Monitored = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "series_monitored_total",
			Namespace: "sonarr",
			Help:      "Total number of monitored series",
		},
		[]string{"hostname"},
	)

	// Unmonitored - Total number of Seriess unmonitored
	Unmonitored = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "series_unmonitored_total",
			Namespace: "sonarr",
			Help:      "Total number of unmonitored series",
		},
		[]string{"hostname"},
	)

	// Qualities - Total number of Episodes by Quality
	Qualities = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "episode_quality_total",
			Namespace: "sonarr",
			Help:      "Total number of downloaded episodes by quality",
		},
		[]string{"hostname", "quality"},
	)

	// FileSize - Total size of all Series
	FileSize = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "episode_bytes",
			Namespace: "sonarr",
			Help:      "Total filesize of all series in bytes",
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
		[]string{"hostname", "status", "download_status", "download_state"},
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

	// SystemHealth - System Health issues
	SystemHealth = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "system_health_issues",
			Namespace: "sonarr",
			Help:      "System Health issues in Sonarr",
		},
		[]string{"hostname", "type", "message", "wikiurl"},
	)
)

// Init initializes all Prometheus metrics made available by Sonarr Exporter.
func Init() {
	prometheus.MustRegister(SystemStatus)
	prometheus.MustRegister(Series)
	prometheus.MustRegister(Seasons)
	prometheus.MustRegister(Episodes)
	prometheus.MustRegister(Monitored)
	prometheus.MustRegister(Unmonitored)
	prometheus.MustRegister(FileSize)
	prometheus.MustRegister(Qualities)
	prometheus.MustRegister(History)
	prometheus.MustRegister(Wanted)
	prometheus.MustRegister(Queue)
	prometheus.MustRegister(RootFolder)
	prometheus.MustRegister(SystemHealth)
}
