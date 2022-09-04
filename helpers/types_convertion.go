package helpers

import (
	"fmt"
	"strconv"
)

func GetUint(s string) uint {
	u64, err := strconv.ParseUint(s, 10, 32)
    if err != nil {
        fmt.Println(err)
    }
    return uint(u64)
}