package analyzer

import "math"

func CalculateSkew(distribution map[string]int) float64 {
	var sum, count float64
	values := []float64{}

	for _, v := range distribution {
		val := float64(v)
		sum += val
		values = append(values, val)
		count++
	}

	mean := sum / count

	var varianceSum float64
	for _, v := range values {
		varianceSum += math.Pow(v-mean, 2)
	}

	variance := varianceSum / count

	return math.Sqrt(variance)
}
