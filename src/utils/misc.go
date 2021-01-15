package utils

func isInSlice(target string, slice []ChartObj) (bool, ChartObj) {
	for _, obj := range slice {
		switch v := obj.(type) {
		case *Directory:
			if target == v.Name {
				return true, v
			}
		case *File:
			continue
		}
	}
	return false, nil
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
