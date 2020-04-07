package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/onedr0p/sonarr-exporter/pkg/config"
	"github.com/onedr0p/sonarr-exporter/pkg/metrics"
	"github.com/onedr0p/sonarr-exporter/pkg/server"
	"github.com/onedr0p/sonarr-exporter/pkg/sonarr"
)

const (
	name = "sonarr-exporter"
)

var (
	s *server.Server
)

func main() {
	conf := config.Load()

	if conf.StartupDelay.Seconds() > 0.0 {
		fmt.Printf(fmt.Sprintf("Startup delay configured... sleeping for %v seconds", conf.StartupDelay.Seconds()))
		time.Sleep(conf.StartupDelay)
	}

	metrics.Init()

	initSonarrClient(conf.Hostname, conf.ApiKey, conf.Interval)
	initHttpServer(conf.Port)

	handleExitSignal()
}

func initSonarrClient(hostname, apiKey string, interval time.Duration) {
	client := sonarr.NewClient(hostname, apiKey, interval)
	go client.Scrape()
}

func initHttpServer(port string) {
	s = server.NewServer(port)
	go s.ListenAndServe()
}

func handleExitSignal() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	s.Stop()
	fmt.Println(fmt.Sprintf("\n%s HTTP server stopped", name))
}
