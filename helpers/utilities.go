package helpers

import (
	"fmt"
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func GetUint(s string) uint {
	u64, err := strconv.ParseUint(s, 10, 32)
    if err != nil {
        fmt.Println(err)
    }
    return uint(u64)
}

func HashAndSalt(pwd []byte) string {
    
    // Use GenerateFromPassword to hash & salt pwd.
    // MinCost is just an integer constant provided by the bcrypt
    // package along with DefaultCost & MaxCost. 
    // The cost can be any value you want provided it isn't lower
    // than the MinCost (4)
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        log.Println(err)
    }
    // GenerateFromPassword returns a byte slice so we need to
    // convert the bytes to a string and return it
    return string(hash)
}