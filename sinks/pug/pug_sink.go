package pug

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/stripe/veneur/samplers"
	"github.com/stripe/veneur/sinks"
	"github.com/stripe/veneur/ssf"
	"github.com/stripe/veneur/trace"
	"google.golang.org/grpc"
)

// Sink represents a gRPC metrics sink
type Sink struct {
	target      string
	client      MetricSinkClient
	grpcConn    *grpc.ClientConn
	traceClient *trace.Client
	log         *logrus.Logger
}

var _ sinks.MetricSink = &Sink{}

// NewPugSink creates a new gRPC sink for metrics
func NewPugSink(ctx context.Context, target string, log *logrus.Logger, opts ...grpc.DialOption) (sinks.MetricSink, error) {
	conn, err := grpc.DialContext(ctx, target, opts...)
	if err != nil {
		log.WithError(err).WithFields(logrus.Fields{
			"target": target,
		}).Error("Error establishing connection to gRPC server")
		return nil, err
	}

	return &Sink{
		client:   NewMetricSinkClient(conn),
		grpcConn: conn,
		target:   target,
		log:      log,
	}, nil
}

// Name returns this sink's name. As the gRPC sink is generic, it's expected
// that this is set via configuration and injected.
func (s *Sink) Name() string {
	return "pug"
}

// Start performs final preparations on the sink before it is
// ready to begin ingesting spans.
func (s *Sink) Start(cl *trace.Client) error {
	s.traceClient = cl
	return nil
}

// Flush flushes the metrics
func (s *Sink) Flush(ctx context.Context, metrics []samplers.InterMetric) error {
	if len(metrics) == 0 {
		return nil
	}

	s.log.Debugf("Flushing %d metrics:", len(metrics))
	payload := make([]*Metric, 0, len(metrics))
	for _, m := range metrics {
		payload = append(payload, &Metric{
			Name:  m.Name,
			Type:  int32(m.Type),
			Time:  m.Timestamp,
			Value: m.Value,
			Tags:  m.Tags,
		})
	}

	if _, err := s.client.SendMetric(ctx, &Request{
		Metrics: payload,
	}); err != nil {
		s.log.WithFields(logrus.Fields{
			logrus.ErrorKey: err,
			"target":        s.target,
			"message":       err.Error(),
		}).Error("Error sending span to gRPC sink target")
		return err
	}

	return nil
}

// FlushOtherSamples ...
func (s *Sink) FlushOtherSamples(ctx context.Context, samples []ssf.SSFSample) {
	return
}
