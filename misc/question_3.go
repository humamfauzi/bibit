package question
func NewFunction(str string) string {
	if len(str) == 0 {
		return ""
	}
	indexFirstBracketFound := strings.Index(str, "(")
	indexClosingBracketFound := strings.Index(str, ")")
	if indexFirstBracketFound < 0 || indexClosingBracketFound < 0 {
		return ""
	}
	runes := []rune(str)
	return string(runes[indexFirstBracketFound+1 : indexClosingBracketFound-1])
}
