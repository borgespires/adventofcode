package day8

import (
    "bytes"

    "adventofcode/io"
)

func memorySpace(s string) int {
    count := 0
    
    for i := 1; i < len(s)-1; {
        c := string(s[i:i+2])

        switch c {
        case "\\\\": i+=2
        case "\\\"": i+=2
        case "\\x": i+=4
        default: i++
        }

        count++
    }

    return count
}

func encode(s string) string {
    var b bytes.Buffer

    b.WriteString("\"")

    for i := 0; i < len(s); i++ {
        c := string(s[i])

        switch c {
        case "\\": b.WriteString("\\\\")
        case "\"": b.WriteString("\\\"")
        default: b.WriteString(c)
        }
    }

    b.WriteString("\"")

    return b.String()
}

func Solve(filepath string) (int, int) {
    strings := io.Readlines(filepath)

    cInString := 0
    cInMemory := 0
    cEncoded := 0

    for _, str := range strings {
        cInString += len(str)
        cInMemory += memorySpace(str)
        cEncoded += len(encode(str))
    }

    return cInString - cInMemory, cEncoded - cInString
}