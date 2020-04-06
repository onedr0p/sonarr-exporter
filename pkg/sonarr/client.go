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
		var sizeOnDisk int64

		var (
			totalSeasons     = 0
			totalEpisodes    = 0
			totalMonitored   = 0
			totalUnmonitored = 0
			qualities        = map[string]int{}
		)
		series := Series{}
		c.apiRequest(fmt.Sprintf("%s/api/v3/%s", c.hostname, "series"), &series)
		for _, s := range series {
			if s.Monitored {
				totalMonitored++
			} else {
				totalUnmonitored++
			}
			totalSeasons += s.Statistics.SeasonCount
			totalEpisodes += s.Statistics.EpisodeFileCount
			sizeOnDisk += s.Statistics.SizeOnDisk

			// Get Episode Qualities
			episodeFile := EpisodeFile{}
			c.apiRequest(fmt.Sprintf("%s/api/v3/%s?seriesId=%d", c.hostname, "episodefile", s.Id), &episodeFile)
			for _, e := range episodeFile {
				if e.Quality.Quality.Name != "" {
					qualities[e.Quality.Quality.Name]++
				}
			}
		}
		metrics.Series.WithLabelValues(c.hostname).Set(float64(len(series)))
		metrics.Seasons.WithLabelValues(c.hostname).Set(float64(totalSeasons))
		metrics.Episodes.WithLabelValues(c.hostname).Set(float64(totalEpisodes))
		metrics.FileSize.WithLabelValues(c.hostname).Set(float64(sizeOnDisk))
		metrics.Monitored.WithLabelValues(c.hostname).Set(float64(totalMonitored))
		metrics.Unmonitored.WithLabelValues(c.hostname).Set(float64(totalUnmonitored))

		for qualityName, count := range qualities {
			metrics.Qualities.WithLabelValues(c.hostname, qualityName).Set(float64(count))
		}

		// History
		history := History{}
		c.apiRequest(fmt.Sprintf("%s/api/v3/%s", c.hostname, "history"), &history)
		metrics.History.WithLabelValues(c.hostname).Set(float64(history.TotalRecords))

		// Wanted
		// sortKey required param for some reason :?
		wanted := WantedMissing{}
		c.apiRequest(fmt.Sprintf("%s/api/v3/%s?sortKey=%s", c.hostname, "wanted/missing", "airDateUtc"), &wanted)
		metrics.Wanted.WithLabelValues(c.hostname).Set(float64(wanted.TotalRecords))

		// Queue
		queue := Queue{}
		c.apiRequest(fmt.Sprintf("%s/api/v3/%s", c.hostname, "queue"), &queue)
		// Calculate total pages
		var totalPages = (queue.TotalRecords + queue.PageSize - 1) / queue.PageSize
		// Paginate
		var queueStatusAll = make([]QueueRecords, 0, queue.TotalRecords)
		queueStatusAll = append(queueStatusAll, queue.Records...)
		if totalPages > 1 {
			for page := 2; page <= totalPages; page++ {
				c.apiRequest(fmt.Sprintf("%s/api/v3/%s?page=%d", c.hostname, "queue", page), &queue)
				queueStatusAll = append(queueStatusAll, queue.Records...)
			}
		}
		// Group Status, TrackedDownloadStatus and TrackedDownloadState
		for i, s := range queueStatusAll {
			metrics.Queue.WithLabelValues(c.hostname, s.Status, s.TrackedDownloadStatus, s.TrackedDownloadState).Set(float64(i + 1))
		}

		// Root Folder
		rootFolders := RootFolder{}
		c.apiRequest(fmt.Sprintf("%s/api/v3/%s", c.hostname, "rootfolder"), &rootFolders)
		for _, rootFolder := range rootFolders {
			metrics.RootFolder.WithLabelValues(c.hostname, rootFolder.Path).Set(float64(rootFolder.FreeSpace))
		}

		// Health Issues
		systemHealth := SystemHealth{}
		c.apiRequest(fmt.Sprintf("%s/api/v3/%s", c.hostname, "health"), &systemHealth)
		for _, h := range systemHealth {
			metrics.SystemHealth.WithLabelValues(c.hostname, h.Source, h.Type, h.Message, h.WikiURL).Set(float64(1))
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
