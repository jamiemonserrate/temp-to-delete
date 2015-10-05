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

	flag.StringVar(&configPath, "config", "", "Path to redis config")
	flag.Parse()

	if configPath == "" {
		log.Fatal("Missing flag --config")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("Redis config does not exist")
	}

	conf, err := redisconf.Load(configPath)
	if err != nil {
		log.Fatal("Invalid redis config")
	}

	_, err = client.Connect(client.Host(conf.Host()), client.Port(conf.Port()))
	if err != nil {
		log.Fatal("Can not connect to Redis")
	}

	m := metrics.Metrics{
		metrics.Metric{
			Name:  "foo",
			Value: 1,
			Unit:  "bar",
		},
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(m)
}
