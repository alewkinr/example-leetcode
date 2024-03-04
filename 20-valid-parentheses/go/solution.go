package solution

func isValid(s string) bool {
	opens := []string{"(", "{", "["}
	set := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	stack := make([]string, 1)
	for _, char := range s {
		if isin(opens, char) {
			stack = append(stack, string(char))
			continue
		}

		if pop(stack) != set[string(char)] {
			return false
		}
		remove(&stack)
	}

	if len(stack) != 1 {
		return false
	}
	return true

}

func isin(s []string, v rune) bool {
	_v := string(v)
	for _, item := range s {
		if _v == item {
			return true
		}
	}
	return false
}

func pop(stack []string) string {
	return stack[len(stack)-1]
}

func remove(stack *[]string) {
	v := *stack
	*stack = v[:len(v)-1]
}
