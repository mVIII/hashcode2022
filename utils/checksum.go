package utils

import (
	"fmt"
	"hash/fnv"
)

func Checksum(s string) string {
	hash := fnv.New128()
	hash.Write([]byte(s))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
