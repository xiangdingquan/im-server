package util

import (
	"strconv"
)

func StringToInt32(s string) (int32, error) {
	i, err := strconv.Atoi(s)
	return int32(i), err
}

func StringToUint32(s string) (uint32, error) {
	i, err := strconv.Atoi(s)
	return uint32(i), err
}

func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func StringToUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

func StringToFloat32(s string) (float32, error) {
	i, err := strconv.ParseFloat(s, 32)
	return float32(i), err
}

func StringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func UInt64ToString(i uint64) string {
	return strconv.FormatUint(i, 10)
}

func Int32ToString(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

func UInt32ToString(i int32) string {
	return strconv.FormatUint(uint64(i), 10)
}

func Float32ToString(i float32) string {
	return strconv.FormatFloat(float64(i), 'f', -1, 64)
}

func Float64ToString(i float64) string {
	return strconv.FormatFloat(i, 'f', -1, 64)
}

func BoolToInt8(b bool) int8 {
	if b {
		return 1
	}
	return 0
}

func Int8ToBool(b int8) bool {
	return b != 0
}

func Int8Bool(b int8) bool {
	return b == 1
}
