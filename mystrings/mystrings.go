package mystrings

// we need to capitalize the first letter of the function, so that this function can be exported
func Reverse(s string) string {
	res := ""
	for _, v :=  range s {
		res = string(v) + res
	}
	return res
}