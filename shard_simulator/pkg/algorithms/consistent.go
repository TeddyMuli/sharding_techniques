package algorithms

import (
	"github.com/buraksezer/consistent"
	"github.com/cespare/xxhash"
)

type myMember string
type ConsistentSharding struct {
    ring *consistent.Consistent
}

func (m myMember) String() string {
	return string(m)
}

type hasher struct{}

func (h hasher) Sum64(data []byte) uint64 {
	return xxhash.Sum64(data)
}

func NewConsistent() *ConsistentSharding {
	cfg := consistent.Config{
		PartitionCount:    271,
		ReplicationFactor: 20,
		Load:        1.25,
		Hasher:            hasher{},
	}
	
	return &ConsistentSharding{
		ring: consistent.New(nil, cfg),
	}
}

func (c *ConsistentSharding) Name() string {
    return "Consistent Hashing"
}

func (c *ConsistentSharding) AddNode(nodeID string) {
    c.ring.Add(myMember(nodeID))
}

func (c *ConsistentSharding) GetShard(key string) string {
    owner := c.ring.LocateKey([]byte(key))
    return owner.String()
}
