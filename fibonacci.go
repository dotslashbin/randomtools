package randomtools

func GetFiboIndex(input int) int {
	if input <= 2 {
		return input
	}

	index := GetFiboIndex(input-1) + GetFiboIndex(input-2)
	return index
}
