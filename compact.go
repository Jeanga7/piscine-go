package main

func Compact(ptr *[]string) int {
	var tab []string
	for _, i := range *ptr {
		if i != "" {
			tab = append(tab, i)
		}
	}
	return len(tab)
}
