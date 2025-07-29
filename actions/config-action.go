package actions

var MappedFuncActions = map[string]interface{}{
	"action.add":           addTwoNumbers,
	"action.sub":           subTwoNumbers,
	"action.string_concat": stringConcat,
}

func addTwoNumbers(a, b int) int {
	return a + b
}

func subTwoNumbers(a, b int) int {
	return a - b
}

func stringConcat(a, b string) string {
	return a + b
}
