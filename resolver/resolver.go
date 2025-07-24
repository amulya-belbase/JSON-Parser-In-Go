package resolver

import (
	"fmt"
	"json_parser/actions"
	"json_parser/cache"
	"strconv"
)

func ResolveWithActionProcessID(action, processId string) interface{} {
	actionData, exists := cache.AllActions[action]
	if !exists {
		return fmt.Sprintf("action '%s' not found", action)
	}

	for _, process := range actionData.Processes {
		if process.ProcessId == processId {
			return ResolveMappedFunc(process.MappedFunc, process.First, process.Second)
		}
	}

	return fmt.Sprintf("process '%s' not found", processId)
}

func ResolveMappedFunc(mappedFunc, first, second string) interface{} {
	fn, exists := actions.MappedFuncActions[mappedFunc]
	if !exists {
		return fmt.Sprintf("mapped function '%s' not found", mappedFunc)
	}

	// string to int
	a, err1 := strconv.Atoi(first)
	b, err2 := strconv.Atoi(second)
	if err1 != nil || err2 != nil {
		return fmt.Sprintf("invalid number inputs: %s, %s", first, second)
	}

	// execute the function
	result := fn(a, b)
	return result
}
