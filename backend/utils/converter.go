package utils

import (
	"log"
	"strconv"
	"time"
	"unsafe"
)

func ToUint(s string) (uint, error) {
	const fnToUint = "ToUint"

	if s == "undefined" {
		return 0, nil
	}

	var (
		result  uint
		byteStr []byte
		sLen    int
		counter int
	)

	sLen = len(s)
	byteStr = append(byteStr, s...)

	// Fast Conversion
	if 0 < sLen && sLen < 10 {

		if byteStr[0] == '-' {
			result = 0
			return result, nil
		}

		if byteStr[0] == '+' {
			byteStr = byteStr[1:]
			if len(byteStr) < 1 {
				return 0, &strconv.NumError{Func: fnToUint, Num: s, Err: strconv.ErrSyntax}
			}
		}

		result = 0
		for counter = 0; counter < sLen; counter++ {
			byteStr[counter] -= '0'
			if byteStr[counter] > 9 {
				return 0, &strconv.NumError{Func: fnToUint, Num: s, Err: strconv.ErrSyntax}
			}
			result = result*10 + uint(byteStr[0])
		}

		return result, nil
	}

	// Slow path for invalid, big, or underscored integers.
	i64, err := strconv.ParseUint(s, 10, 0)
	if nErr, ok := err.(*strconv.NumError); ok {
		nErr.Func = fnToUint
	}
	return uint(i64), err
}

func TryParseInt(value string, defaultValue int) int {
	numValue, err := strconv.Atoi(value)
	if err != nil {
		log.Print(err)
		numValue = defaultValue
	}
	return numValue
}

func TryParseDuration(value string, defaultValue int) time.Duration {
	numValue := TryParseInt(value, defaultValue)
	result := time.Duration(numValue)
	return result
}

func ToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
