package utils

import (
	"os"
	"fmt"
	"encoding/csv"
	"strconv"
)

func WriteDistributionCSV(filename string, rows []DistRow) {
    file, _ := os.Create(filename)
    defer file.Close()
    
    writer := csv.NewWriter(file)
    defer writer.Flush()

    writer.Write([]string{"Algorithm", "ShardID", "KeyCount"}) // Header
    for _, r := range rows {
        writer.Write([]string{r.Algorithm, r.ShardID, strconv.Itoa(r.KeyCount)})
    }
}

func WriteMovementCSV(filename string, rows []MoveRow) {
    file, _ := os.Create(filename)
    defer file.Close()
    
    writer := csv.NewWriter(file)
    defer writer.Flush()

    writer.Write([]string{"Algorithm", "PercentMoved"}) // Header
    for _, r := range rows {
        writer.Write([]string{r.Algorithm, fmt.Sprintf("%.2f", r.PercentMoved)})
    }
}
