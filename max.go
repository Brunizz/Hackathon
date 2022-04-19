package piscine

func Max(a []int) int {
	BiggestNum := a[0]
	for i := 1; i < len(a); i++ {
		if BiggestNum < a[i] {
			BiggestNum = a[i]
		}
	}
	return BiggestNum
}
