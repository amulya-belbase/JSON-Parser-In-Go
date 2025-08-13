package resolver

import (
	"fmt"
	"json_parser/actions"
	"json_parser/cache"
	"json_parser/utils"
	"strings"
)

func ResolveWithActionProcessID(action, processId string) interface{} {
	actionData, exists := cache.AllActions[action]
	if !exists {
		return fmt.Sprintf("action '%s' not found", action)
	}

	for _, process := range actionData.Processes {
		if process.ProcessId == processId {
			val := ResolveMappedFunc(process.MappedFunc, process.First, process.Second)
			if process.Predicate != "" {
				res := ResolvePredicate(process.Predicate, val)
				// TODO: use template engine to evaluate predicate check
				fmt.Println("Evaluated predicate: ", res)
			}
			return val
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

func ResolvePredicate(predicate string, val interface{}) interface{} {
	// parse lisp-like predicate
	parsedPredicateMap := parsePredicates(predicate)

	for k, _ := range parsedPredicateMap {
		fn, exists := utils.PredicateMapFuncs[k]
		if !exists {
			return fmt.Sprintf("predicate function '%s' not found", k)
		}

		parsedPredicateMap[k] = castFunction(fn, []interface{}{val})
	}

	return parsedPredicateMap
}

func parsePredicates(predicate string) map[string]interface{} {
	parsedMap := make(map[string]interface{})

	// trim all the syntactical garbage
	rc := strings.TrimPrefix(predicate, "{{")
	rc = strings.TrimSuffix(rc, "}}")
	rc = strings.ReplaceAll(rc, "(", "")
	rc = strings.ReplaceAll(rc, ")", "")

	for i := 0; i < len(rc); i++ {
		// . is an indicator of a function
		if rc[i] == '.' {
			start := i + 1

			// get function name from . index to next space index
			end := strings.IndexByte(rc[start:], ' ')
			if end == -1 {
				// -1 indicates the end of the string
				end = len(rc)
			} else {
				// adjust substring index to original string
				end += start
			}

			// get the func name
			variable := rc[start:end]
			parsedMap[variable] = nil

			// move i to the end of this variable
			i = end - 1
		}
	}
	return parsedMap
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

	if f, ok := fn.(func(int) bool); ok {
		if argsLen != 1 {
			return fmt.Sprintf("expected 1 argument, got %d", argsLen)
		}

		a, err := utils.ConvertToInt(args[0])
		if err != nil {
			return fmt.Sprintf("invalid first argument: %v. error: %v", args[0], err)
		}

		return f(a)
	}

	return fmt.Sprintf("invalid function signature: %v", fn)
}
