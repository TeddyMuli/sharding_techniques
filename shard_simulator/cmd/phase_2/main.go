package main

import (
	"context"
	"fmt"
	"log"
	"sort"
	"sync"
	"time"

	"github.com/TeddyMuli/sharding_techniques/shard_simulator/phase_1/pkg/algorithms"
	"github.com/TeddyMuli/sharding_techniques/shard_simulator/phase_1/pkg/generator"
	"github.com/TeddyMuli/sharding_techniques/shard_simulator/phase_1/pkg/transport"
	"github.com/TeddyMuli/sharding_techniques/shard_simulator/phase_1/pkg/utils"
)

var registry = map[string]string{
	"node-00": "localhost:6370",
	"node-01": "localhost:6371",
	"node-02": "localhost:6372",
	"node-03": "localhost:6373",
	"node-04": "localhost:6374",
	"node-05": "localhost:6375",
	"node-06": "localhost:6376",
	"node-07": "localhost:6377",
	"node-08": "localhost:6378",
	"node-09": "localhost:6379",
}

func main() {
	var latencyRows []utils.LatencyRow
	
	fmt.Println("Connecting to Docker Cluster...")
	clientPool, err := transport.NewShardClient(registry)
	
	if err != nil {
		log.Fatalf("Infrastructure Error: %v", err)
	}
	defer clientPool.Cleanup()
	fmt.Println("Connected to all shards!")
	
	keys := generator.GenerateKeys(1000)

	for _, algo := range algorithms.Competitors {
		row := RunConcurrentBenchmark(algo, clientPool, keys)
		latencyRows = append(latencyRows, row)
	}
	
	utils.WriteLatencyCSV("visualization/phase_2/latency_results.csv", latencyRows)
}

func CalculateMetrics(latencies []time.Duration) (p50, p90, p99 time.Duration) {
	if len(latencies) == 0 {
		return 0, 0, 0
	}

	sort.Slice(latencies, func(i, j int) bool {
		return latencies[i] < latencies[j]
	})

	p50Idx := int(float64(len(latencies)) * 0.50)
	p90Idx := int(float64(len(latencies)) * 0.90)
	p99Idx := int(float64(len(latencies)) * 0.99)

	return latencies[p50Idx], latencies[p90Idx], latencies[p99Idx]
}

func RunConcurrentBenchmark(sharder algorithms.Sharder, clientPool *transport.ShardClient, keys []string) utils.LatencyRow {
    var wg sync.WaitGroup
    var mu sync.Mutex
    
    for i := range 10 {
        sharder.AddNode(fmt.Sprintf("node-%02d", i))
    }
    
    latencies := make([]time.Duration, 0, len(keys))
    workerCount := 20
    ctx := context.Background()

    chunkSize := len(keys) / workerCount
    startAll := time.Now()

    for i :=range workerCount {
        wg.Add(1)
        
        startIdx := i * chunkSize
        endIdx := startIdx + chunkSize
        if i == workerCount-1 { endIdx = len(keys) }

        go func(kChunk []string) {
            defer wg.Done()
            localLatencies := []time.Duration{}

            for _, key := range kChunk {
                target := sharder.GetShard(key)
                if target == "" {
                    continue
                }
                latency, err := clientPool.Write(ctx, target, key, "value")
                if err == nil {
                    localLatencies = append(localLatencies, latency)
                }
            }

            mu.Lock()
            latencies = append(latencies, localLatencies...)
            mu.Unlock()
        }(keys[startIdx:endIdx])
    }

    wg.Wait()
    totalTime := time.Since(startAll)

    p50, p90, p99 := CalculateMetrics(latencies)
    tps := float64(len(keys)) / totalTime.Seconds()

    fmt.Printf("\n--- Results for %s ---\n", sharder.Name())
    fmt.Printf("Throughput:   %.2f req/sec\n", tps)
    fmt.Printf("p50 Latency:  %v\n", p50)
    fmt.Printf("p90 Latency:  %v\n", p90)
    fmt.Printf("p99 Latency:  %v (Tail Latency)\n", p99)
    fmt.Println("--------------------------------")
    
    return utils.LatencyRow{
        Algorithm:  sharder.Name(),
        Throughput: tps,
        P50:        float64(p50.Microseconds()) / 1000.0,
        P90:        float64(p90.Microseconds()) / 1000.0,
        P99:        float64(p99.Microseconds()) / 1000.0,
    }
}
