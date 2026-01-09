package main

import (
	"fmt"

	"github.com/TeddyMuli/sharding_techniques/shard_simulator/phase_1/pkg/algorithms"
	"github.com/TeddyMuli/sharding_techniques/shard_simulator/phase_1/pkg/generator"
	"github.com/TeddyMuli/sharding_techniques/shard_simulator/phase_1/pkg/utils"
)

func main() {
	competitors := []algorithms.Sharder{
		algorithms.NewModulo(),
		algorithms.NewConsistent(),
		algorithms.NewRange(10000),
		algorithms.NewDirectory(),
		algorithms.NewGeo(),
	}
	
	keyCount := 100_000
	keys := generator.GenerateKeys(keyCount)
	fmt.Printf("Generated %d keys for benchmarking...\n\n", keyCount)

	for _, algo := range competitors {
		utils.RunBenchmark(algo, keys)
	}
}
