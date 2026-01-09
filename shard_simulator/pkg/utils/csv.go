package utils

import (
	"os"
	"fmt"
	"encoding/csv"
	"strconv"
	"log"
)

type LatencyRow struct {
    Algorithm  string
    Throughput float64
    P50        float64
    P90        float64
    P99        float64
}

func WriteDistributionCSV(filename string, rows []DistRow) {
    file, _ := os.Create(filename)
    defer file.Close()
    
    writer := csv.NewWriter(file)
    defer writer.Flush()

    writer.Write([]string{"Algorithm", "ShardID", "KeyCount"})
    for _, r := range rows {
        writer.Write([]string{r.Algorithm, r.ShardID, strconv.Itoa(r.KeyCount)})
    }
}

func WriteMovementCSV(filename string, rows []MoveRow) {
    file, _ := os.Create(filename)
    defer file.Close()
    
    writer := csv.NewWriter(file)
    defer writer.Flush()

    writer.Write([]string{"Algorithm", "PercentMoved"})
    for _, r := range rows {
        writer.Write([]string{r.Algorithm, fmt.Sprintf("%.2f", r.PercentMoved)})
    }
}

func WriteLatencyCSV(filename string, rows []LatencyRow) {
    file, err := os.Create(filename)
    if err != nil {
        log.Fatalf("Could not create CSV: %v", err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    writer.Write([]string{"Algorithm", "Throughput", "P50_ms", "P90_ms", "P99_ms"})

    for _, r := range rows {
        writer.Write([]string{
            r.Algorithm,
            fmt.Sprintf("%.2f", r.Throughput),
            fmt.Sprintf("%.4f", r.P50),
            fmt.Sprintf("%.4f", r.P90),
            fmt.Sprintf("%.4f", r.P99),
        })
    }
}
