# sonarr-exporter

Prometheus Exporter for Sonarr

[![Docker Pulls](https://img.shields.io/docker/pulls/onedr0p/sonarr-exporter)](https://hub.docker.com/r/onedr0p/sonarr-exporter) [![Go Report Card](https://goreportcard.com/badge/github.com/onedr0p/sonarr-exporter)](https://goreportcard.com/report/github.com/onedr0p/sonarr-exporter)

## Usage

|Name             |Description                                                  |Default|
|-----------------|-------------------------------------------------------------|-------|
|`SONARR_HOSTNAME`|You Sonarr instance's URL                                    |       |
|`SONARR_APIKEY`  |Your Sonarr instance's API Key                               |       |
|`INTERVAL`       |The duration of which the exporter will scrape the Sonarr API|`10m`  |
|`PORT`           |The port the exporter will listen on                         |`9811` |

### Docker Compose Example

```yaml
version: '3.7'
services:
  sonarr-exporter:
    image: onedr0p/sonarr-exporter:v1.0.0
    environment:
      SONARR_HOSTNAME: "http://localhost:7878"
      SONARR_APIKEY: "..."
      INTERVAL: "1h"
```

### Metrics

```bash
# HELP sonarr_episode_download_total Total number of episodes downloaded for all series
# TYPE sonarr_episode_download_total gauge
sonarr_episode_download_total{hostname="http://localhost:8989"} 49233
# HELP sonarr_episode_missing_total Total number of missing episodes
# TYPE sonarr_episode_missing_total gauge
sonarr_episode_missing_total{hostname="http://localhost:8989"} 28
# HELP sonarr_episode_bytes Total file size of all episodes in bytes
# TYPE sonarr_episode_bytes gauge
sonarr_episode_bytes{hostname="http://localhost:8989"} 4.4574724013551e+13
# HELP sonarr_episode_quality_total Total number of downloaded episodes by quality
# TYPE sonarr_episode_quality_total gauge
sonarr_episode_quality_total{hostname="http://localhost:8989",quality="Bluray-1080p"} 4876
sonarr_episode_quality_total{hostname="http://localhost:8989",quality="Bluray-1080p Remux"} 6
sonarr_episode_quality_total{hostname="http://localhost:8989",quality="Bluray-2160p Remux"} 1
sonarr_episode_quality_total{hostname="http://localhost:8989",quality="Bluray-480p"} 2059
sonarr_episode_quality_total{hostname="http://localhost:8989",quality="Bluray-720p"} 705
sonarr_episode_quality_total{hostname="http://localhost:8989",quality="DVD"} 13710
sonarr_episode_quality_total{hostname="http://localhost:8989",quality="HDTV-1080p"} 1171
sonarr_episode_quality_total{hostname="http://localhost:8989",quality="HDTV-720p"} 1683
sonarr_episode_quality_total{hostname="http://localhost:8989",quality="Raw-HD"} 2
sonarr_episode_quality_total{hostname="http://localhost:8989",quality="SDTV"} 7518
sonarr_episode_quality_total{hostname="http://localhost:8989",quality="WEBDL-1080p"} 8793
sonarr_episode_quality_total{hostname="http://localhost:8989",quality="WEBDL-480p"} 4302
sonarr_episode_quality_total{hostname="http://localhost:8989",quality="WEBDL-720p"} 1776
sonarr_episode_quality_total{hostname="http://localhost:8989",quality="WEBRip-1080p"} 297
sonarr_episode_quality_total{hostname="http://localhost:8989",quality="WEBRip-480p"} 535
sonarr_episode_quality_total{hostname="http://localhost:8989",quality="WEBRip-720p"} 364
# HELP sonarr_history_total Total number of records in history
# TYPE sonarr_history_total gauge
sonarr_history_total{hostname="http://localhost:8989"} 106681
# HELP sonarr_queue_total Total number of episodes in queue
# TYPE sonarr_queue_total gauge
sonarr_queue_total{hostname="http://localhost:8989"} 0
# HELP sonarr_rootfolder_freespace_bytes Root folder space in bytes
# TYPE sonarr_rootfolder_freespace_bytes gauge
sonarr_rootfolder_freespace_bytes{folder="/media/Library/Television/",hostname="http://localhost:8989"} 2.499789602816e+13
# HELP sonarr_season_total Total number of seasons for all series
# TYPE sonarr_season_total gauge
sonarr_season_total{hostname="http://localhost:8989"} 3472
# HELP sonarr_series_monitored_total Total number of monitored series
# TYPE sonarr_series_monitored_total gauge
sonarr_series_monitored_total{hostname="http://localhost:8989"} 308
# HELP sonarr_series_total Total number of series
# TYPE sonarr_series_total gauge
sonarr_series_total{hostname="http://localhost:8989"} 892
# HELP sonarr_series_unmonitored_total Total number of unmonitored series
# TYPE sonarr_series_unmonitored_total gauge
sonarr_series_unmonitored_total{hostname="http://localhost:8989"} 584
# HELP sonarr_status System Status
# TYPE sonarr_status gauge
sonarr_status{hostname="http://localhost:8989"} 1
# HELP sonarr_health_issues Health issues in Sonarr
# TYPE sonarr_health_issues gauge
sonarr_health_issues{hostname="http://localhost:8989",message="No download client is available",type="warning",wikiurl="https://github.com/Sonarr/Sonarr/wiki/Health-checks#no-download-client-is-available"} 1
sonarr_health_issues{hostname="http://localhost:8989",message="No indexers available with Automatic Search enabled, Sonarr will not provide any automatic search results",type="warning",wikiurl="https://github.com/Sonarr/Sonarr/wiki/Health-checks#no-indexers-available-with-automatic-search-enabled-sonarr-will-not-provide-any-automatic-search-results"} 1
sonarr_health_issues{hostname="http://localhost:8989",message="No indexers available with RSS sync enabled, Sonarr will not grab new releases automatically",type="error",wikiurl="https://github.com/Sonarr/Sonarr/wiki/Health-checks#no-indexers-available-with-rss-sync-enabled-sonarr-will-not-grab-new-releases-automatically"} 1
```
