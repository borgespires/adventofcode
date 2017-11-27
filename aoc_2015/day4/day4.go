package day4

import (
    "strings"
    "strconv"
    "crypto/md5"
    "encoding/hex"
)

func mine(key string, nzeros int) int {
    hasher := md5.New()

    for i:=1; i < 10000000; i++ {
        nkey := []byte(key + strconv.Itoa(i))
        
        hasher.Write(nkey)
        hash := hex.EncodeToString(hasher.Sum(nil))

        if strings.Count(string(hash[:nzeros]), "0") == nzeros {
            return i
        }

        hasher.Reset()
    }

    return 0
}

func Solve(key string) (int, int) {
    return mine(key, 5), mine(key, 6)
}
