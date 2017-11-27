package day20

import "math"

func firstToGet_old(min int) int {
    var countpresents func(int) int
    h := 1

    countpresents = func(h int) int {
        sum := 0;
        d := int(math.Sqrt(float64(h))) + 1
        
        for i := 1; i <= d; i++ {
            if h % i == 0 {
                sum += i
                sum += h/i
            }
        }

        return sum;
    }

    for countpresents(h) < min/10 { h++ }

    return h
}

func firstToGet(min int) int {
    houses := make([]int, min/10 + 1)
    for elf := 1; elf < len(houses); elf++ {
        for house := elf; house < len(houses); house += elf {
            houses[house] += elf * 10
        }
    }

    for house := 1; house < len(houses); house++ {
        if (houses[house] > min) { return house }
    }
    return 0
}

func _firstToGet(min int) int {
    houses := make([]int, min/11 + 1)
    for elf := 0; elf < len(houses); elf++ {
        for house, n := elf, 0; house < len(houses) && n<50; house, n = house+elf, n+1 {
            houses[house] += elf * 11
        }
    }

    for house := 1; house < len(houses); house++ {
        if (houses[house] > min) { return house }
    }
    return 0
}

func Solve(min int) (int, int) {
    return firstToGet(min), _firstToGet(min)
}