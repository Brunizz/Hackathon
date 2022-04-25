package piscine

func Unmatch(a []int) int {
	for _, i := range a {
		notpair := 0
		for _, j := range a {
			if j == i {
				notpair++
			}
		}
		if notpair == 1 || notpair%2 == 2 {
			return i
		}
	}
	return -1
}
