package day18

import "adventofcode/io"

const ON = 1
const OFF = 0

func countNeighborsOn(lights []int, width int) (res []int) {
    size := len(lights)

    for i := 0; i < size; i++ {
        count := 0
        line := i/width

        if u := i - width; u >= 0 {
            if ur := u+1; ur/width == line-1 { count += lights[ur] }
            if ul := u-1; ul >= 0 && ul/width == line-1 { count += lights[ul] }
            count += lights[u]
        }

        if d := i + width; d < size {
            if dr := d+1; dr/width == line+1 { count += lights[dr] }
            if dl := d-1; dl/width == line+1 { count += lights[dl] }
            count += lights[d]
        }

        l := i-1
        r := i+1

        if l >= 0 && l/width == line { count += lights[l] }
        if r < size && r/width == line { count += lights[r] }

        res = append(res, count)
    }
    return
}

func animate(l []int, w int, stuck []int, n int) []int {
    var contains func(s []int, ele int) bool
    lights := make([]int, len(l))
    copy(lights, l)
    
    contains = func(s []int, ele int) bool {
        for _, a := range s {
            if a == ele { return true }
        }
        return false
    } 

    for _, i := range stuck { lights[i] = ON }

    for i := 0; i < n; i++ {
        neighbors := countNeighborsOn(lights, w)

        for i := 0; i < len(lights); i++ {
            if contains(stuck, i) { continue }
            if lights[i] == ON && (neighbors[i] < 2 || neighbors[i] > 3) { lights[i] = OFF }
            if lights[i] == OFF && neighbors[i] == 3 { lights[i] = ON }
        }
    }

    return lights
}

func Solve(filepath string) (int, int) {
    lines := io.Readlines(filepath)
    lights := []int{}
    width := len(lines[0])
    
    cnt := 0
    cntstuck := 0

    for _, str := range lines {
        for _, r := range str {
            var state int
            if r == 35 { state = ON }
            
            lights = append(lights, state)
        }
    }

    size := len(lights)
    stuck := []int{0, width-1, size-width, size-1}

    final := animate(lights, width, []int{}, 100)
    finalStuck := animate(lights, width, stuck, 100)

    for i := 0; i < size; i++ { 
        cnt += final[i]
        cntstuck += finalStuck[i]
    }

    return cnt, cntstuck
}