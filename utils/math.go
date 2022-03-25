package utils

import "math"

func RoundFloat(val float64) float64 {
	return math.Floor(val*100) / 100
}
