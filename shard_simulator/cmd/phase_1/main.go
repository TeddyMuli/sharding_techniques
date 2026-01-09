package main

import (
	"fmt"

	"github.com/TeddyMuli/sharding_techniques/shard_simulator/phase_1/pkg/algorithms"
	"github.com/TeddyMuli/sharding_techniques/shard_simulator/phase_1/pkg/generator"
	"github.com/TeddyMuli/sharding_techniques/shard_simulator/phase_1/pkg/utils"
)

func main() {
	var allDistRows []utils.DistRow
    var allMoveRows []utils.MoveRow
    
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
		dRows, mRow := utils.RunBenchmark(algo, keys)
		allDistRows = append(allDistRows, dRows...)
        allMoveRows = append(allMoveRows, mRow)
	}
	
	utils.WriteDistributionCSV("../visualization/phase_1/distribution.csv", allDistRows)
    utils.WriteMovementCSV("../visualization/phase_1/movement.csv", allMoveRows)
    
    fmt.Println("\nFinal CSVs generated with all algorithms.")
}
