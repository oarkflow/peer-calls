package server

import (
  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promauto"
)

var prometheusHomeViewsTotal = promauto.NewCounter(prometheus.CounterOpts{
  Name: "pageviews_index_total",
  Help: "Total number of homepage views",
})

var prometheusCallJoinTotal = promauto.NewCounter(prometheus.CounterOpts{
  Name: "pageviews_call_join_total",
  Help: "Total number of new call requests",
})

var prometheusCallViewsTotal = promauto.NewCounter(prometheus.CounterOpts{
  Name: "pageviews_call_total",
  Help: "Total number of homepage views",
})

var prometheusWSConnTotal = promauto.NewCounter(prometheus.CounterOpts{
  Name: "ws_conn_total",
  Help: "Total number of opened websocket connections",
})

var prometheusWSConnActive = promauto.NewGauge(prometheus.GaugeOpts{
  Name: "ws_conn_active",
  Help: "Total number of active websocket connections",
})

var prometheusWSConnErrTotal = promauto.NewCounter(prometheus.CounterOpts{
  Name: "ws_conn_err_total",
  Help: "Total number of errored out websocket connections",
})

var prometheusWSConnDuration = promauto.NewHistogram(prometheus.HistogramOpts{
  Name:    "ws_conn_duration",
  Help:    "Duration of websocket connections",
  Buckets: []float64{1, 60, 5 * 60, 15 * 60, 30 * 60, 45 * 60, 60 * 60, 120 * 60},
})

var prometheusWebRTCConnTotal = promauto.NewCounter(prometheus.CounterOpts{
  Name: "webrtc_conn_total",
  Help: "Total number of opened webrtc connections",
})

var prometheusWebRTCConnActive = promauto.NewGauge(prometheus.GaugeOpts{
  Name: "webrtc_conn_active",
  Help: "Total number of active webrtc connections",
})

// var prometheusWebRTCConnErrTotal = promauto.NewCounter(prometheus.CounterOpts{
// 	Name: "webrtc_conn_err_total",
// 	Help: "Total number of errored out webrtc connections",
// })

var prometheusWebRTCConnDuration = promauto.NewHistogram(prometheus.HistogramOpts{
  Name:    "webrtc_conn_duration_seconds",
  Help:    "Duration of webrtc connections",
  Buckets: []float64{1, 60, 5 * 60, 15 * 60, 30 * 60, 45 * 60, 60 * 60, 120 * 60},
})
