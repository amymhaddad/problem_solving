//hamming problem
package hamming

import "errors"

//Count the differences between two strands
func Distance(a, b string) (int, error) {
	var count int

	if len(a) != len(b) {
		return count, errors.New("the strands must be equal length")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count += 1
		}
	}
	return count, nil
}
