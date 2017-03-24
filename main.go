package main

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"https://github.com/rancher/go-rancher/client"
	"github.com/rancher/go-rancher/client"
)

const (
	appName = "chaosrancher"
	image   = "quay.io/bscott/chaosrancher"
	version = "v0.1.0"
)

func init() {
	kingpin.Flag("debug", "Enable debug logging.").BoolVar(&debug)
}

func main() {

	// load Kingpin
	kingpin.Version(version)
	kingpin.Parse()

	// Check if Debug logging should be enabled.
	if debug {
		log.SetLevel(log.DebugLevel)
	}

	// create a new rancher client
	client, err := newClient()
	if err != nil {
		log.Fatal(err)
	}

}

// function newClient for rancher

func newClient() (*client.RancherClient, error) {

}
