package metrics

import (
	"strconv"

	redis "github.com/pivotal-cf/cf-redis-broker/redis/client"
)

type Metric struct {
	Key   string
	Value float64
	Unit  string
}

func CPULoad(c redis.Client) (Metric, error) {
	infoMap, err := c.Info()
	if err != nil {
		return Metric{}, err
	}

	value, err := strconv.ParseFloat(infoMap["used_cpu_sys"], 64)
	if err != nil {
		return Metric{}, err
	}

	return Metric{"cpu_load", value, "load"}, nil
}
