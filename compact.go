package piscine

func Compact(ptr *[]string) int {
	count := 0
	length := len([])
	for i := 0; i < length; i++ {
		if (*ptr)[i] != " " {
			count++
		}
	}
	return count
}
