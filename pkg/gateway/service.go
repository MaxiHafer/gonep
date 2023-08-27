package gateway

import (
	"context"
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

func NewService(client *resty.Client) Service {
	return &service{
		client: client,
	}
}

type Service interface {
	Today(ctx context.Context, id string) ([]*Metric, error)
}

type service struct {
	client *resty.Client
}

func (s *service) Today(ctx context.Context, id string) ([]*Metric, error) {
	resp, err := s.client.R().
		SetContext(ctx).
		SetPathParam("id", id).
		Post("/detail/{id}")

	if err != nil {
		return nil, err
	}

	var metrics []*Metric
	if err := json.Unmarshal(resp.Body(), &metrics); err != nil {
		return nil, err
	}

	return metrics, nil
}
