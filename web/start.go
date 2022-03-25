package web

import (
	"cloudprobe/internal/g"
	"cloudprobe/prom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	metricsPath = "/metrics"
	ns      = "cloud_probe"
)

func Start() {
	metrics := prom.NewMetrics(ns)
	registry := prometheus.NewRegistry()
	registry.MustRegister(metrics)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET(metricsPath, gin.WrapH(promhttp.HandlerFor(registry, promhttp.HandlerOpts{})))
	router.Run(*g.ListenAndPort)
}
