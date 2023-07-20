package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	var r []int
	for i := min; i < max; i++ {
		r = append(r, i)
	}
	return r
}
