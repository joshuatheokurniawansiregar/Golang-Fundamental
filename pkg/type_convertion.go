package pkg

import "log"

func CallTypeConvertionIntToInt() int32 {
	var a int8 = 1
	var b int16 = 2
	var c int32 = int32(a) + int32(b)
	return c
}

func CallTypeConvertionFloat32ToFloat64() float64 {
	var a float32 = 2.5
	var b float32 = 2.5
	var c float64 = float64(a) * float64(b)
	return c
}

func CallTypeConvertionInt64ToFloat64(val1 int64) float64 {
	var var_result = float64(val1)
	var abstract_var_int64 interface{} = val1

	_, b := abstract_var_int64.(int64)

	if b != true {
		log.Fatal("Log fatal: val1.Expectation: argument 1: int64")
		log.Panic("Log panic: val1.Expectation: argument 1: int64")
	}
	return var_result
}
