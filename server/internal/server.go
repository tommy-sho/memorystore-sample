package server

import (
	"context"
	"fmt"

	"github.com/tommy-sho/memory-store-sample/server/redis"

	sample "github.com/tommy-sho/memory-store-sample/server/genproto"
)

type sampleServcer struct {
	client *redis.Client
}

func NewsampleServer(addr string) (sample.SampleServiceServer, error) {
	r, err := redis.NewClient(addr)
	if err != nil {
		return &sampleServcer{}, fmt.Errorf("new sample server error: %v", err)
	}

	return &sampleServcer{client: r}, nil
}

func (s *sampleServcer) Register(ctx context.Context, req *sample.RegisterRequest) (*sample.RegisterResponse, error) {
	err := s.client.Set(ctx, req.Key, req.Value)
	if err != nil {
		return nil, fmt.Errorf("sample server register error: %v", err)
	}

	return &sample.RegisterResponse{}, nil
}

func (s *sampleServcer) Get(ctx context.Context, req *sample.GetRequest) (*sample.GetResponse, error) {
	v, err := s.client.Get(ctx, req.Key)
	if err != nil {
		return nil, fmt.Errorf("sample server get error: %v", err)
	}

	return &sample.GetResponse{Key: req.Key, Value: v}, nil
}
