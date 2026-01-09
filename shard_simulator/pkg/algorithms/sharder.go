package algorithms

type Sharder interface {
	Name() string
	AddNode(nodeId string)
	GetShard(key string) string
}
