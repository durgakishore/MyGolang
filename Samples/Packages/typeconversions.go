package typeconversions

import (
	"strconv"
)

// StringToInt is to convert string to int
func StringToInt(str string) int {

	intValue, _ := strconv.Atoi(str)

	return intValue
}

// StringToFloat is to convert string to float
func StringToFloat(str string) float64 {

	f, _ := strconv.ParseFloat(str, 8)

	return f
}
// StringToBoolean is to convert string to bool
func StringToBoolean(str string) bool {

	b1, _ := strconv.ParseBool(str)
	return b1
}
// BooleanToString is to convert boolean to string
func BooleanToString(val bool) string {

	return strconv.FormatBool(val)

}
//FloatToString is to convert float to string
func FloatToString(val float64) string {

	return strconv.FormatFloat(val, 'E', -1, 32)

}
//IntToString is to convert int to string
func IntToString(val int64) string {

	return strconv.FormatInt(val, 10)

}
//Float32ToFloat64 is to convert float32 to float64
func Float32ToFloat64(val float32) float64 {

	return float64(val)

}
//Float64ToFloat32 is to convert float64 to float32
func Float64ToFloat32(val float64) float32 {

	return float32(val)

}
//IntToFloat is to convert int to float64
func IntToFloat(val int) float64 {

	return float64(val)

}
