package internal

import "math"

// Percentage calculates percentage.
func Percentage(total, current, max float64) float64 {
	return (current / total) * 100
}

// PercentageRound returns a rounded Percentage.
func PercentageRound(total, current, max float64) float64 {
	return math.Round(Percentage(total, current, max))
}
