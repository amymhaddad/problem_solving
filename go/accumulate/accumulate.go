package accumulate

type converter func(string) string

//Accumulate applies a function to a collection
func Accumulate(collection []string, fn converter) []string {

	updatedCollection := []string{}
	for _, item := range collection {
		updatedItem := fn(item)
		updatedCollection = append(updatedCollection, updatedItem)
	}
	return updatedCollection

}
