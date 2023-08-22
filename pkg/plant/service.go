package plant

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
	Status(ctx context.Context, id string) (*Status, error)
	List(ctx context.Context) ([]*Plant, error)
}

type service struct {
	client *resty.Client
}

func (s *service) Status(ctx context.Context, id string) (*Status, error) {
	resp, err := s.client.R().
		SetContext(ctx).
		SetPathParam("id", id).
		Get("/status/{id}")

	if err != nil {
		return nil, err
	}

	status := new(Status)
	if err := json.Unmarshal(resp.Body(), status); err != nil {
		return nil, err
	}

	return status, nil
}

func (s *service) List(ctx context.Context) ([]*Plant, error) {
	resp, err := s.client.R().
		SetContext(ctx).
		Get("/pvlist")
	if err != nil {
		return nil, err
	}

	var plants []*Plant
	if err := json.Unmarshal(resp.Body(), &plants); err != nil {
		return nil, err
	}

	return plants, nil
}
