package generator

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateKeys(count int) []string {
	rand.NewSource(time.Now().UnixNano())
	
	keys := make([]string, count)
	for i:= range count {
		keys[i] = fmt.Sprintf("user-%d", rand.Intn(10000000))
	}
	
	return keys
}
