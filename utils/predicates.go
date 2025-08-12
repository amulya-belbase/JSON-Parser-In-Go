package utils

var PredicateMapFuncs = map[string]interface{}{
	"is_odd_or_even": isOddOrEven,
}

func isOddOrEven(num int) bool {
	return num%2 == 0
}
