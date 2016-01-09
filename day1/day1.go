package day1

import "adventofcode/io"

// func endFloor(filepath string) int {
//     s := readlines(filepath)[0]
//     up := strings.Count(s, "(")
//     return up - (len(s) - up)
// }

func Solve(filepath string) (int, int) {
    s := io.Readlines(filepath)[0]
    c := 0
    var ibelowzero int

    for i, r := range s {
        if r == 40 { c++ } else { c-- }

        if (ibelowzero == 0 && c < 0) {
            ibelowzero = i + 1
        }
    }

    return c, ibelowzero
}