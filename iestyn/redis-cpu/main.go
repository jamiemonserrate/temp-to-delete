package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/pivotal-cf/cf-redis-broker/redis/client"
	"github.com/pivotal-cf/cf-redis-broker/redisconf"
	"github.com/pivotal-cf/redis-cpu/metrics"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "", "Path to redis config file")
	flag.Parse()

	if configPath == "" {
		log.Fatal("Missing flag --config")
	}

	config, err := redisconf.Load(configPath)
	if err != nil {
		log.Fatalf("Error loading Redis config file: %s", err)
	}

	port := config.Port()
	host := config.Host()
	password := config.Password()

	clientConnection, err := client.Connect(
		client.Port(port),
		client.Host(host),
		client.Password(password),
	)
	if err != nil {
		log.Fatalf("Error connecting to Redis: %s", err)
	}

	m, err := metrics.CPULoad(clientConnection)
	if err != nil {
		log.Fatalf("Error getting metrics: %s", err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(m); err != nil {
		log.Fatalf("Error encoding json: %s", err)
	}
}
