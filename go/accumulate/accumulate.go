package accumulate

type converter func(string) string

func Accumulate(collection []string, fn converter) []string{

	return ["hello", "world"]
}
