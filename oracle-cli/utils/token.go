package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ConvertToXfund(amount string) float64 {
	tokens := strings.TrimSpace(amount)
	i, err := strconv.Atoi(tokens)
	if err != nil {
		fmt.Println("error", err)
		return 0
	}
	return float64(i) / math.Pow10(9)
}
