package day2

import (
    "strings"
    "sort"
    "strconv"

    "adventofcode/io"
)

func getSize(box string) (int, int, int) {
    dim := make([]int, 3)
    split := strings.Split(box, "x")

    for i, e := range split {
        v, _ := strconv.Atoi(e)
        
        dim[i] = v
    }

    sort.Ints(dim)
    return dim[0], dim[1], dim[2]
}

func Solve(filepath string) (int, int) {
    presents := io.Readlines(filepath)
    tpaper := 0
    tribbon := 0

    for _, box := range presents {
        l, w, h := getSize(box)
        paper := 2 * (l*w + w*h + h*l) + l*w
        ribbon := 2 * (l+w) + (l*w*h)

        tpaper += paper
        tribbon += ribbon
    }

    return tpaper, tribbon
}