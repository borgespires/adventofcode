package day24

import (
    "sort"
    "strconv"

    "adventofcode/io"
)

type Group []int
func (g Group) minus(o Group) Group {
    set := make(map[int]int)
    res := Group{}

    for _, w := range g {
        if v, exists := set[w]; exists { set[w] = v+1 } else { set[w] = 1 }
    }

    for _, w:= range o {
        if v, exists := set[w]; exists {
            if v > 1 { set[w] = v-1 } else { delete(set, w) }
        }
    }

    for w, n := range set {
        for n > 0 { 
            res = append(res, w)
            n--
        }
    }

    return res
}
func (g Group) sum() (acc int) {
    for _, w := range g { acc += w }
    return
}
func (g Group) qe() int {
    acc := 1
    for _, w := range g { acc *= w }
    return acc
}

type ByQE []Group
func (g ByQE) Len() int { return len(g) }
func (g ByQE) Swap(i, j int) { g[i], g[j] = g[j], g[i] }
func (g ByQE) Less(i, j int) bool {
    if len(g[i]) == len(g[j]) {
        return g[i].qe() < g[j].qe()
    }

    return len(g[i]) < len(g[j])
}

func canbesplit(items Group, nbins, target int) bool {
    groups := combinations(items, target, nbins)

    if nbins==2 && len(groups) > 0 { return true }

    for _, g := range groups {
        if canbesplit(items.minus(g), nbins-1, target) { return true }
    }

    return false
}

func combinations(items Group, target, nbins int) []Group {
    var rec func(l int, idx int, acc []int)
    res := []Group{}

    maxGroupSize := len(items)/nbins

    rec = func(spaceLeft int, idx int, acc []int) {
        weight := items[idx]

        if (len(acc) >= maxGroupSize) { return }

        if weight == spaceLeft {
            res = append(res, append(acc, weight))
            return
        }

        if weight > spaceLeft { return }

        for i := idx+1; i < len(items); i++ {
            rec(spaceLeft-weight, i, append(acc, weight))
        }
    }

    for i := 0; i < len(items); i++ {
        rec(target, i, []int{})
    }

    return res
}

func balanceSleigh(items []int, nbins int) int {
    initial := Group(items)

    totalWeight := initial.sum()
    target := totalWeight / nbins;

    if totalWeight % nbins != 0 { return -1 }

    c := combinations(initial, target, nbins)
    sort.Sort(ByQE(c))

    if nbins > 2 {
        for _, opt := range c {
            if canbesplit(initial.minus(opt), nbins-1, target) { return opt.qe() }
        }
    } else {
        return c[0].qe()
    }

    return -1
}

func Solve(filepath string) (int, int) {
    lines := io.Readlines(filepath)

    in := []int{}

    for _, s := range lines {
        w, _ := strconv.Atoi(s)
        in = append(in, w)
    }

    return balanceSleigh(in, 3), balanceSleigh(in, 4)
}