//hamming problem
package hamming

//Count the differences between two strings
func Distance(s1, s2 string) (int, error) {

	b1 := []byte(s1)
	b2 := []byte(s2)

	var count int

	for i := 0; i < len(s1); i++ {
		if b1[i] != b2[i] {
			count += 1
		}
	}
	return count, nil
}
