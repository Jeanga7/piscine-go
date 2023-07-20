package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		tab1 := []int(nil)
		return tab1
	}
	ans := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		ans[i] = min + i
	}
	return ans
}
