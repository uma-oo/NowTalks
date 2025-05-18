package utils

func CheckPOSTCategories(categories []any) bool {
	for _, category := range categories {
		_, ok := category.(float64)
		if !ok {
			return false
		}
	}
	return true
}
