package utils

import 	(
	"sync"
	"fmt"
	"github.com/TeddyMuli/sharding_techniques/shard_simulator/phase_1/pkg/analyzer"
	"github.com/TeddyMuli/sharding_techniques/shard_simulator/phase_1/pkg/algorithms"
)

type DistRow struct {
    Algorithm string
    ShardID   string
    KeyCount  int
}

type MoveRow struct {
    Algorithm    string
    PercentMoved float64
}


func RunBenchmark(algo algorithms.Sharder, keys []string) {
	var distRows []DistRow
    var moveRows []MoveRow

	fmt.Printf("================ %s ================\n", algo.Name())

	for i := range 10 {
		algo.AddNode(fmt.Sprintf("node-%02d", i))
	}

	distribution := make(map[string]int)
	keyLocation := make(map[string]string)

	for _, key := range keys {
		node := algo.GetShard(key)
		distribution[node]++
		keyLocation[key] = node
	}
	
	for node, count := range distribution {
            distRows = append(distRows, DistRow{
                Algorithm: algo.Name(),
                ShardID:   node,
                KeyCount:  count,
            })
        }
	var mu sync.Mutex

	var wg sync.WaitGroup
	workerCount := 10

	chunkSize := len(keys) / workerCount

	for i := range workerCount {
		wg.Add(1)

		start := i * chunkSize
		end := start + chunkSize

		if i == workerCount-1 {
			end = len(keys)
		}

		go func(keyChunk []string) {
			defer wg.Done()
			localStats := make(map[string]int)

			for _, key := range keyChunk {
				node := algo.GetShard(key)
				localStats[node]++
			}

			mu.Lock()
			for node, count := range localStats {
				distribution[node] += count
			}
			mu.Unlock()
		}(keys[start:end])
	}

	wg.Wait()

	skew := analyzer.CalculateSkew(distribution)
	fmt.Printf("[Skew] Standard Deviation: %.2f (Lower is better)\n", skew)

	algo.AddNode("node-10")

	movedCount := 0
	for _, key := range keys {
		newNode := algo.GetShard(key)
		if keyLocation[key] != newNode {
			movedCount++
		}
	}

	percentMoved := (float64(movedCount) / float64(len(keys))) * 100
	moveRows = append(moveRows, MoveRow{
            Algorithm:    algo.Name(),
            PercentMoved: percentMoved,
        })
	writeDistributionCSV("distribution.csv", distRows)
    writeMovementCSV("movement.csv", moveRows)
	fmt.Printf("[Movement] Keys Moved: %.2f%%\n", percentMoved)
	fmt.Println("---------------------------------------------")
}
