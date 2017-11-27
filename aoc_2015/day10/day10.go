package day10

import (
    "bytes"
    "strconv"
)

func lookandsay(s string) string {
    var b bytes.Buffer
    prev := string(s[0])
    count := 1
    s = s + "_"

    for i := 1; i < len(s); i++ {
        c := string(s[i])

        if c == prev { count++; continue }

        b.WriteString(strconv.Itoa(count) + prev)

        count = 1
        prev = c
    }

    return b.String()
}

func looper(s string, n int) string {
    for n > 0 {
        s = lookandsay(s)
        n--
    }

    return s
}

func Solve(s string) (int, int) {
    return len(looper(s, 40)), len(looper(s, 50))
}