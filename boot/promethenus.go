package boot
// promethenus 指标采集
import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Standard default metrics
var reqCnt = &Metric{
	ID:          "reqCnt",
	Name:        "requests_counter",
	Description: "How many HTTP requests processed, partitioned by status code and HTTP method.",
	Type:        "counter_vec",
	Args:        []string{"code", "method", "handler", "host", "url"},
}

var reqDur = &Metric{
	ID:          "reqDur",
	Name:        "request_latencies",
	Description: "The HTTP request latencies in Millisecond.",
	Type:        "histogram_vec",
	Args:        []string{"code", "method", "host", "url"},
}

var resSz = &Metric{
	ID:          "resSz",
	Name:        "response_bytes_size",
	Description: "The HTTP response sizes in bytes.",
	Type:        "summary"}

var reqSz = &Metric{
	ID:          "reqSz",
	Name:        "request_bytes_size",
	Description: "The HTTP request sizes in bytes.",
	Type:        "summary"}

var standardMetrics = []*Metric{
	reqCnt,
	reqDur,
	resSz,
	reqSz,
}

// Metric is a definition for the name, description, type, ID, and
// prometheus.Collector type (i.e. CounterVec, Summary, etc) of each metric
type Metric struct {
	MetricCollector prometheus.Collector
	ID              string
	Name            string
	Description     string
	Type            string
	Args            []string
}

// Prometheus contains the metrics gathered by the instance and its path
type Prometheus struct {
	reqCnt *prometheus.CounterVec
	reqDur *prometheus.HistogramVec
	reqSz, resSz  prometheus.Summary

	MetricsList []*Metric
}

// NewPrometheus generates a new set of metrics with a certain subsystem name
func NewPrometheus(subsystem string) *Prometheus {
	var metricsList []*Metric
	for _, metric := range standardMetrics {
		metricsList = append(metricsList, metric)
	}

	p := &Prometheus{
		MetricsList: metricsList,
	}

	p.registerMetrics(subsystem)

	return p
}

// NewMetric associates prometheus.Collector based on Metric.Type
func NewMetric(m *Metric, subsystem string) prometheus.Collector {
	var metric prometheus.Collector
	switch m.Type {
	case "counter_vec":
		metric = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
			m.Args,
		)
	case "counter":
		metric = prometheus.NewCounter(
			prometheus.CounterOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
		)
	case "gauge_vec":
		metric = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
			m.Args,
		)
	case "gauge":
		metric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
		)
	case "histogram_vec":
		metric = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
			m.Args,
		)
	case "histogram":
		metric = prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
		)
	case "summary_vec":
		metric = prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
			m.Args,
		)
	case "summary":
		metric = prometheus.NewSummary(
			prometheus.SummaryOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
		)
	}
	return metric
}

func (p *Prometheus) registerMetrics(subsystem string) {
	for _, metricDef := range p.MetricsList {
		metric := NewMetric(metricDef, subsystem)
		if err := prometheus.Register(metric); err != nil {
			log.Fatal(fmt.Errorf("%s could not be registered in Prometheus)", metricDef.Name))
		}
		switch metricDef {
		case reqCnt:
			p.reqCnt = metric.(*prometheus.CounterVec)
		case reqDur:
			p.reqDur = metric.(*prometheus.HistogramVec)
		case resSz:
			p.resSz = metric.(prometheus.Summary)
		case reqSz:
			p.reqSz = metric.(prometheus.Summary)
		}
		metricDef.MetricCollector = metric
	}
}

// Use adds the middleware to a gin engine.
func (p *Prometheus) Use(e *gin.Engine) {
	e.Use(p.HandlerFunc())
}

// HandlerFunc defines handler function for middleware
func (p *Prometheus) HandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.String() == "/metrics" || c.Request.URL.String() == "/sys/metrics" {
			c.Next()
			return
		}

		start := time.Now()
		reqSz := computeApproximateRequestSize(c.Request)

		c.Next()

		statusCode := strconv.Itoa(c.Writer.Status())
		latency := float64(time.Since(start)) / float64(time.Millisecond)
		resSz := float64(c.Writer.Size())

		p.reqDur.WithLabelValues(statusCode, c.Request.Method, c.Request.Host, c.Request.URL.Path).Observe(latency)
		p.reqCnt.WithLabelValues(statusCode, c.Request.Method, c.HandlerName(), c.Request.Host, c.Request.URL.Path).Inc()
		p.reqSz.Observe(float64(reqSz))
		p.resSz.Observe(resSz)
	}
}

func HandlerPrometheus() func(w http.ResponseWriter, req *http.Request) {
	h := promhttp.Handler()
	return func(w http.ResponseWriter, req *http.Request) {
		h.ServeHTTP(w, req)
	}
}

// From https://github.com/DanielHeckrath/gin-prometheus/blob/master/gin_prometheus.go
func computeApproximateRequestSize(r *http.Request) int {
	s := 0
	if r.URL != nil {
		s = len(r.URL.Path)
	}

	s += len(r.Method)
	s += len(r.Proto)
	for name, values := range r.Header {
		s += len(name)
		for _, value := range values {
			s += len(value)
		}
	}
	s += len(r.Host)

	// N.B. r.Form and r.MultipartForm are assumed to be included in r.URL.

	if r.ContentLength != -1 {
		s += int(r.ContentLength)
	}
	return s
}
