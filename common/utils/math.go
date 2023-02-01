package utils

func MinInt(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func MaxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func MaxInt64(i, j int64) int64 {
	if i > j {
		return i
	}
	return j
}

func MinInt64(i, j int64) int64 {
	if i <= j {
		return i
	}
	return j
}

func MaxInt32(i, j int32) int32 {
	if i > j {
		return i
	}
	return j
}

func MinInt32(i, j int32) int32 {
	if i <= j {
		return i
	}
	return j
}
