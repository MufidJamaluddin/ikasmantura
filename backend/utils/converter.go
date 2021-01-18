package utils

import (
	"log"
	"strconv"
	"time"
	"unsafe"
)

func ToUint(s string) (uint, error) {
	i64, err := strconv.ParseUint(s, 10, 32)
	if nErr, ok := err.(*strconv.NumError); ok {
		nErr.Func = "ToUint"
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
