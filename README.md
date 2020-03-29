# sonarr-exporter

Prometheus Exporter for Sonarr

[![Docker Pulls](https://img.shields.io/docker/pulls/onedr0p/sonarr-exporter)](https://hub.docker.com/r/onedr0p/sonarr-exporter)

## Usage

|Name             |Description                                                  |Default|
|-----------------|-------------------------------------------------------------|-------|
|`SONARR_HOSTNAME`|You Sonarr instance's URL                                    |       |
|`SONARR_APIKEY`  |Your Sonarr instance's API Key                               |       |
|`INTERVAL`       |The duration of which the exporter will scrape the Sonarr API|`2m`   |
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
# HELP sonarr_episodes_total Total number of episodes downloaded for all series
# TYPE sonarr_episodes_total gauge
sonarr_episodes_total{hostname="http://localhost:8989"} 49822
# HELP sonarr_health_issues Amount of health issues in Sonarr
# TYPE sonarr_health_issues gauge
sonarr_health_issues{hostname="http://localhost:8989",type="error"} 1
# HELP sonarr_history_total Total number of records in history
# TYPE sonarr_history_total gauge
sonarr_history_total{hostname="http://localhost:8989"} 107134
# HELP sonarr_missing_episodes_total Total number of missing episodes
# TYPE sonarr_missing_episodes_total gauge
sonarr_missing_episodes_total{hostname="http://localhost:8989"} 36
# HELP sonarr_queue_total Total number of episodes in queue
# TYPE sonarr_queue_total gauge
sonarr_queue_total{hostname="http://localhost:8989"} 1
# HELP sonarr_root_folder_space Root folder space
# TYPE sonarr_root_folder_space gauge
sonarr_root_folder_space{folder="/media/Library/Television/",hostname="http://localhost:8989"} 2.5840012099584e+13
# HELP sonarr_seasons_total Total number of seasons for all series
# TYPE sonarr_seasons_total gauge
sonarr_seasons_total{hostname="http://localhost:8989"} 3426
# HELP sonarr_series_total Total number of series
# TYPE sonarr_series_total gauge
sonarr_series_total{hostname="http://localhost:8989"} 881
# HELP sonarr_status System Status
# TYPE sonarr_status gauge
sonarr_status{hostname="http://localhost:8989"} 1
```
