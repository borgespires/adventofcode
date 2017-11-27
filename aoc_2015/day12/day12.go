package day12

import (
    "strconv"
    "regexp"
    "encoding/json"

    "adventofcode/io"
)

func addAll(blob string) int {
    var out int

    regexp.MustCompile(`[\-0-9]+`).ReplaceAllStringFunc(blob, func(match string) string {
        i, _ := strconv.Atoi(match)
        out += i
        return match
    })
    return out
}

func ignoreRed(blob string) int {
    var f interface{}
    json.Unmarshal([]byte(blob), &f)

    return int(addDataType(f))
}

func addDataType(f interface{}) (out float64) {
    out:
    switch fv := f.(type) {
        case []interface{}: // if array
            for _, val := range fv {
                out += addDataType(val)
            }
        case float64: // if number
            out += fv
        case map[string]interface{}: // if object
            for _, val := range fv {
                if val == "red" {
                    break out
                }
            }
            for _, val := range fv {
                out += addDataType(val)
            }
    }
    return
}

func Solve(filepath string) (int, int) {
    blob := io.Readlines(filepath)[0]

    return addAll(blob), ignoreRed(blob)
}