package utils

func CheckPOSTCategories(categories []any) bool {
	for _, category := range categories {
		_, ok := category.(int)
		if !ok {
			return false
		}
	}
	return true
}
