package day13

import (
    "regexp"
    "strconv"

    "adventofcode/io"
)

/*
/* Permutator
 */
type Permutator struct {
    guestList []string
    state []int
}

func newPermutator(gl []string) *Permutator {
    return &Permutator {
        guestList: gl,
        state: make([]int, len(gl)),
    }
}

func (p *Permutator) updateState() {
    maxidx := len(p.state) - 1
    for i := maxidx; i >= 0; i-- {
        if i == 0 || p.state[i] < maxidx-i { p.state[i]++; return }
        p.state[i] = 0
    }
}

func (p *Permutator) next() []string {
    var swap func(int, int)
    res := append([]string{}, p.guestList...)

    swap = func(i int, j int) {
        res[i], res[j] = res[j], res[i]
    }

    for i, v := range p.state { swap(i, i+v) }

    p.updateState()

    return res
}

func (p *Permutator) hasNext() bool {
    return p.state[0] < len(p.state)
}

func parse(s string) (guest string, happiness int, nextTo string) {
    r := regexp.MustCompile(`^(\w+)\swould\s(gain|lose)\s(\d+).*next\sto\s(\w+).$`)
    match := r.FindStringSubmatch(s)

    if match != nil {
        guest = match[1]
        nextTo = match[4]

        switch match[2] {
        case "gain":
            happiness, _ = strconv.Atoi(match[3])
        case "lose":
            happiness, _ = strconv.Atoi(match[3])
            happiness *= -1
        }
    }

    return
}

func evalTable(table []string, guestList map[string]map[string]int) (total int) {
    size := len(table)
    var abs func(int) int

    abs = func(x int) int {
        switch {
        case x < 0:
            return -x
        case x == 0:
            return 0 // return correctly abs(-0)
        }
        return x
    }

    for i := 0; i < size; i++ {
        prev := table[abs((i-1)+size)%size]
        guest := table[i]
        next := table[(i+1)%size]

        if guest == "me" { continue }

        total += guestList[guest][prev] + guestList[guest][next]
    }
    return
}

func optimalSeating(guestList map[string]map[string]int, table []string) int {
    optimal := 0
    p := newPermutator(table)

    for p.hasNext() {
        if scenario := evalTable(p.next(), guestList); scenario > optimal { optimal = scenario }
    }

    return optimal
}

func Solve(filepath string) (int, int) {
    lines := io.Readlines(filepath)
    guestList := make(map[string]map[string]int)
    table :=  []string{}
    
    var guest map[string]int
    var exists bool

    for _, str := range lines {
        name, happiness, nextTo := parse(str)

        if guest, exists = guestList[name]; !exists {
            guest = map[string]int{}
            guestList[name] = guest
            table = append(table, name)
        }

        guest[nextTo] = happiness
    }

    o := optimalSeating(guestList, table)

    // add me
    for _, guest := range guestList {
        guest["me"] = 0
    }
    table = append(table, "me")
    o2 := optimalSeating(guestList, table)

    return o, o2
}