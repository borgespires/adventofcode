package day17

import (
    "strconv"

    "adventofcode/io"
)

func fill(liters int, containers []int) int {
    var rec func(l int, idx int, acc []int)
    count := 0

    rec = func(l int, idx int, acc []int) {
        bucket := containers[idx]

        if bucket == l || l == 0 { count++; return }

        if bucket > l { return }

        for i := idx+1; i < len(containers); i++ {
            rec(l-bucket, i, append(acc, bucket))
        }
    }

    for i := 0; i < len(containers); i++ {
        rec(liters, i, []int{})
    }

    return count
}

func fillMin(liters int, containers []int) int {
    var rec func(l int, idx int, acc []int)
    count := 0
    minbuckets := len(containers)

    rec = func(l int, idx int, acc []int) {
        bucket := containers[idx]

        if len(acc)+1 > minbuckets { return }

        if bucket == l || l == 0 {
            if nbuckets := len(acc)+1; nbuckets < minbuckets {
                count = 1
                minbuckets = nbuckets
            } else if nbuckets == minbuckets {
                count++
            }

            return 
        }

        if bucket > l { return }

        for i := idx+1; i < len(containers); i++ {
            rec(l-bucket, i, append(acc, bucket))
        }
    }

    for i := 0; i < len(containers); i++ {
        rec(liters, i, []int{})
    }

    return count
}

func Solve(filepath string) (int, int) {
    lines := io.Readlines(filepath)
    containers := []int{}

    for _, str := range lines {
        v, _ := strconv.Atoi(str)
        containers = append(containers, v)
    }

    return fill(150, containers), fillMin(150, containers)
}