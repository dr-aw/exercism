package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	d := 0
	if len(a) != len(b) {
		return 0, fmt.Errorf("different length (%d) and (%d)", len(a), len(b))
	}
	for i := range a {
		if a[i] != b[i] {
			d++
		}
	}
	return d, nil
}
