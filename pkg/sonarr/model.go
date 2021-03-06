package sonarr

// RootFolder - Stores struct of JSON response
type RootFolder []struct {
	Path      string `json:"path"`
	FreeSpace int64  `json:"freeSpace"`
}

// SystemStatus - Stores struct of JSON response
type SystemStatus struct {
	Version string `json:"version"`
	AppData string `json:"appData"`
	Branch  string `json:"branch"`
}

// Queue - Stores struct of JSON response
type Queue struct {
	Page          int            `json:"page"`
	PageSize      int            `json:"pageSize"`
	SortKey       string         `json:"sortKey"`
	SortDirection string         `json:"sortDirection"`
	TotalRecords  int            `json:"totalRecords"`
	Records       []QueueRecords `json:"records"`
}

// QueueRecords - Stores struct of JSON response
type QueueRecords struct {
	Size                  float64 `json:"size"`
	Title                 string  `json:"title"`
	Status                string  `json:"status"`
	TrackedDownloadStatus string  `json:"trackedDownloadStatus"`
	TrackedDownloadState  string  `json:"trackedDownloadState"`
	StatusMessages        []struct {
		Title    string   `json:"title"`
		Messages []string `json:"messages"`
	} `json:"statusMessages"`
	ErrorMessage string `json:"errorMessage"`
}

// History - Stores struct of JSON response
type History struct {
	TotalRecords int `json:"totalRecords"`
}

// WantedMissing - Stores struct of JSON response
type WantedMissing struct {
	TotalRecords int `json:"totalRecords"`
}

// SystemHealth - Stores struct of JSON response
type SystemHealth []struct {
	Source  string `json:"source"`
	Type    string `json:"type"`
	Message string `json:"message"`
	WikiURL string `json:"wikiUrl"`
}

// Series - Stores struct of JSON response
type Series []struct {
	Id         int  `json:"id"`
	Monitored  bool `json:"monitored"`
	Statistics struct {
		SeasonCount       int   `json:"seasonCount"`
		EpisodeFileCount  int   `json:"episodeFileCount"`
		EpisodeCount      int   `json:"episodeCount"`
		TotalEpisodeCount int   `json:"totalEpisodeCount"`
		SizeOnDisk        int64 `json:"sizeOnDisk"`
	} `json:"statistics"`
}

// EpisodeFile - Stores struct of JSON response
type EpisodeFile []struct {
	Size    int64 `json:"size"`
	Quality struct {
		Quality struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			Source     string `json:"source"`
			Resolution int    `json:"resolution"`
		} `json:"quality"`
	} `json:"quality"`
}
