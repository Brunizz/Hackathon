package piscine

func Join(strs []string, sep string) string {
	concat := ""
	for i, result := range strs {
		concat += result
		if i != len(strs)-1 {
			concat += sep
		}
	}
	return concat
}
