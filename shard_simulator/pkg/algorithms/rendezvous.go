package algorithms

import (
	"github.com/cespare/xxhash"
)

type RendezvousHashing struct {
	nodes []string
}

func NewRendezvous() *RendezvousHashing {
	return &RendezvousHashing{nodes: []string{}}
}

func (r *RendezvousHashing) Name() string {
	return "Rendezvous Hashing"
}

func (r *RendezvousHashing) AddNode(nodeId string) {
	r.nodes = append(r.nodes, nodeId)
}

func (r *RendezvousHashing) GetShard(key string) string {
	var maxScore uint64
	var champion string

	for _, node := range r.nodes {
		hashInput := []byte(key + node)
		score := xxhash.Sum64(hashInput)

		if score > maxScore {
			maxScore = score
			champion = node
		}
	}

	return champion
}
