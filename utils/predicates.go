package utils

var PredicateMapFuncs = map[string]interface{}{
	"is_even": isEven,
}

func isEven(num int) bool {
	return num%2 == 0
}
