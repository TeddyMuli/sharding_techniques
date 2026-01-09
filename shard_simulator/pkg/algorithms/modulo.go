package algorithms

import "hash/crc32"

type ModuloSharding struct {
	nodes []string
}

func NewModulo() *ModuloSharding {
	return &ModuloSharding{nodes: []string{}}
}

func (m *ModuloSharding) Name() string {
	return "Modulo Hashing"
}

func (m *ModuloSharding) AddNode(nodeId string) {
	m.nodes = append(m.nodes, nodeId)
}

func (m *ModuloSharding) GetShard(key string) string {
	if len(m.nodes) == 0 {
		return ""
	}

	hash := crc32.ChecksumIEEE([]byte(key))
	index := int(hash) % len(m.nodes)

	return m.nodes[index]
}
