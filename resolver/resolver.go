package resolver

import (
	"fmt"
	"json_parser/actions"
	"json_parser/cache"
	"json_parser/utils"
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

func ResolveMappedFunc(mappedFunc string, args ...interface{}) interface{} {
	fn, exists := actions.MappedFuncActions[mappedFunc]
	if !exists {
		return fmt.Sprintf("mapped function '%s' not found", mappedFunc)
	}

	return castFunction(fn, args)
}

func castFunction(fn interface{}, args []interface{}) interface{} {
	argsLen := len(args)

	// perform type assertions based on expected function signatures
	// can add more function signature checks if needed
	if f, ok := fn.(func(int, int) int); ok {
		if argsLen != 2 {
			return fmt.Sprintf("expected 2 arguments, got %d", argsLen)
		}

		a, err := utils.ConvertToInt(args[0])
		if err != nil {
			return fmt.Sprintf("invalid first argument: %v. error: %v", args[0], err)
		}

		b, err := utils.ConvertToInt(args[1])
		if err != nil {
			return fmt.Sprintf("invalid second argument: %v. error: %v", args[1], err)
		}

		return f(a, b)
	}

	if f, ok := fn.(func(string, string) string); ok {
		if argsLen != 2 {
			return fmt.Sprintf("expected 2 arguments, got %d", argsLen)
		}

		a, err := utils.ConvertToString(args[0])
		if err != nil {
			return fmt.Sprintf("invalid first argument: %v. error: %v", args[0], err)
		}

		b, err := utils.ConvertToString(args[1])
		if err != nil {
			return fmt.Sprintf("invalid second argument: %v. error: %v", args[1], err)
		}

		return f(a, b)
	}

	return fmt.Sprintf("invalid function signature: %v", fn)
}
