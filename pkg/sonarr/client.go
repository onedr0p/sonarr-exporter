package sonarr

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/onedr0p/sonarr-exporter/pkg/metrics"
)

// Client struct is a Sonarr client to request an instance of a Sonarr
type Client struct {
	httpClient http.Client
	interval   time.Duration
	hostname   string
	apiKey     string
}

// NewClient method initializes a new Sonarr client.
func NewClient(hostname, apiKey string, interval time.Duration) *Client {
	return &Client{
		hostname: hostname,
		apiKey:   apiKey,
		interval: interval,
		httpClient: http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

// Scrape method logins and retrieves statistics from Sonarr JSON API
// and then pass them as Prometheus metrics.
func (c *Client) Scrape() {
	for range time.Tick(c.interval) {

		// System Status
		status := SystemStatus{}
		if err := c.apiRequest(fmt.Sprintf("%s/api/v3/%s", c.hostname, "system/status"), &status); err != nil {
			metrics.Status.WithLabelValues(c.hostname).Set(0.0)
			return
		} else if (SystemStatus{}) == status {
			metrics.Status.WithLabelValues(c.hostname).Set(0.0)
			return
		} else {
			metrics.Status.WithLabelValues(c.hostname).Set(1.0)
		}

		// Series, Seasons, and Episodes
		var episodeFileSize int64

		var (
			totalSeasons      = 0
			totalEpisodes     = 0
			seriesMonitored   = 0
			seriesUnmonitored = 0
			episodeQualities  = map[string]int{}
		)
		series := Series{}
		c.apiRequest(fmt.Sprintf("%s/api/v3/%s", c.hostname, "series"), &series)
		for _, s := range series {
			if s.Monitored {
				seriesMonitored++
			} else {
				seriesUnmonitored++
			}
			totalSeasons += s.SeasonCount
			totalEpisodes += s.EpisodeCount

			// Get Episode Qualities
			episodeFile := EpisodeFile{}
			c.apiRequest(fmt.Sprintf("%s/api/v3/%s?seriesId=%d", c.hostname, "episodefile", s.Id), &episodeFile)
			for _, e := range episodeFile {
				if e.Quality.Quality.Name != "" {
					episodeQualities[e.Quality.Quality.Name]++
				}
				if e.Size != 0 {
					episodeFileSize += e.Size
				}
			}
		}
		metrics.Series.WithLabelValues(c.hostname).Set(float64(len(series)))
		metrics.Seasons.WithLabelValues(c.hostname).Set(float64(totalSeasons))
		metrics.Episodes.WithLabelValues(c.hostname).Set(float64(totalEpisodes))
		metrics.SeriesMonitored.WithLabelValues(c.hostname).Set(float64(seriesMonitored))
		metrics.SeriesUnmonitored.WithLabelValues(c.hostname).Set(float64(seriesUnmonitored))

		for qualityName, count := range episodeQualities {
			metrics.EpisodeQualities.WithLabelValues(c.hostname, qualityName).Set(float64(count))
		}

		metrics.EpisodeFileSize.WithLabelValues(c.hostname).Set(float64(episodeFileSize))

		// History
		history := History{}
		c.apiRequest(fmt.Sprintf("%s/api/v3/%s", c.hostname, "history"), &history)
		metrics.History.WithLabelValues(c.hostname).Set(float64(history.TotalRecords))

		// Wanted
		wanted := WantedMissing{}
		c.apiRequest(fmt.Sprintf("%s/api/v3/%s", c.hostname, "wanted/missing"), &wanted)
		metrics.Wanted.WithLabelValues(c.hostname).Set(float64(wanted.TotalRecords))

		// Queue by Status
		var queueStatus = map[string]int{}
		queue := Queue{}
		c.apiRequest(fmt.Sprintf("%s/api/v3/%s", c.hostname, "queue"), &queue)
		for _, s := range queue {
			if s.TrackedDownloadStatus != "" {
				queueStatus[s.TrackedDownloadStatus]++
			}
		}
		for trackedDownloadStatus, count := range queueStatus {
			metrics.Queue.WithLabelValues(c.hostname, trackedDownloadStatus).Set(float64(count))
		}

		// Root Folder
		rootFolders := RootFolder{}
		c.apiRequest(fmt.Sprintf("%s/api/v3/%s", c.hostname, "rootfolder"), &rootFolders)
		for _, rootFolder := range rootFolders {
			metrics.RootFolder.WithLabelValues(c.hostname, rootFolder.Path).Set(float64(rootFolder.FreeSpace))
		}

		// Health Issues
		health := Health{}
		c.apiRequest(fmt.Sprintf("%s/api/v3/%s", c.hostname, "health"), &health)
		for _, h := range health {
			metrics.Health.WithLabelValues(c.hostname, h.Source, h.Type, h.Message, h.WikiURL).Set(float64(1))
		}
	}
}

func (c *Client) apiRequest(endpoint string, target interface{}) error {
	log.Printf("Sending HTTP request to %s", endpoint)

	req, err := http.NewRequest("GET", endpoint, nil)
	// req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("")))
	req.Header.Set("X-Api-Key", c.apiKey)
	if err != nil {
		log.Fatal("An error has occurred when creating HTTP statistics request", err)
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatal("An error has occurred during retrieving Sonarr statistics", err)
		return err
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}
