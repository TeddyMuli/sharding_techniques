package transport

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type ShardClient struct {
	connections map[string]*redis.Client
}

func NewShardClient(registry map[string]string) (*ShardClient, error) {
	clients := make(map[string]*redis.Client)
	
	for nodeId, addr := range registry {
		rdb := redis.NewClient(&redis.Options{
			Addr: addr,
			Password: "",
			DB: 0,
		})
		
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		
		if err := rdb.Ping(ctx).Err(); err != nil {
			return nil, fmt.Errorf("failed to connect to %s at %s: %v", nodeId, addr, err)
		}
		
		clients[nodeId] = rdb
	}
	
	return &ShardClient{connections: clients}, nil
}

func (s *ShardClient) Write(ctx context.Context, shardId string, key string, value string) (time.Duration, error) {
	client, exists := s.connections[shardId]
	if !exists {
		return 0, fmt.Errorf("shad %s not found in pool", shardId)
	}
	
	start := time.Now()
	
	err := client.Set(ctx, key, value, 0).Err()
	
	latency := time.Since(start)
	return latency, err
}

func (s *ShardClient) Cleanup() {
	for _, client := range s.connections {
		client.Close()
	}
}
