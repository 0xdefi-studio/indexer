package cmd

import (
	"fmt"
	"math/big"
	"strings"
)

func DurationToTimeFormat(durationInSeconds *big.Int) string {
	secondsInMinute := big.NewInt(60)
	secondsInHour := big.NewInt(3600)

	// Extract hours, minutes, and seconds
	hours := new(big.Int).Div(durationInSeconds, secondsInHour)
	durationInSeconds.Sub(durationInSeconds, new(big.Int).Mul(hours, secondsInHour))

	minutes := new(big.Int).Div(durationInSeconds, secondsInMinute)
	durationInSeconds.Sub(durationInSeconds, new(big.Int).Mul(minutes, secondsInMinute))

	seconds := durationInSeconds

	// Build the time format string
	var timeParts []string
	if hours.Cmp(big.NewInt(0)) > 0 {
		timeParts = append(timeParts, fmt.Sprintf("%dh", hours))
	}

	if minutes.Cmp(big.NewInt(0)) > 0 {
		timeParts = append(timeParts, fmt.Sprintf("%dmin", minutes))
	}

	if seconds.Cmp(big.NewInt(0)) > 0 {
		timeParts = append(timeParts, fmt.Sprintf("%dsec", seconds))
	}

	return strings.Join(timeParts, " ")
}
