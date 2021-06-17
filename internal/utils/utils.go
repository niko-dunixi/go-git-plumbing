package utils

func IndexOfStringInSlice(item string, items ...string) int {
	for i := range items {
		if item == items[i] {
			return i
		}
	}
	return -1
}
