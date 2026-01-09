package algorithms

import (
	"strings"
)

type GeoSharding struct {
	regionMap map[string]string 
	nodes     []string
}

func NewGeo() *GeoSharding {
	return &GeoSharding{
		regionMap: make(map[string]string),
		nodes:     []string{},
	}
}

func (g *GeoSharding) Name() string {
	return "Geo-Sharding"
}

func (g *GeoSharding) AddNode(nodeID string) {
	g.nodes = append(g.nodes, nodeID)
	
	count := len(g.nodes)
	if count <= 2 {
		g.regionMap["US"] = nodeID
	} else if count <= 4 {
		g.regionMap["EU"] = nodeID
	} else {
		g.regionMap["ASIA"] = nodeID
	}
}

func (g *GeoSharding) GetShard(key string) string {
	lastChar := key[len(key)-1:]
	
	region := "ASIA"
	if strings.ContainsAny(lastChar, "0123") {
		region = "US"
	} else if strings.ContainsAny(lastChar, "456") {
		region = "EU"
	}

	if node, ok := g.regionMap[region]; ok {
		return node
	}
	if len(g.nodes) > 0 { return g.nodes[0] }
	return ""
}
