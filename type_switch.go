

func sum(numbers interface{}, length int) (float64, bool) {
	switch nums := numbers.(type) {
	case []float64:
		sum := 0.0
		for _, num := range nums {
			sum += num
		}
		return sum, true
	case []int:
		sum := 0
		for _, num := range nums {
			sum += num
		}
		return float64(sum), true
	default:
		return 0, false
	}
}
