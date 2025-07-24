package actions

var MappedFuncActions = map[string]func(int, int) int{
	"action.add": addTwoNumbers,
	"action.sub": subTwoNumbers,
}

func addTwoNumbers(a, b int) int {
	return a + b
}

func subTwoNumbers(a, b int) int {
	return a - b
}
