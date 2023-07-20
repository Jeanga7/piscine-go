package piscine

func ConcatParams(args []string) string {
	a := make([]string, len(args))
	result := ""
	for i := 0; i < len(a); i++ {

		result += args[i]
		if i < len(args)-1 {
			result += "\n"
		}

	}
	return result
}
