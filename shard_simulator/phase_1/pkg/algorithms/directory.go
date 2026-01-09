package algorithms

import "sync"

type DirectorySharding struct {
	nodes     []string
	directory map[string]string
	mu        sync.Mutex
	nextIdx   int
}

func NewDirectory() *DirectorySharding {
	return &DirectorySharding{
		nodes:     []string{},
		directory: make(map[string]string),
	}
}

func (d *DirectorySharding) Name() string {
	return "Directory (Lookup) Sharding"
}

func (d *DirectorySharding) AddNode(nodeID string) {
	d.nodes = append(d.nodes, nodeID)
}

func (d *DirectorySharding) GetShard(key string) string {
	d.mu.Lock()
	defer d.mu.Unlock()

	if shard, exists := d.directory[key]; exists {
		return shard
	}

	if len(d.nodes) == 0 { return "" }
	
	assignedNode := d.nodes[d.nextIdx]
	d.directory[key] = assignedNode
	
	d.nextIdx = (d.nextIdx + 1) % len(d.nodes)
	
	return assignedNode
}
