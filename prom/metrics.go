package prom

import (
	"cloudprobe/internal/g"
	"cloudprobe/internal/log"
	"cloudprobe/internal/tasks"
	"cloudprobe/utils"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"sync"
	"time"
)

var (
	ms = make(map[string]*prometheus.Desc)
)

type Metrics struct {
	Ms map[string] *prometheus.Desc
	mutex sync.Mutex
}

func NewCommonMetric(ns, metricsName, help string, labels []string) *prometheus.Desc {
	return prometheus.NewDesc(ns+"_"+metricsName, help, labels, nil)
}

func NewMetrics(ns string) *Metrics {
	// icmp
	ms["icmp"] = NewCommonMetric(ns, "icmp_probe_with_args", "For icmp probe. " +
		"if return '-1001' represent error and ErrorMes is label of status!",
		[]string{"name", "status"})

	// nc
	ms["nc"] = NewCommonMetric(ns, "nc_probe_with_args", "for nc probe. " +
		"if return '-1001' represent error and ErrorMes is label of status",
		[]string{"name", "status"})

	// pos
	ms["pos"] = NewCommonMetric(ns, "pos_probe_no_args", "for pos probe. " +
		"if return '-1001' represent error and ErrorMes is label of status",
		[]string{"name", "status"})

	return &Metrics{
		Ms: ms,
	}
}

func (c *Metrics) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range c.Ms {
		ch <- m
	}
}

func (c *Metrics) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	//job
	tasks.NewServer(g.IcmpScripts,g.IcmpArgs,g.NcScripts,g.NcArgs,g.PosScripts).Handler()

	//icmp range
	icmpRes := func(k,v interface{}) bool {
		ok := utils.RangeSyncMap(c, ch, "icmp", k,v)
		return ok
	}
	g.IcmpResMap.Range(icmpRes)

	// nc range
	ncRes := func(k,v interface{}) bool {
		ok := utils.RangeSyncMap(c, ch, "nc", k,v)
		return ok
	}
	g.NcResMap.Range(ncRes)

	// pos range
	posRes := func(k, v interface{}) bool {
		ok := utils.RangeSyncMap(c, ch, "pos", k,v)
		return ok
	}
	g.PosResMap.Range(posRes)

	// END
	log.Info(fmt.Sprintf("Finished at %v", time.Now().Unix()))
}
