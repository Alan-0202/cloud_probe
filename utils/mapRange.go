package utils

import (
	"cloudprobe/prom"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)

func RangeSyncMap(c *prom.Metrics, ch chan<- prometheus.Metric, metricName string, k,v interface{}) bool {
	name := CleanStr(k.(string))
	valFloat64, err := strconv.ParseFloat(v.(map[string]string)["value"], 64)
	if err != nil {
		ch <- prometheus.MustNewConstMetric(c.Ms[metricName], prometheus.GaugeValue, valFloat64,
			name, v.(map[string]string)["status"])
		return true
	}
	ch <- prometheus.MustNewConstMetric(c.Ms[metricName], prometheus.GaugeValue, valFloat64,
		name, v.(map[string]string)["status"])
	return true
}
