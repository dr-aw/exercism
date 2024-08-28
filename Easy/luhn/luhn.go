package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if id == "0" {
		return false
	}
	n := len(id)
	sum := 0
	double := false // we need every second digit

	for i := n - 1; i >= 0; i-- { // iteration from the end
		digit, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return false
		}
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}

	return sum%10 == 0
}
