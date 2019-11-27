package books

// HasSuffixAndPreffix префиксный и суффиксный проверка вхождения строки
func HasSuffixAndPreffix(s, substr string) bool {
	return len(s) >= len(substr) && (s[len(s)-len(substr):] == substr || s[:len(substr)] == substr)
}

// Contains Старт
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasSuffixAndPreffix(s[i:], substr) {
			return true
		}
		// else if HasPrefix(s[i:], substr) {
		// 	return true
		// }
	}
	return false
}
