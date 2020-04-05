package sonarr

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/onedr0p/sonarr-exporter/pkg/metrics"
)

var (
	apiUrlPattern = "%s/api/%s"
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
		if err := c.apiRequest(fmt.Sprintf(apiUrlPattern, c.hostname, "system/status"), &status); err != nil {
			metrics.Status.WithLabelValues(c.hostname).Set(0.0)
			return
		} else if (SystemStatus{}) == status {
			metrics.Status.WithLabelValues(c.hostname).Set(0.0)
			return
		} else {
			metrics.Status.WithLabelValues(c.hostname).Set(1.0)
		}

		// Series, Seasons, and Episodes
		var (
			totalSeasons      = 0
			totalEpisodes     = 0
			seriesMonitored   = 0
			seriesUnmonitored = 0
		)
		series := Series{}
		c.apiRequest(fmt.Sprintf(apiUrlPattern, c.hostname, "series"), &series)
		for _, s := range series {
			if s.Monitored {
				seriesMonitored++
			} else {
				seriesUnmonitored++
			}
			totalSeasons += s.SeasonCount
			totalEpisodes += s.EpisodeCount
		}
		metrics.Series.WithLabelValues(c.hostname).Set(float64(len(series)))
		metrics.Seasons.WithLabelValues(c.hostname).Set(float64(totalSeasons))
		metrics.Episodes.WithLabelValues(c.hostname).Set(float64(totalEpisodes))
		metrics.SeriesMonitored.WithLabelValues(c.hostname).Set(float64(seriesMonitored))
		metrics.SeriesUnmonitored.WithLabelValues(c.hostname).Set(float64(seriesUnmonitored))

		// History
		history := History{}
		c.apiRequest(fmt.Sprintf(apiUrlPattern, c.hostname, "history"), &history)
		metrics.History.WithLabelValues(c.hostname).Set(float64(history.TotalRecords))

		// Wanted
		wanted := WantedMissing{}
		c.apiRequest(fmt.Sprintf(apiUrlPattern, c.hostname, "wanted/missing"), &wanted)
		metrics.Wanted.WithLabelValues(c.hostname).Set(float64(wanted.TotalRecords))

		// Queue
		queue := Queue{}
		c.apiRequest(fmt.Sprintf(apiUrlPattern, c.hostname, "queue"), &queue)
		metrics.Queue.WithLabelValues(c.hostname).Set(float64(len(queue)))

		// Root Folder
		rootFolders := RootFolder{}
		c.apiRequest(fmt.Sprintf(apiUrlPattern, c.hostname, "rootfolder"), &rootFolders)
		for _, rootFolder := range rootFolders {
			metrics.RootFolder.WithLabelValues(c.hostname, rootFolder.Path).Set(float64(rootFolder.FreeSpace))
		}

		// Health Issues
		healthIssuesByType := map[string]int{}
		health := Health{}
		c.apiRequest(fmt.Sprintf(apiUrlPattern, c.hostname, "health"), &health)
		for _, h := range health {
			healthIssuesByType[h.Type]++
		}
		for issueType, count := range healthIssuesByType {
			metrics.Health.WithLabelValues(c.hostname, issueType).Set(float64(count))
		}
	}
}

func (c *Client) apiRequest(endpoint string, target interface{}) error {
	log.Printf("Sending HTTP request to %s", endpoint)

	req, err := http.NewRequest("GET", endpoint, nil)
	// req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("")))
	req.Header.Set("X-Api-Key", c.apiKey)
	if err != nil {
		log.Fatal("An error has occured when creating HTTP statistics request", err)
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatal("An error has occured during retrieving Sonarr statistics", err)
		return err
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}
