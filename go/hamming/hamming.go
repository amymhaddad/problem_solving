//hamming problem
package hamming

import "errors"

//Count the differences between two strands
func Distance(a, b string) (int, error) {
	var count int

	if len(a) != len(b) {
		return count, errors.New("The strands must be equal length.")
	}

	strand1 := []byte(a)
	strand2 := []byte(b)

	for i := 0; i < len(a); i++ {
		if strand1[i] != strand2[i] {
			count += 1
		}
	}
	return count, nil
}
