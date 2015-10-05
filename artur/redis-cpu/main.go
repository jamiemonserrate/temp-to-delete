package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gosexy/redis"
)

func main() {
	var (
		host     string
		password string
		port     uint
	)

	flag.StringVar(&host, "host", "", "Redis host")
	flag.StringVar(&password, "password", "", "Redis password")
	flag.UintVar(&port, "port", 0, "Redis port")

	flag.Parse()

	if host == "" {
		log.Fatal("Host not defined")
	}

	if port == 0 {
		log.Fatal("Port not defined")
	}

	client := *redis.New()
	err := client.Connect(host, port)
	if err != nil {
		log.Fatalf("Cannot connect to redis: %s", err)
	}

	fmt.Println("{}")
}
