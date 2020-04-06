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
type Queue []struct {
	Title string `json:"title"`
	Size  int32  `json:"size"`
}

// History - Stores struct of JSON response
type History struct {
	TotalRecords int `json:"totalRecords"`
}

// WantedMissing - Stores struct of JSON response
type WantedMissing struct {
	TotalRecords int `json:"totalRecords"`
}

// Health - Stores struct of JSON response
type Health []struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	WikiURL string `json:"wikiUrl"`
}

// Series - Stores struct of JSON response
type Series []struct {
	Id           int  `json:"id"`
	Monitored    bool `json:"monitored"`
	SeasonCount  int  `json:"seasonCount"`
	EpisodeCount int  `json:"episodeCount"`
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
