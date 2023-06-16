package utils

// If
// @Description: 三目表达式
// @param condition
// @param trueVal
// @param falseVal
// @return interface{}
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
