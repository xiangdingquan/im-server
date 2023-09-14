package math2

// EPSILON very small number
var EPSILON float32 = 0.00000001

// FloatEquals float equals
func FloatEquals(a, b float32) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}

// EPSILON64 very small number
var EPSILON64 = 0.00000001

// Float64Equals float equal
func Float64Equals(a, b float64) bool {
	if (a-b) < EPSILON64 && (b-a) < EPSILON64 {
		return true
	}
	return false
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Int32Min(x, y int32) int32 {
	if x < y {
		return x
	}
	return y
}

func Int32Max(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}

func Int64Min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func Int64Max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
