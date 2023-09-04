package gateway

import (
	"context"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"time"
)

func NewService(client *resty.Client) Service {
	return &service{
		client: client,
	}
}

type Metric interface {
	Time() *time.Time
	KilowattHours() float64
}

type DetailMetric interface {
	Time() *time.Time
	Watts() int
}

type Service interface {
	Today(ctx context.Context, id string) ([]DetailMetric, error)
	Week(ctx context.Context, id string) ([]Metric, error)
	Month(ctx context.Context, id string) ([]Metric, error)
	Year(ctx context.Context, id string) ([]Metric, error)
}

type service struct {
	client *resty.Client
}

func (s *service) Today(ctx context.Context, id string) ([]DetailMetric, error) {
	resp, err := s.client.R().
		SetContext(ctx).
		SetPathParam("id", id).
		Post("/pv_monitor/appservice/detail/{id}/0")

	if err != nil {
		return nil, err
	}

	var tsMetrics []*timestampMetric
	if err := json.Unmarshal(resp.Body(), &tsMetrics); err != nil {
		return nil, err
	}

	metrics := make([]DetailMetric, len(tsMetrics))
	for i := range tsMetrics {
		metrics[i] = tsMetrics[i]
	}

	return metrics, nil
}

func (s *service) Week(ctx context.Context, id string) ([]Metric, error) {
	resp, err := s.client.R().
		SetContext(ctx).
		SetPathParam("id", id).
		Post("/pv_monitor/appservice/week/{id}/0")

	if err != nil {
		return nil, err
	}

	var dayMetrics []*dayMetric
	if err := json.Unmarshal(resp.Body(), &dayMetrics); err != nil {
		return nil, err
	}

	metrics := make([]Metric, len(dayMetrics))
	for i := range dayMetrics {
		metrics[i] = dayMetrics[i]
	}

	return metrics, nil
}

func (s *service) Month(ctx context.Context, id string) ([]Metric, error) {
	resp, err := s.client.R().
		SetContext(ctx).
		SetPathParam("id", id).
		Post("/pv_monitor/appservice/month/{id}/0")

	if err != nil {
		return nil, err
	}

	var monthMetrics []*monthMetric
	if err := json.Unmarshal(resp.Body(), &monthMetrics); err != nil {
		return nil, err
	}

	metrics := make([]Metric, len(monthMetrics))
	for i := range monthMetrics {
		metrics[i] = monthMetrics[i]
	}

	return metrics, nil
}

func (s *service) Year(ctx context.Context, id string) ([]Metric, error) {
	resp, err := s.client.R().
		SetContext(ctx).
		SetPathParam("id", id).
		Post("/pv_monitor/appservice/year/{id}/0")

	if err != nil {
		return nil, err
	}

	var monthMetrics []*monthMetric
	if err := json.Unmarshal(resp.Body(), &monthMetrics); err != nil {
		return nil, err
	}

	metrics := make([]Metric, len(monthMetrics))
	for i := range monthMetrics {
		metrics[i] = monthMetrics[i]
	}

	return metrics, nil
}
