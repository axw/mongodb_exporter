package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	indexGaugesMissRatio = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: Namespace,
		Subsystem: "index_counters",
		Name:      "miss_ratio",
		Help:      "The missRatio value is the ratio of hits to misses. This value is typically 0 or approaching 0",
	})
)

var (
	indexGaugesTotal = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: Namespace,
		Name:      "index_counters_total",
		Help:      "Total indexes by type",
	}, []string{"type"})
)

//IndexCounterStats index counter stats
type IndexCounterStats struct {
	Accesses  float64 `bson:"accesses"`
	Hits      float64 `bson:"hits"`
	Misses    float64 `bson:"misses"`
	Resets    float64 `bson:"resets"`
	MissRatio float64 `bson:"missRatio"`
}

// Export exports the data to prometheus.
func (indexGaugesStats *IndexCounterStats) Export(ch chan<- prometheus.Metric) {
	indexGaugesTotal.WithLabelValues("accesses").Set(indexGaugesStats.Accesses)
	indexGaugesTotal.WithLabelValues("hits").Set(indexGaugesStats.Hits)
	indexGaugesTotal.WithLabelValues("misses").Set(indexGaugesStats.Misses)
	indexGaugesTotal.WithLabelValues("resets").Set(indexGaugesStats.Resets)

	indexGaugesMissRatio.Set(indexGaugesStats.MissRatio)

	indexGaugesTotal.Collect(ch)
	indexGaugesMissRatio.Collect(ch)

}

// Describe describes the metrics for prometheus
func (indexGaugesStats *IndexCounterStats) Describe(ch chan<- *prometheus.Desc) {
	indexGaugesTotal.Describe(ch)
	indexGaugesMissRatio.Describe(ch)
}
