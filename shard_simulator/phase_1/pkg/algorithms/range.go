package algorithms

import (
	"strconv"
	"strings"
)

type RangeSharding struct {
	nodes [] string
	rangeSize int
}

func NewRange(rangeSize int) *RangeSharding {
	return &RangeSharding{
		nodes: []string{},
		rangeSize: rangeSize,
	}
}

func (r *RangeSharding) Name() string {
	return "Range-Based Sharding"
}

func (r *RangeSharding) AddNode(nodeID string) {
	r.nodes = append(r.nodes, nodeID)
}

func (r *RangeSharding) GetShard(key string) string {
	if len(r.nodes) == 0 { return "" }
	parts := strings.Split(key, "-")
	if len(parts) < 2 { return r.nodes[0] }
	
	id, _ := strconv.Atoi(parts[1])
	
	chunkIndex := id / r.rangeSize
	
	if chunkIndex > len(r.nodes) {
		return r.nodes[len(r.nodes)-1]
	}
	
	return r.nodes[chunkIndex]
}
