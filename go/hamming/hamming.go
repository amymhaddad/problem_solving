//hamming problem
package hamming

import "errors"

//Count the differences between two strings
func Distance(s1, s2 string) (int, error) {

	var count int

	if len(s1) != len(s2) {
		return count, errors.New("The strands must be equal length.")
	}

	b1 := []byte(s1)
	b2 := []byte(s2)

	for i := 0; i < len(s1); i++ {
		if b1[i] != b2[i] {
			count += 1
		}
	}
	return count, nil
}
